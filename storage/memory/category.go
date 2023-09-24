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

type categoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *categoryRepo {
	return &categoryRepo{db: db}
}

func (s *categoryRepo) CreateCategory(req models.CreateCategory) (string, error) {
	fmt.Println("category create")
	id := uuid.NewString()
	query := `
	INSERT INTO 
	 category(
		id,
		name,
		parent_id) 
     VALUES($1,$2,$3)`
	_, err := s.db.Exec(context.Background(), query,
		id,
		req.Name,
		req.ParentId,
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return "", err
	}
	return id, nil
}

func (s *categoryRepo) UpdateCategory(req models.Category) (string, error) {
	query := `
	update 
		category
	 set 
	 	name=$2,
		parent_id=$3,
		updated_at=NOW()
	 where 
	 	id=$1`
	resp, err := s.db.Exec(context.Background(), query,
		req.Id,
		req.Name,
		req.ParentId,
	)
	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}
	return "Updated", nil
}

func (s *categoryRepo) GetCategory(req models.IdRequestCategory) (models.Category, error) {
	query := `
	select
		id,
		name,
		parent_id,
		created_at,
		updated_at 
	 from 
	 	category
	 where
		id=$1`
	category := models.Category{}
	err := s.db.QueryRow(context.Background(), query, req.Id).Scan(
		&category.Id,
		&category.Name,
		&category.ParentId,
		&category.CreatedAt,
		&category.UpdatedAt,
	)
	if err != nil {
		fmt.Println("error scan", err.Error())
	}
	return category, errors.New("not found")
}

func (b *categoryRepo) GetAllCategory(req models.GetAllCategoryRequest) (resp models.GetAllCategory, err error) {
	var (
		params  = make(map[string]interface{})
		filter  = "WHERE true "
		offsetQ = " OFFSET 0 "
		limit   = " LIMIT 10 "
		offset  = (req.Page - 1) * req.Limit
	)
	s := `
	SELECT
		id,
		name,
		parent_id,
		created_at,
		updated_at 
	FROM
		category
	`
	if req.Search != "" {
		filter += ` WHERE name ILIKE '%' || @search || '%' `
		params["search"] = req.Search
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
		var category models.Category
		err := rows.Scan()
		if err != nil {
			return resp, err
		}
		resp.Categories = append(resp.Categories, category)
		resp.Count = len(resp.Categories)
	}
	return resp, nil
}
func (s *categoryRepo) DeleteCategory(req models.IdRequestCategory) (string, error) {

	query := `
	delete from
		category
	where
		id=$1 `
	resp, err := s.db.Exec(context.Background(), query,
		req.Id,
	)
	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}

	return "deleted", nil
}
