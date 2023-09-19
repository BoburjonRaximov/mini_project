package memory

import (
	"context"
	"errors"
	"fmt"
	"new_project/models"
	"new_project/pkg/helper"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type staffRepo struct {
	db *pgxpool.Pool
}

func NewStaffRepo(db *pgxpool.Pool) *staffRepo {
	return &staffRepo{db: db}
}

func (s *staffRepo) CreateStaff(req models.CreateStaff) (string, error) {
	fmt.Println("staff create")
	id := uuid.NewString()
	query :=
		`INSERT INTO 
	staffs(id,name,branchId,tarifId,type,balance,createdAt,birthDate) 
VALUES($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := s.db.Exec(context.Background(), query,
		id,
		req.Name,
		req.BranchId,
		req.TarifId,
		req.Type,
		req.Balance,
		req.BirthDate,
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return "", err
	}
	return id, nil
}

func (s *staffRepo) UpdateStaff(req models.Staff) (string, error) {
	query := `
	update staffs
	set name=$2,branchId=$3,tarifId=$4,type=$5,balance=$6,createdAt=$7,birthDate=$8
	where id=$1`
	resp, err := s.db.Exec(context.Background(), query,
		req.Id,
		req.Name,
		req.BranchId,
		req.TarifId,
		req.Type,
		req.Balance,
		req.BirthDate,
	)
	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}
	return "Updated", nil
}

func (s *staffRepo) GetStaff(req models.IdRequestStaff) (models.Staff, error) {
	query := `
	select * from staffs
	where id=$1`
	staff := models.Staff{}
	err := s.db.QueryRow(context.Background(), query, req.Id).Scan(
		&staff.Id,
		&staff.Name,
		&staff.BranchId,
		&staff.TarifId,
		&staff.Type,
		&staff.Balance,
		&staff.BirthDate,
	)
	if err != nil {
		fmt.Println("error scan", err.Error())
	}
	return staff, errors.New("not found")
}

func (b *staffRepo) GetAllStaff(req models.GetAllStaffRequest) (resp models.GetAllStaff, err error) {
	var (
		params  = make(map[string]interface{})
		filter  = "WHERE true "
		offsetQ = " OFFSET 0 "
		limit   = " LIMIT 10 "
		offset  = (req.Page - 1) * req.Limit
	)
	s := `
	SELECT *
	FROM staffs
	`
	if req.Search != "" {
		filter += ` AND name ILIKE '%@search%' `
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

	rows, err := b.db.Query(context.Background(), q, pArr...)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var staff models.Staff
		err := rows.Scan()
		if err != nil {
			return resp, err
		}
		resp.Staffs = append(resp.Staffs, staff)
		resp.Count = len(resp.Staffs)
	}
	return resp, nil
}

func (s *staffRepo) DeleteStaff(req models.IdRequestStaff) (string, error) {
	query := `
	delete from staffs
	where id=$1 `
	resp, err := s.db.Exec(context.Background(), query,
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
