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

type staffTarifRepo struct {
	db *pgxpool.Pool
}

func NewStaffTarifRepo(db *pgxpool.Pool) *staffTarifRepo {
	return &staffTarifRepo{db: db}
}

func (s *staffTarifRepo) CreateStaffTarif(req models.CreateStaffTarif) (string, error) {
	id := uuid.NewString()
	query :=
		`INSERT INTO 
	staffTarifs(id,name,type,amountForCash,amountForCard,createdAt,foundedAt) 
VALUES($1,$2,$3,$4,$5,$6,$7)`
	_, err := s.db.Exec(context.Background(), query,
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

func (s *staffTarifRepo) UpdateStaffTarif(req models.StaffTarif) (string, error) {
	query := `
	update staffTarifs
	set name=$2,type=$3,amountForCash,AmountForCard=$5,createdAt=$6,foundedAt=$7
	where id=$1`
	resp, err := s.db.Exec(context.Background(), query,
		req.Id,
		req.Name,
		req.Type,
		req.AmountForCash,
		req.AmountForCard,
		req.CreatedAt,
		req.FoundedAt,
	)
	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}
	return "updated", nil
}
func (s *staffTarifRepo) GetStaffTarif(req models.IdRequestStaffTarif) (models.StaffTarif, error) {

	query := `
	select * from staffTarif
	where id=$1`
	staffTarif := models.StaffTarif{}
	err := s.db.QueryRow(context.Background(), query, req.Id).Scan(
		&staffTarif.Id,
		&staffTarif.Name,
		&staffTarif.Type,
		&staffTarif.AmountForCash,
		&staffTarif.AmountForCard,
		&staffTarif.CreatedAt,
		&staffTarif.FoundedAt,
	)
	if err != nil {
		fmt.Println("error scan", err.Error())
	}
	return staffTarif, errors.New("not found")
}

func (st *staffTarifRepo) GetAllStaffTarif(req models.GetAllStaffTarifRequest) (resp models.GetAllStaffTarif, err error) {
	var (
		params  = make(map[string]interface{})
		filter  = "WHERE true "
		offsetQ = " OFFSET 0 "
		limit   = " LIMIT 10 "
		offset  = (req.Page - 1) * req.Limit
	)
	s := `
	SELECT *
	FROM staffTarifs
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

	rows, err := st.db.Query(context.Background(), q, pArr...)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var staffTarif models.StaffTarif
		err := rows.Scan()
		if err != nil {
			return resp, err
		}
		resp.StaffTarifs = append(resp.StaffTarifs, staffTarif)
		resp.Count = len(resp.StaffTarifs)
	}
	return resp, nil
}

func (s *staffTarifRepo) DeleteStaffTarif(req models.IdRequestStaffTarif) (string, error) {
	query := `
	delete from staffTarifs
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
