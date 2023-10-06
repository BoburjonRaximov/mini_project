package models

type CreateStaff struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Name     string  `json:"name"`
	BranchId string  `json:"branch_id"`
	TarifId  string  `json:"tarif_id"`
	Type     int     `json:"type"`
	Balance  float64 `json:"balance"`
}

type Staff struct {
	Id        string  `json:"id"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	Name      string  `json:"name"`
	BranchId  string  `json:"branch_id"`
	TarifId   string  `json:"tarif_id"`
	Type      int     `json:"type"`
	Balance   float64 `json:"balance"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
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
	Id string `json:"id"`
}

// Login
type RequestByUsername struct {
	Username string
}
type LoginRes struct {
	Token string `json:"token"`
}

// it's no use
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
