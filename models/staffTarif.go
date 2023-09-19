package models

type CreateStaffTarif struct {
	Name          string
	Type          int // fixed, percent
	AmountForCash float64
	AmountForCard float64
	FoundedAt     string
}

type StaffTarif struct {
	Id            string
	Name          string
	Type          int // fixed, percent
	AmountForCash float64
	AmountForCard float64
	CreatedAt     string
	FoundedAt     string
}

type GetAllStaffTarifRequest struct {
	Page  int
	Limit int
	Search string
	Name  string
}

type GetAllStaffTarif struct {
	StaffTarifs []StaffTarif
	Count       int
}

type IdRequestStaffTarif struct {
	Id string
}
