package storage

import "new_project/models"

type StorageI interface {
	Branch() BranchesI
	Staff() StaffsI
	StaffTransaction() StaffTransactionI
	Sale() SalesI
	StaffTariff() StaffTariffsI
}

type BranchesI interface {
	//CreateBranch method creates new branch with given name and address and returns its id
	CreateBranch(models.CreateBranch) (string, error)
	UpdateBranch(models.Branch) (string, error)
	GetBranch(models.IdRequest) (models.Branch, error)
	GetAllBranch(models.GetAllBranchRequest) (models.GetAllBranch, error)
	DeleteBranch(models.IdRequest) (string, error)
}

type StaffsI interface {
	//CreateBranch method creates new branch with given name and address and returns its id
	CreateStaff(models.CreateStaff) (string, error)
	UpdateStaff(models.Staff) (string, error)
	GetStaff(models.IdRequestStaff) (models.Staff, error)
	GetAllStaff(models.GetAllStaffRequest) (models.GetAllStaff, error)
	DeleteStaff(models.IdRequestStaff) (string, error)
}
type StaffTariffsI interface {
	//CreateBranch method creates new branch with given name and address and returns its id
	CreateStaffTariff(models.CreateStaffTariff) (string, error)
	UpdateStaffTariff(models.StaffTariff) (string, error)
	GetStaffTariff(models.IdRequestStaffTariff) (models.StaffTariff, error)
	GetAllStaffTariff(models.GetAllStaffTariffRequest) (models.GetAllStaffTariff, error)
	DeleteStaffTariff(models.IdRequestStaffTariff) (string, error)
}
type StaffTransactionI interface {
	//CreateBranch method creates new branch with given name and address and returns its id
	CreateStaffTransaction(models.CreateStaffTransaction) (string, error)
	UpdateStaffTransaction(models.StaffTransaction) (string, error)
	GetStaffTransaction(models.IdRequestStaffTransaction) (models.StaffTransaction, error)
	GetAllStaffTransaction(models.GetAllStaffTransactionRequest) (models.GetAllStaffTransaction, error)
	DeleteStaffTransaction(models.IdRequestStaffTransaction) (string, error)
}

type SalesI interface {
	//CreateBranch method creates new branch with given name and address and returns its id
	CreateSale(models.CreateSale) (string, error)
	UpdateSale(models.Sale) (string, error)
	GetSale(models.IdRequestSale) (models.Sale, error)
	GetAllSale(models.GetAllSaleRequest) (models.GetAllSale, error)
	DeleteSale(models.IdRequestSale) (string, error)
}
