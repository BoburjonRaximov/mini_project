package storage

import (
	"context"
	"new_project/models"
	"time"
)

type StorageI interface {
	Branch() BranchesI
	Staff() StaffsI
	StaffTransaction() StaffTransactionI
	Sale() SalesI
	StaffTariff() StaffTariffsI
}

type CacheI interface {
	Cache() RedisI
}

type BranchesI interface {
	//CreateBranch method creates new branch with given name and address and returns its id
	CreateBranch(context.Context, models.CreateBranch) (string, error)
	UpdateBranch(context.Context, models.Branch) (string, error)
	GetBranch(context.Context, models.IdRequest) (models.Branch, error)
	GetAllBranch(context.Context, models.GetAllBranchRequest) (models.GetAllBranch, error)
	DeleteBranch(context.Context, models.IdRequest) (string, error)
}

type StaffsI interface {
	//CreateBranch method creates new branch with given name and address and returns its id
	CreateStaff(context.Context, models.CreateStaff) (string, error)
	UpdateStaff(context.Context, models.Staff) (string, error)
	GetStaff(context.Context, models.IdRequestStaff) (models.Staff, error)
	GetAllStaff(context.Context, models.GetAllStaffRequest) (models.GetAllStaff, error)
	DeleteStaff(context.Context, models.IdRequestStaff) (string, error)
}
type StaffTariffsI interface {
	//CreateBranch method creates new branch with given name and address and returns its id
	CreateStaffTariff(context.Context, models.CreateStaffTariff) (string, error)
	UpdateStaffTariff(context.Context, models.StaffTariff) (string, error)
	GetStaffTariff(context.Context, models.IdRequestStaffTariff) (models.StaffTariff, error)
	GetAllStaffTariff(context.Context, models.GetAllStaffTariffRequest) (models.GetAllStaffTariff, error)
	DeleteStaffTariff(context.Context, models.IdRequestStaffTariff) (string, error)
}
type StaffTransactionI interface {
	//CreateBranch method creates new branch with given name and address and returns its id
	CreateStaffTransaction(context.Context, models.CreateStaffTransaction) (string, error)
	UpdateStaffTransaction(context.Context, models.StaffTransaction) (string, error)
	GetStaffTransaction(context.Context, models.IdRequestStaffTransaction) (models.StaffTransaction, error)
	GetAllStaffTransaction(context.Context, models.GetAllStaffTransactionRequest) (models.GetAllStaffTransaction, error)
	DeleteStaffTransaction(context.Context, models.IdRequestStaffTransaction) (string, error)
}

type SalesI interface {
	//CreateBranch method creates new branch with given name and address and returns its id
	CreateSale(context.Context, models.CreateSale) (string, error)
	UpdateSale(context.Context, models.Sale) (string, error)
	GetSale(context.Context, models.IdRequestSale) (models.Sale, error)
	GetAllSale(context.Context, models.GetAllSaleRequest) (models.GetAllSale, error)
	DeleteSale(context.Context, models.IdRequestSale) (string, error)
}

type RedisI interface {
	Create(ctx context.Context, key string, obj interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string, res interface{}) (bool, error)
	Delete(ctx context.Context, id string) error
}
