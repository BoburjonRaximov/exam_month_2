package models

type CreateRemaining struct {
	BranchId   string  `json:"branch_id"`
	CategoryId string  `json:"category_id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Barcode    string  `json:"barcode"`
	Count      int     `json:"count"`
	TotalPrice float64 `json:"total_price"`
}

type Remaining struct {
	Id         string  `json:"id"`
	BranchId   string  `json:"branch_id"`
	CategoryId string  `json:"category_id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Barcode    string  `json:"barcode"`
	Count      int     `json:"count"`
	TotalPrice float64 `json:"total_price"`
	CreatedAt  string  `json:"cteated_at"`
	FoundedAt  string  `json:"founded_at"`
}

type GetAllRemainingRequest struct {
	Page       int
	Limit      int
	Search     string
	BranchId   string
	CategoryId string
	Barcode    string
}

type GetAllRemaining struct {
	Remainings []Remaining
	Count      int
}

type IdRequestRemaining struct {
	Id string `json:"id"`
}
