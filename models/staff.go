package models

type CreateStaff struct {
	Name      string
	BranchId  string
	TarifId   string
	Type      int
	Balance   float64
	BirthDate string
}

type Staff struct {
	Id        string
	Name      string
	BranchId  string
	TarifId   string
	Type      int // cashier, shopAssistant
	Balance   float64
	CreatedAt string
	BirthDate string
}

type GetAllStaffRequest struct {
	Page   int
	Limit  int
	Search string
	Name   string
}

type GetAllStaff struct {
	Staffs []Staff
	Count  int
}
type IdRequestStaff struct {
	Id string
}

type ChangeBalanceStaff struct {
	Id     string
	Amount float64
}

type UpdateBalanceRequest struct {
	SaleId          string
	TransactionType string
	SourceType      string
	Cashier         StaffIdAmount
	ShopAssisstant  StaffIdAmount
	Text            string
}

type StaffIdAmount struct {
	StaffId string
	Amount  float32
}
