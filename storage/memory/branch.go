package memory

import (
	"context"
	"fmt"
	"new_project/models"
	"new_project/pkg/helper"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type branchRepo struct {
	db *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) *branchRepo {
	return &branchRepo{db: db}
}

func (b *branchRepo) CreateBranch(req models.CreateBranch) (string, error) {
	fmt.Println("branch create")
	id := uuid.NewString()

	query := `
	INSERT INTO 
		branch(
			id,
			name,
			address,
			phone_number) 
	VALUES($1,$2,$3,$4)`
	_, err := b.db.Exec(context.Background(), query,
		id,
		req.Name,
		req.Address,
		req.PhoneNumber,
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return "error exec", err
	}
	return id, nil
}

func (b *branchRepo) UpdateBranch(req models.Branch) (string, error) {
	query := `
	UPDATE branch
	 SET 
	 	name=$2,
	 	address=$3,
		phone_number=$4,
		updated_at = NOW()
	WHERE 
		id=$1
	`
	resp, err := b.db.Exec(context.Background(), query,
		req.Id,
		req.Name,
		req.Address,
		req.PhoneNumber,
	)
	if err != nil {
		return "error exec", err
	}
	if resp.RowsAffected() == 0 {
		return "error row", pgx.ErrNoRows
	}
	return "Updated", nil
}

func (b *branchRepo) GetBranch(req models.IdRequest) (models.Branch, error) {
	query := `
	SELECT
	    id,
		name,
		address,
		phone_number,
		created_at::text,
		updated_at::text 
	 FROM 
	 	branch
	WHERE 
		id = $1`
	resp := b.db.QueryRow(context.Background(), query,
		req.Id,
	)
	var branch models.Branch
	err := resp.Scan(
		&branch.Id,
		&branch.Name,
		&branch.Address,
		&branch.PhoneNumber,
		&branch.CreatedAt,
		&branch.UpdatedAt,
	)
	if err != nil {
		fmt.Println("error scan", err.Error())
	}
	return branch, nil
}
func (b *branchRepo) GetAllBranch(req models.GetAllBranchRequest) (resp models.GetAllBranch, err error) {
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
		address,
		phone_number,
		created_at::text,
		updated_at::text 
	FROM 
		branch
	`
	if req.Search != "" {
		filter += ` AND name ILIKE '%' || @search || '%' `
		params["search"] = req.Search
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf("LIMIT %d", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf("OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ
	fmt.Println(query)

	q, pArr := helper.ReplaceQueryParams(query, params)

	rows, err := b.db.Query(context.Background(), q, pArr...)
	if err != nil {
		fmt.Println("error")
		return resp, err

	}
	defer rows.Close()
	for rows.Next() {
		var branch models.Branch
		err := rows.Scan(
			&branch.Id,
			&branch.Name,
			&branch.Address,
			&branch.PhoneNumber,
			&branch.CreatedAt,
			&branch.UpdatedAt,
		)
		if err != nil {
			fmt.Println("err")
			return resp, err
		}
		resp.Branches = append(resp.Branches, branch)
		resp.Count = len(resp.Branches)
		fmt.Println("ok")
	}
	return resp, nil
}

// delete branch
func (b *branchRepo) DeleteBranch(req models.IdRequest) (string, error) {
	query := `
	DELETE FROM 
		branch
	WHERE 
		id=$1 `
	resp, err := b.db.Exec(context.Background(), query,
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
