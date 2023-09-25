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

type remainingRepo struct {
	db *pgxpool.Pool
}

func NewRemainingRepo(db *pgxpool.Pool) *remainingRepo {
	return &remainingRepo{db: db}
}

func (s *remainingRepo) CreateRemaining(req models.CreateRemaining) (string, error) {
	id := uuid.NewString()
	query := `
	INSERT INTO 
		remaining(
			id,
			branch_id,
			category_id,
			name,
			price,
			barcode,
			count,
			total_price) 
	VALUES($1,$2,$3,$4,$5,$6,$7,$8)
	`
	_, err := s.db.Exec(context.Background(), query,
		id,
		req.BranchId,
		req.CategoryId,
		req.Name,
		req.Price,
		req.Barcode,
		req.Count,
		req.TotalPrice,
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return "error exec", err
	}
	return id, nil
}

func (s *remainingRepo) UpdateRemaining(req models.Remaining) (string, error) {
	query := `
	UPDATE
		remmaining
	SET 
		branch_id=$2,
		category_id=$3,
		name=$4,
		price=$5,
		barcode=$6,
		count=$7,
		total_price=$8,
		updated_at=NOW()
	WHERE 
		id=$1
	`
	resp, err := s.db.Exec(context.Background(), query,
		req.Id,
		req.BranchId,
		req.CategoryId,
		req.Name,
		req.Price,
		req.Barcode,
		req.Count,
		req.TotalPrice,
	)
	if err != nil {
		return "error exec", err
	}
	if resp.RowsAffected() == 0 {
		return "error RowAffected", pgx.ErrNoRows
	}
	return "Updated", nil
}

func (s *remainingRepo) GetRemaining(req models.IdRequestRemaining) (models.Remaining, error) {
	query := `
	SELECT
		id,
		branch_id,
		category_id,
		name,
		price,
		barcode,
		count,
		total_price,
		created_at :: text,
		updated_at :: text
	FROM 
		remaining
	WHERE 
		id=$1
	`
	remaining := models.Remaining{}
	err := s.db.QueryRow(context.Background(), query, req.Id).Scan(
		&remaining.Id,
		&remaining.BranchId,
		&remaining.CategoryId,
		&remaining.Name,
		&remaining.Price,
		&remaining.Barcode,
		&remaining.Count,
		&remaining.TotalPrice,
	)
	if err != nil {
		fmt.Println("error scan", err.Error())
	}
	return remaining, errors.New("not found")
}
func (b *remainingRepo) GetAllRemaining(req models.GetAllRemainingRequest) (resp models.GetAllRemaining, err error) {
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
		branch_id,
		category_id,
		name,
		price,
		barcode,
		count,
		total_price,
		created_at :: text,
		updated_at :: text
	FROM
		remaining
	`

	if req.Branch != "" {
		filter += ` AND branch_id IN(
			SELECT
				r.branch_id
			FROM 
				remaining as r
			JOIN 
				branch as b ON b.id = r.branch_id
			WHERE
				b.name ILIKE '%' || @branch || '%'
		)`
		params["branch"] = req.Branch
	}
	if req.Category != "" {
		filter += ` AND category_id IN(
			SELECT
				r.category_id
			FROM 
				remaining as r
			JOIN 
				category as c ON c.id = r.category_id
			WHERE
				c.name ILIKE '%' || @category || '%'
			) `
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

	rows, err := b.db.Query(context.Background(), q, pArr...)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var remaining models.Remaining
		err := rows.Scan()
		if err != nil {
			return resp, err
		}
		resp.Remainings = append(resp.Remainings, remaining)
		resp.Count = len(resp.Remainings)
	}
	return resp, nil
}
func (s *remainingRepo) DeleteRemaining(req models.IdRequestRemaining) (string, error) {
	query := `
	DELETE FROM
		remaining
	WHERE
		id=$1 
	`
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
