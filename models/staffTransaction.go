package models

type CreateStaffTransaction struct {
	SaleId     string
	StaffId    string
	Type       int    // withdraw, topup
	SourceType string //sales, bonus
	Amount     float64
	Text       string
}

type StaffTransaction struct {
	Id         string
	SaleId     string
	StaffId    string
	Type       int    // withdraw, topup
	SourceType string //sales, bonus
	Amount     float64
	Text       string
	CreatedAt  string
}

type GetAllStaffTransactionRequest struct {
	Page  int
	Limit int
	Search string
}

type GetAllStaffTransaction struct {
	StaffTransactions []StaffTransaction
	Count             int
}

type CalculateReq struct {
	DateFrom string
	DateTo   string
}
 type IdRequestStaffTransaction struct{
	Id string
 }