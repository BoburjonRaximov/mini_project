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

func (s *staffRepo) CreateStaff(ctx context.Context, req models.CreateStaff) (string, error) {
	fmt.Println("staff create")
	id := uuid.NewString()
	query := `
	INSERT INTO 
	  staffs(id,
			name,
			username,
			password,
			branch_id,
			tariff_id,
			type,
			balance) 
    VALUES($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := s.db.Exec(ctx, query,
		id,
		req.Name,
		req.Username,
		req.Password,
		req.BranchId,
		req.TarifId,
		req.Type,
		req.Balance,
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return "", err
	}
	return id, nil
}

func (s *staffRepo) UpdateStaff(ctx context.Context, req models.Staff) (string, error) {
	query := `
	UPDATE
		staffs
	 SET 
		name=$2,
		branch_id=$3,
		tarif_id=$4,
		type=$5,
		balance=$6,
		updated_at=NOW()
	WHERE 
		id=$1`
	resp, err := s.db.Exec(ctx, query,
		req.Id,
		req.Name,
		req.BranchId,
		req.TarifId,
		req.Type,
		req.Balance,
	)
	if err != nil {
		return "ERROR EXEC", err
	}
	if resp.RowsAffected() == 0 {
		return "ERROR RowsAffected", pgx.ErrNoRows
	}
	return "Updated", nil
}

func (s *staffRepo) GetStaff(ctx context.Context, req models.IdRequestStaff) (models.Staff, error) {
	query := `
	SELECT
		id,
		name,
		branch_id,
		tariff_id,
		type,
		balance,
		created_at::text,
		updated_at::text
	FROM
		staffs
	WHERE
		id=$1`
	staff := models.Staff{}
	err := s.db.QueryRow(ctx, query, req.Id).Scan(
		&staff.Id,
		&staff.Name,
		&staff.BranchId,
		&staff.TarifId,
		&staff.Type,
		&staff.Balance,
	)
	if err != nil {
		fmt.Println("error scan", err.Error())
	}
	return staff, errors.New("not found")
}

func (b *staffRepo) GetAllStaff(ctx context.Context, req models.GetAllStaffRequest) (resp models.GetAllStaff, err error) {
	var (
		params  = make(map[string]interface{})
		filter  = " WHERE true "
		offsetQ = " OFFSET 0 "
		limit   = " LIMIT 10 "
		offset  = (req.Page - 1) * req.Limit
	)
	s := `
	SELECT
		id,
		name,
		branch_id,
		tariff_id,
		type,
		balance,
		created_at::text,
		updated_at::text	
	FROM 
		staffs
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

func (s *staffRepo) DeleteStaff(ctx context.Context, req models.IdRequestStaff) (string, error) {
	query := `
	DELETE FROM 
		staffs
	WHERE 
		id=$1 `
	resp, err := s.db.Exec(ctx, query,
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

// BALANCE
func (s *staffRepo) UpdateBalance(ctx context.Context, req models.UpdateBalanceRequest) (string, error) {
	tr, err := s.db.Begin(ctx)
	defer func() {
		if err != nil {
			tr.Rollback(ctx)
		} else {
			tr.Commit(ctx)
		}
	}()

	cqb := `
	update staffs
	set balance=+$2
	where id=$1`
	if req.TransactionType == "withdraw" {
		req.Cashier.Amount = -req.Cashier.Amount
		req.ShopAssisstant.Amount = -req.ShopAssisstant.Amount
	}
	_, err = tr.Exec(ctx, cqb, req.Cashier.StaffId, req.Cashier.Amount)
	if err != nil {
		return "error exec", err
	}
	// strq := `
	// insert into transactions(
	// 	id,
	// 	staff_id,
	// 	sale_id,
	// 	amount,
	// 	type,
	// 	source_type,
	// 	text
	// )`
	// _, err := tr.Exec(context.Background(), strq,
	//  uuid.NewString(), req)
	if err != nil {
		return "error exec", err
	}
	return "balance updated", nil
}
