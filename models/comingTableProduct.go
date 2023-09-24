package models

type CreateComingTableProduct struct {
	CategoryId    string  `json:"category_id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Barcode       string  `json:"barcode"`
	Count         int     `json:"count"`
	TotalPrice    float64 `json:"total_price"`
	ComingTableId string  `json:"coming_table_id"`
}

type ComingTableProduct struct {
	Id            string  `json:"id"`
	CategoryId    string  `json:"category_id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Barcode       string  `json:"barcode"`
	Count         int     `json:"count"`
	TotalPrice    float64 `json:"total_price"`
	ComingTableId string  `json:"coming_table_id"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

type GetAllComingTableProductRequest struct {
	Page       int
	Limit      int
	Search     string
	CategoryId string
	Barcode    string
}

type GetAllComingTableProduct struct {
	ComingTableProducts []ComingTableProduct
	Count               int
}

type IdRequestComingTableProduct struct {
	Id string `json:"id"`
}
