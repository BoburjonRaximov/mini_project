package models

type CreateSale struct {
	BranchId        string
	ShopAssistantId string
	CashierId       string
	Price           float64
	PaymentType     int //1-card , chash-2
	ClientName      string
}

type Sale struct {
	Id              string
	BranchId        string
	ShopAssistantId string
	CashierId       string
	Price           float64
	PaymentType     int //1-card, 2-cash
	Status          int //1-success, 2-cancel
	ClientName      string
	CreatedAt       string
}

type GetAllSaleRequest struct {
	Page            int
	Search          string
	Limit           int
	ClientName      string
	BranchId        string
	ShopAssistantId string
	CashierId       string
	PaymentType     int //1-card, 2-cash
	Status          int //1-success, 2-cancel
	PriceFrom       float64
	PriceTo         float64
}

type GetAllSale struct {
	Sales []Sale
	Count int
}

type DailySales struct {
	Day         string
	BranchId    string
	SalesAmount float64
}

type IdRequestSale struct {
	Id string
}
