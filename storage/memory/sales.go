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

type saleRepo struct {
	db *pgxpool.Pool
}

func NewSaleRepo(db *pgxpool.Pool) *saleRepo {
	return &saleRepo{db: db}
}

func (s *saleRepo) CreateSale(ctx context.Context, req models.CreateSale) (string, error) {
	fmt.Println("staff create")
	id := uuid.NewString()
	query :=`
	INSERT INTO 
		sales(
			id,
			branch_id,
			shop_assistant_id,
			cashier_id,
			price,
			payment_type,
			client_name) 
VALUES($1,$2,$3,$4,$5,$6,$7)`
	_, err := s.db.Exec(ctx, query,
		id,
		req.BranchId,
		req.ShopAssistantId,
		req.CashierId,
		req.Price,
		req.PaymentType,
		req.ClientName,
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return "", err
	}
	return id, nil
}

func (s *saleRepo) UpdateSale(ctx context.Context,req models.Sale) (string, error) {
	query := `
	UPDATE 
		sales
	SET 
		branch_id=$2,
		shop_assistant_id=$3,
		cashier_id=$4,
		price=$5,
		payment_type=$6,
		client_name=$7
	where 
		id=$1`
	resp, err := s.db.Exec(ctx, query,
		req.Id,
		req.BranchId,
		req.ShopAssistantId,
		req.CashierId,
		req.Price,
		req.PaymentType,
		req.ClientName,
	)
	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}
	return "Updated", nil
}

func (s *saleRepo) GetSale(ctx context.Context,req models.IdRequestSale) (models.Sale, error) {
	query := `
	SELECT
		branch_id,
		shop_assistant_id,
		cashier_id,
		price,
		payment_type,
		client_name
	FROM 
		sales
	WHERE 
		id=$1`
	sale := models.Sale{}
	err := s.db.QueryRow(ctx, query, req.Id).Scan(
		&sale.Id,
		&sale.BranchId,
		&sale.ShopAssistantId,
		&sale.CashierId,
		&sale.Price,
		&sale.PaymentType,
		&sale.ClientName,
	)
	if err != nil {
		fmt.Println("error scan", err.Error())
	}
	return sale, errors.New("not found")
}

func (b *saleRepo) GetAllSale(ctx context.Context,req models.GetAllSaleRequest) (resp models.GetAllSale, err error) {
	var (
		params  = make(map[string]interface{})
		filter  = "WHERE true "
		offsetQ = " OFFSET 0 "
		limit   = " LIMIT 10 "
		offset  = (req.Page - 1) * req.Limit
	)
	s := `
	SELECT
		branch_id,
		shop_assistant_id,
		cashier_id,
		price,
		payment_type,
		client_name
	FROM 
		sales
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
		var sale models.Sale
		err := rows.Scan()
		if err != nil {
			return resp, err
		}
		resp.Sales = append(resp.Sales, sale)
		resp.Count = len(resp.Sales)
	}
	return resp, nil
}
func (s *saleRepo) DeleteSale(ctx context.Context,req models.IdRequestSale) (string, error) {

	query := `
	DELETE FROM 
		sales
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
