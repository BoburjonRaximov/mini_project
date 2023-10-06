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

type staffTariffRepo struct {
	db *pgxpool.Pool
}

func NewStaffTariffRepo(db *pgxpool.Pool) *staffTariffRepo {
	return &staffTariffRepo{db: db}
}

func (s *staffTariffRepo) CreateStaffTariff(ctx context.Context,req models.CreateStaffTariff) (string, error) {
	id := uuid.NewString()
	query :=`
	INSERT INTO 
		tariffs(
			id,
			name,
			type,
			amount_for_cash,
			amount_for_card,
			created_at,
			founded_at) 
	VALUES($1,$2,$3,$4,$5,$6,$7)`
	_, err := s.db.Exec(ctx, query,
		id,
		req.Name,
		req.Type,
		req.AmountForCash,
		req.AmountForCard,
		req.FoundedAt,
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return "", err
	}
	return id, nil
}

func (s *staffTariffRepo) UpdateStaffTariff(ctx context.Context,req models.StaffTariff) (string, error) {
	query := `
	UPDATE 
		tariffs
	 SET
		name=$2,
		type=$3,
		amount_for_cash,
		amount_for_card=$5,
		created_at=$6,
		founded_at=$7
	WHERE 
		id=$1`
	resp, err := s.db.Exec(ctx, query,
		req.Id,
		req.Name,
		req.Type,
		req.AmountForCash,
		req.AmountForCard,
		req.CreatedAt,
		req.FoundedAt,
	)
	if err != nil {
		return "ERROR EXEC", err
	}
	if resp.RowsAffected() == 0 {
		return "error RowsAffected", pgx.ErrNoRows
	}
	return "updated", nil
}

func (s *staffTariffRepo) GetStaffTariff(ctx context.Context,req models.IdRequestStaffTariff) (models.StaffTariff, error) {

	query := `
	SETECT
		name,
		type,
		amount_for_cash,
		amount_for_card,
		created_at::text,
		founded_at::text
	FROM 
		tariffs
	WHERE
		id=$1`
	staffTariff := models.StaffTariff{}
	err := s.db.QueryRow(ctx, query, req.Id).Scan(
		&staffTariff.Id,
		&staffTariff.Name,
		&staffTariff.Type,
		&staffTariff.AmountForCash,
		&staffTariff.AmountForCard,
		&staffTariff.CreatedAt,
		&staffTariff.FoundedAt,
	)
	if err != nil {
		fmt.Println("error scan", err.Error())
	}
	return staffTariff, errors.New("not found")
}

func (st *staffTariffRepo) GetAllStaffTariff(ctx context.Context,req models.GetAllStaffTariffRequest) (resp models.GetAllStaffTariff, err error) {
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
		type,
		amount_for_cash,
		amount_for_card,
		created_at::text,
		founded_at::text
	FROM 
		tariffs
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

	rows, err := st.db.Query(ctx, q, pArr...)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var staffTariff models.StaffTariff
		err := rows.Scan()
		if err != nil {
			return resp, err
		}
		resp.StaffTariffs = append(resp.StaffTariffs, staffTariff)
		resp.Count = len(resp.StaffTariffs)
	}
	return resp, nil
}

func (s *staffTariffRepo) DeleteStaffTariff(ctx context.Context,req models.IdRequestStaffTariff) (string, error) {
	query := `
	DELETE FROM 
		tariffs
	WHERE 
		id=$1 `
	resp, err := s.db.Exec(ctx, query,
		req.Id,
	)
	if err != nil {
		return "error exec", err
	}
	if resp.RowsAffected() == 0 {
		return "RowsAffected", pgx.ErrNoRows
	}

	return "deleted", nil
}
