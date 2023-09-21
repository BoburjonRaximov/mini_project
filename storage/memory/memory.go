package memory

import (
	"context"
	"fmt"
	"new_project/config"
	"new_project/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type strg struct {
	db               *pgxpool.Pool
	branch           *branchRepo
	staff            *staffRepo
	staffTariff      *staffTariffRepo
	sale             *saleRepo
	staffTransaction *staffTransactionRepo
}

// Branch implements storage.StorageI.
func (b *strg) Branch() storage.BranchesI {
	if b.branch == nil {
		b.branch = NewBranchRepo(b.db)
	}
	return b.branch
}

// Sale implements storage.StorageI.
func (s *strg) Sale() storage.SalesI {
	if s.sale == nil {
		s.sale = NewSaleRepo(s.db)
	}
	return s.sale
}

// Staff implements storage.StorageI.
func (s *strg) Staff() storage.StaffsI {
	if s.staff == nil {
		s.staff = NewStaffRepo(s.db)
	}
	return s.staff
}

// StaffTarif implements storage.StorageI.
func (s *strg) StaffTariff() storage.StaffTariffsI {
	if s.staffTariff == nil {
		s.staffTariff = NewStaffTariffRepo(s.db)
	}
	return s.staffTariff
}

// StaffTransaction implements storage.StorageI.
func (s *strg) StaffTransaction() storage.StaffTransactionI {
	if s.staffTransaction == nil {
		s.staffTransaction = NewStaffTransactionRepo(s.db)
	}
	return s.staffTransaction
}

func NewStorage(ctx context.Context, cfg config.ConfigPostgres) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.PostgresUser,
			cfg.PostgresPassword,
			cfg.PostgresHost,
			cfg.PostgresPort,
			cfg.PostgresDatabase,
		),
	)
	if err != nil {
		fmt.Println("ParseConfig:", err.Error())
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections
	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		fmt.Println("ConnectConfig:", err.Error())
		return nil, err
	}
	return &strg{
		db: pool,
	}, nil
}
