package memory

import (
	"context"
	"fmt"
	"new_project/models"
	"new_project/pkg/helper"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type branchRepo struct {
	db *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) *branchRepo {
	return &branchRepo{db: db}
}


func (b *branchRepo) CreateBranch(ctx context.Context, req models.CreateBranch) (string, error) {
	fmt.Println("branch create")
	id := uuid.NewString()

	query := `
	INSERT INTO 
		branches(
			id,
			name,
			adress) 
	VALUES($1,$2,$3)`
	_, err := b.db.Exec(ctx, query,
		id,
		req.Name,
		req.Address,
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return "error exec", err
	}
	return id, nil
}

func (b *branchRepo) UpdateBranch(ctx context.Context,req models.Branch) (string, error) {
	query := `
	UPDATE 
		branches
	SET 
		name=$2,
		adress=$3
	WHERE 
		id=$1`
	resp, err := b.db.Exec(ctx, query,
		req.Id,
		req.Name,
		req.Address,
	)
	if err != nil {
		return "warning", err
	}
	if resp.RowsAffected() == 0 {
		return "error row", pgx.ErrNoRows
	}
	return "OK", nil
}

func (b *branchRepo) GetBranch(ctx context.Context,req models.IdRequest) (models.Branch, error) {
	query := `
	SELECT
		name,
		adress
	FROM 
		branches
	WHERE 
		id = $1`
	resp := b.db.QueryRow(ctx, query,
		req.Id,
	)
	var branch models.Branch
	err := resp.Scan(
		&branch.Id,
		&branch.Name,
		&branch.Address,
	)
	if err != nil {
		fmt.Println("error scan", err.Error())
	}
	return branch, nil
}
func (b *branchRepo) GetAllBranch(ctx context.Context,req models.GetAllBranchRequest) (resp models.GetAllBranch, err error) {
	var (
		params  = make(map[string]interface{})
		filter  = " WHERE true "
		offsetQ = " OFFSET 0 "
		limit   = " LIMIT 10 "
		offset  = (req.Page - 1) * req.Limit
	)
	s := `
	SELECT
		name,
		adress	
	FROM 
		branches
	`
	if req.Search != "" {
		filter += ` AND name ILIKE '%' || @search || '%' `
		params["search"] = req.Search
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf("LIMIT %d", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf("OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ

	q, pArr := helper.ReplaceQueryParams(query, params)

	rows, err := b.db.Query(ctx, q, pArr...)
	if err != nil {
		return resp, err
	}
	defer rows.Close()
	for rows.Next() {
		var branch models.Branch
		err := rows.Scan(
			&branch.Id,
			&branch.Name,
			&branch.Address,
		)
		if err != nil {
			return resp, err
		}
		resp.Branches = append(resp.Branches, branch)
		resp.Count = len(resp.Branches)
	}
	return resp, nil
}

// delete branch
func (b *branchRepo) DeleteBranch(ctx context.Context,req models.IdRequest) (string, error) {
	query := `
	DELETE FROM 
		branches
	WHERE 
		id=$1 `
	resp, err := b.db.Exec(ctx, query,
		req.Id,
	)
	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}

	return "deleted", nil
}
