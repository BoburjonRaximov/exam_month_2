package models

type CreateProduct struct {
	Name       string  `json:"name"`
	Price      string  `json:"price"`
	Barcode    float64 `json:"barcode"`
	CategoryId string  `json:"category_id"`
}

type Product struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Barcode    string  `json:"barcode"`
	CategoryId string  `json:"category_id"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type GetAllProductRequest struct {
	Page    int
	Limit   int
	SearchName  string
	Barcode string
}

type GetAllProduct struct {
	Products []Product
	Count    int
}
type IdRequestProduct struct {
	Id string `json:"id"`
}
