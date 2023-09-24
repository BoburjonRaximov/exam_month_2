package models

type CreateCategory struct {
	Name      string `json:"name"`
	ParentId  string `json:"parent_id"`
}

type Category struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	ParentId  string `json:"parent_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetAllCategoryRequest struct {
	Page            int
	Limit           int
	Search          string
}

type GetAllCategory struct {
	Categories []Category
	Count int
}

type IdRequestCategory struct {
	Id string `json:"id"`
}
