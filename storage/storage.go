package storage

import "new_project/models"

type StorageI interface {
	Branch() BranchesI
	Staff() StaffsI
	StaffTransaction() StaffTransactionI
	Sale() SalesI
	StaffTarif() StaffTarifsI
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
type StaffTarifsI interface {
	//CreateBranch method creates new branch with given name and address and returns its id
	CreateStaffTarif(models.CreateStaffTarif) (string, error)
	UpdateStaffTarif(models.StaffTarif) (string, error)
	GetStaffTarif(models.IdRequestStaffTarif) (models.StaffTarif, error)
	GetAllStaffTarif(models.GetAllStaffTarifRequest) (models.GetAllStaffTarif, error)
	DeleteStaffTarif(models.IdRequestStaffTarif) (string, error)
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
