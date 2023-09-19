package models

type CreateBranch struct {
	Name    string
	Address string
}

type Branch struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type IdRequest struct {
	Id string
}

type GetAllBranchRequest struct {
	Page    int
	Limit   int
	Search  string
	Address string
}
type GetAllBranch struct {
	Branches []Branch
	Count    int
}
