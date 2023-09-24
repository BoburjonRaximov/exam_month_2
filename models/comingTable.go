package models

type CreateComingTable struct {
	ComingId string `json:"coming_id"`
	BranchId string `json:"branch_id"`
	DateTime string `json:"date_time"`
	Status   string `json:"status"`
}

type ComingTable struct {
	Id        string `json:"id"`
	ComingId  string `json:"coming_id"`
	BranchId  string `json:"branch_id"`
	DateTime  string `json:"date_time"`
	Status    string `json:"status"`
	CreatedAt string `json:"cteated_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetAllComingTableRequest struct {
	Page     int
	Limit    int
	Search   string
	ComingId string
	BranchId string
}

type GetAllComingTable struct {
	ComingTables []ComingTable
	Count        int
}

type IdRequestComingTable struct {
	Id string `json:"id"`
}
