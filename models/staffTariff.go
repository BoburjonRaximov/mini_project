package models

type CreateStaffTariff struct {
	Name          string
	Type          int // fixed, percent
	AmountForCash float64
	AmountForCard float64
	FoundedAt     string
}

type StaffTariff struct {
	Id            string
	Name          string
	Type          int // fixed, percent
	AmountForCash float64
	AmountForCard float64
	CreatedAt     string
	FoundedAt     string
}

type GetAllStaffTariffRequest struct {
	Page   int
	Limit  int
	Search string
	Name   string
}

type GetAllStaffTariff struct {
	StaffTariffs []StaffTariff
	Count       int
}

type IdRequestStaffTariff struct {
	Id string
}
