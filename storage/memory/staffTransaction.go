package memory

import (
	"context"
	"errors"
	"fmt"
	"new_project/models"
	"new_project/pkg/helper"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type staffTransactionRepo struct {
	db *pgxpool.Pool
}

func NewStaffTransactionRepo(db *pgxpool.Pool) *staffTransactionRepo {
	return &staffTransactionRepo{db: db}
}

func (s *staffTransactionRepo) CreateStaffTransaction(ctx context.Context, req models.CreateStaffTransaction) (string, error) {
	fmt.Println("staff create")
	id := uuid.NewString()
	query := `
	INSERT INTO 
		transactions(
			id,
			sale_id,
			staff_id,
			type,
			source_type,
			amount,text) 
	VALUES($1,$2,$3,$4,$5,$6,$7)
	`
	_, err := s.db.Exec(ctx, query,
		id,
		req.SaleId,
		req.StaffId,
		req.Type,
		req.SourceType,
		req.Amount,
		req.Text,
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return "", err
	}
	return id, nil
}

func (s *staffTransactionRepo) UpdateStaffTransaction(ctx context.Context, req models.StaffTransaction) (string, error) {
	query := `
	UPDATE 
		transactions
	SET 
		sale_id=$2,
		staff_id=$3,
		type=$4,
		source_type=$5,
		amount=$6,
		text=$7
	where 
		id=$1
	`
	resp, err := s.db.Exec(ctx, query,
		req.Id,
		req.SaleId,
		req.StaffId,
		req.Type,
		req.SourceType,
		req.Amount,
		req.Text,
	)
	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}
	return "Updated", nil
}

func (s *staffTransactionRepo) GetStaffTransaction(ctx context.Context, req models.IdRequestStaffTransaction) (models.StaffTransaction, error) {
	query := `
	SELECT
		sale_id,
		staff_id,
		type,
		source_type,
		amount,
		text
	FROM 
		transactions
	WHERE 
		id=$1`
	staffTransaction := models.StaffTransaction{}
	err := s.db.QueryRow(ctx, query, req.Id).Scan(
		&staffTransaction.Id,
		&staffTransaction.SaleId,
		&staffTransaction.StaffId,
		&staffTransaction.Type,
		&staffTransaction.SourceType,
		&staffTransaction.Amount,
		&staffTransaction.Text,
	)
	if err != nil {
		fmt.Println("error scan", err.Error())
	}
	return staffTransaction, errors.New("not found")
}

func (b *staffTransactionRepo) GetAllStaffTransaction(ctx context.Context, req models.GetAllStaffTransactionRequest) (resp models.GetAllStaffTransaction, err error) {
	var (
		params  = make(map[string]interface{})
		filter  = "WHERE true "
		offsetQ = " OFFSET 0 "
		limit   = " LIMIT 10 "
		offset  = (req.Page - 1) * req.Limit
	)
	s := `
	SELECT
		sale_id,
		staff_id,
		type,
		source_type,
		amount,
		text
	FROM 
		transactions
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
		var staffTransaction models.StaffTransaction
		err := rows.Scan()
		if err != nil {
			return resp, err
		}
		resp.StaffTransactions = append(resp.StaffTransactions, staffTransaction)
		resp.Count = len(resp.StaffTransactions)
	}
	return resp, nil
}

func (s *staffTransactionRepo) DeleteStaffTransaction(ctx context.Context, req models.IdRequestStaffTransaction) (string, error) {
	query := `
	DELETE 
		transactions
	WHERE 
		id=$1 `
	resp, err := s.db.Exec(ctx, query,
		req.Id,
	)
	if err != nil {
		return "error exec", err
	}
	if resp.RowsAffected() == 0 {
		return "error RowsAffected", pgx.ErrNoRows
	}
	return "deleted", nil
}
