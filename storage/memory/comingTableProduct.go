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

type comingTableProductRepo struct {
	db *pgxpool.Pool
}

func NewComingTableProductRepo(db *pgxpool.Pool) *comingTableProductRepo {
	return &comingTableProductRepo{db: db}
}

func (s *comingTableProductRepo) CreateComingTableProduct(req models.CreateComingTableProduct) (string, error) {
	id := uuid.NewString()
	query := `
	INSERT INTO 
	 coming_table_product(
		id,
		category_id,
		name,
		price,
		barcode,
		count,
		total_price,
		coming_table_id) 
     VALUES($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := s.db.Exec(context.Background(), query,
		id,
		req.CategoryId,
		req.Name,
		req.Price,
		req.Barcode,
		req.Count,
		req.TotalPrice,
		req.ComingTableId,
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return "error exec", err
	}
	return id, nil
}

func (s *comingTableProductRepo) UpdateComingTableProduct(req models.ComingTableProduct) (string, error) {
	query := `
	UPDATE 
		coming_table_product
	SET 
	    category_id=$2,
		name=$3,
		price=$4,
		barcode=$5,
		count=$6,
		total_price=$7,
		coming_table_id=$8,
		updated_at=NOW()
	WHERE 
		id=$1`
	resp, err := s.db.Exec(context.Background(), query,
		req.Id,
		req.CategoryId,
		req.Name,
		req.Price,
		req.Barcode,
		req.Count,
		req.TotalPrice,
		req.ComingTableId,
	)
	if err != nil {
		return "error exec", err
	}
	if resp.RowsAffected() == 0 {
		return "error RowsAffected", pgx.ErrNoRows
	}
	return "updated", nil
}
func (s *comingTableProductRepo) GetComingTableProduct(req models.IdRequestComingTableProduct) (models.ComingTableProduct, error) {

	query := `
	SELECT
		id,
		category_id,
		name,
		price,
		barcode,
		count,
		total_price,
		coming_table_id,
		created_at :: text,
		updated_at :: text
	FROM
		coming_table_product
	WHERE
		id=$1
	`
	comingTableProduct := models.ComingTableProduct{}
	err := s.db.QueryRow(context.Background(), query, req.Id).Scan(
		&comingTableProduct.Id,
		&comingTableProduct.CategoryId,
		&comingTableProduct.Name,
		&comingTableProduct.Price,
		&comingTableProduct.Barcode,
		&comingTableProduct.Count,
		&comingTableProduct.TotalPrice,
		&comingTableProduct.ComingTableId,
		&comingTableProduct.CreatedAt,
		&comingTableProduct.UpdatedAt,
	)
	if err != nil {
		fmt.Println("error scan", err.Error())
	}
	return comingTableProduct, errors.New("not found")
}

func (st *comingTableProductRepo) GetAllComingTableProduct(req models.GetAllComingTableProductRequest) (resp models.GetAllComingTableProduct, err error) {
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
		category_id,
		name,
		price,
		barcode,
		count,
		total_price,
		coming_table_id,
		created_at :: text,
		updated_at :: text
	FROM 
		coming_table_product
	`

	if req.Category != "" {
		filter += ` AND category_id IN(
	SELECT
		ct.category_id
	FROM 
		coming_table_product as ct
	JOIN 
		category as c ON c.id = ct.category_id
	WHERE
		c.name ILIKE '%' || @category || '%'
	)`
		params["category"] = req.Category
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

	rows, err := st.db.Query(context.Background(), q, pArr...)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var comingTableProduct models.ComingTableProduct
		err := rows.Scan()
		if err != nil {
			return resp, err
		}
		resp.ComingTableProducts = append(resp.ComingTableProducts, comingTableProduct)
		resp.Count = len(resp.ComingTableProducts)
	}
	return resp, nil
}

func (s *comingTableProductRepo) DeleteComingTableProduct(req models.IdRequestComingTableProduct) (string, error) {
	query := `
	DELETE FROM
		coming_table_product
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
