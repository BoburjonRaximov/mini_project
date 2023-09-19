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

func (s *staffTransactionRepo) CreateStaffTransaction(req models.CreateStaffTransaction) (string, error) {
	fmt.Println("staff create")
	id := uuid.NewString()
	query :=
		`INSERT INTO 
	staffTransactions(id,saleId,staffId,type,sourceType,amount,text,createdAt) 
VALUES($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := s.db.Exec(context.Background(), query,
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

func (s *staffTransactionRepo) UpdateStaffTransaction(req models.StaffTransaction) (string, error) {
	query := `
	update staffTransactions
	set saleId=$2,StaffId=$3,type=$4,sourceType=$5,amount=$6,text=$7
	where id=$1`
	resp, err := s.db.Exec(context.Background(), query,
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

func (s *staffTransactionRepo) GetStaffTransaction(req models.IdRequestStaffTransaction) (models.StaffTransaction, error) {
	query := `
	select * from staffTransactions
	where id=$1`
	staffTransaction := models.StaffTransaction{}
	err := s.db.QueryRow(context.Background(), query, req.Id).Scan(
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
func (b *staffTransactionRepo) GetAllStaffTransaction(req models.GetAllStaffTransactionRequest) (resp models.GetAllStaffTransaction, err error) {
	var (
		params  = make(map[string]interface{})
		filter  = "WHERE true "
		offsetQ = " OFFSET 0 "
		limit   = " LIMIT 10 "
		offset  = (req.Page - 1) * req.Limit
	)
	s := `
	SELECT *
	FROM staffTransactions
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
func (s *staffTransactionRepo) DeleteStaffTransaction(req models.IdRequestStaffTransaction) (string, error) {
	query := `
	delete from staffTransaction
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
