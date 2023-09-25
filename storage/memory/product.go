package memory

import (
	"context"
	"errors"
	"fmt"
	"new_project/models"
	"new_project/pkg/helper"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type productRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *productRepo {
	return &productRepo{db: db}
}

func (s *productRepo) CreateProduct(req models.CreateProduct) (string, error) {
	fmt.Println("product create")
	id := uuid.NewString()
	query := `
	INSERT INTO 
	    product(
			id,
			name,
			price,
			barcode,
			category_id) 
    VALUES($1,$2,$3,$4,$5)
	`
	_, err := s.db.Exec(context.Background(), query,
		id,
		req.Name,
		req.Price,
		req.Barcode,
		req.CategoryId,
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return "error exec", err
	}
	return id, nil
}

func (s *productRepo) UpdateProduct(req models.Product) (string, error) {
	query := `
	UPDATE 
		product
	SET 
		name=$2,
		price=$3,
		barcode=$4,
		category_id=$5,
		updated_at = NOW()
	WHERE 
		id=$1
	`
	resp, err := s.db.Exec(context.Background(), query,
		req.Id,
		req.Name,
		req.Price,
		req.Barcode,
		req.CategoryId,
	)
	if err != nil {
		return "error exec", err
	}
	if resp.RowsAffected() == 0 {
		return "error RowsAffected", pgx.ErrNoRows
	}
	return "Updated", nil
}

func (s *productRepo) GetProduct(req models.IdRequestProduct) (models.Product, error) {
	query := `
	SELECT
		id,
		name,
		price,
		barcode,
		category_id,
		created_at :: text,
		updated_at :: text
	FROM
		product
	WHERE
		id=$1
	`
	product := models.Product{}
	err := s.db.QueryRow(context.Background(), query, req.Id).Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.Barcode,
		&product.CategoryId,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		fmt.Println("error scan", err.Error())
	}
	return product, errors.New("not found")
}
func (b *productRepo) GetAllProduct(req models.GetAllProductRequest) (resp models.GetAllProduct, err error) {
	var (
		params  = make(map[string]interface{})
		filter  = " WHERE true "
		offsetQ = " OFFSET 0 "
		limit   = " LIMIT 10 "
		offset  = (req.Page - 1) * req.Limit
	)
	s := `
	SELECT
		id,
		name,
		price,
		barcode,
		category_id,
		created_at :: text,
		updated_at :: text
	FROM
		product
	`
	if req.SearchName != "" {
		filter += ` AND name ILIKE '%' || @search || '%' `
		params["search"] = req.SearchName
	}
	if req.Barcode != "" {
		filter += ` AND barcode=@barcode `
		params["barcode"] = req.Barcode
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf("LIMIT %d", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf("OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ

	q, pArr := helper.ReplaceQueryParams(query, params)

	rows, err := b.db.Query(context.Background(), q, pArr...)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err := rows.Scan()
		if err != nil {
			return resp, err
		}
		resp.Products = append(resp.Products, product)
		resp.Count = len(resp.Products)
	}
	return resp, nil
}
func (s *productRepo) DeleteProduct(req models.IdRequestProduct) (string, error) {
	query := `
	DELLETE FROM
		product
	WHERE
		id=$1 `
	resp, err := s.db.Exec(context.Background(), query,
		req.Id,
	)
	if err != nil {
		return "error exec", err
	}
	if resp.RowsAffected() == 0 {
		return "error RowsAffected", pgx.ErrNoRows
	}

	return "deleted", nil
}
