package models

type CreateBranch struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type Branch struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type IdRequest struct {
	Id string `json:"id"`
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
