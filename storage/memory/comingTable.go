package memory

import (
	"context"
	"errors"
	"fmt"
	"new_project/models"
	"new_project/pkg/helper"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type comingTableRepo struct {
	db *pgxpool.Pool
}

func NewComingTableRepo(db *pgxpool.Pool) *comingTableRepo {
	return &comingTableRepo{db: db}
}

func (s *comingTableRepo) CreateComingTable(req models.CreateComingTable) (string, error) {
	fmt.Println("coming table create")
	id := uuid.NewString()
	query := `
	INSERT INTO 
	 coming_table(
		id,
		coming_id,
		branch_id,
		date_time,
		status) 
     VALUES($1,$2,$3,$4,$5)`
	_, err := s.db.Exec(context.Background(), query,
		id,
		req.ComingId,
		req.BranchId,
		req.DateTime,
		req.Status,
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return "error exec", err
	}
	return id, nil
}

func (s *comingTableRepo) UpdateComingTable(req models.ComingTable) (string, error) {
	query := `
	UPDATE 
		coming_table
	 SET 
		coming_id=$2,
		branch_id=$3,
		date_time=$4,
		status=$5,
		updated_at= NOW()
	 WHERE id=$1
	 `
	resp, err := s.db.Exec(context.Background(), query,
		req.Id,
		req.ComingId,
		req.BranchId,
		req.DateTime,
		req.Status,
	)
	if err != nil {
		return "error exec", err
	}
	if resp.RowsAffected() == 0 {
		return "error RowAffected", pgx.ErrNoRows
	}
	return "Updated", nil
}

func (s *comingTableRepo) GetComingTable(req models.IdRequestComingTable) (models.ComingTable, error) {
	query := `
	SELECT
		id,
		coming_id,
		branch_id,
		date_time,
		status,
		created_at :: text,
		updated_at :: text
	 FROM
		coming_table
	 WHERE 
		id=$1
		`
	comingTable := models.ComingTable{}
	err := s.db.QueryRow(context.Background(), query, req.Id).Scan(
		&comingTable.Id,
		&comingTable.ComingId,
		&comingTable.BranchId,
		&comingTable.DateTime,
		&comingTable.Status,
		&comingTable.CreatedAt,
		&comingTable.UpdatedAt,
	)
	if err != nil {
		fmt.Println("error scan", err.Error())
	}
	return comingTable, errors.New("not found")
}

func (b *comingTableRepo) GetAllComingTable(req models.GetAllComingTableRequest) (resp models.GetAllComingTable, err error) {
	var (
		params  = make(map[string]interface{})
		filter  = "	WHERE true "
		offsetQ = " OFFSET 0 "
		limit   = " LIMIT 10 "
		offset  = (req.Page - 1) * req.Limit
	)
	s := `
	SELECT
		id,
		coming_id,
		branch_id,
		date_time,
		status,
		created_at :: text,
		updated_at :: text
	FROM
		coming_table
	`

	if req.ComingId != "" {
		filter += ` AND coming_id=@coming_id `
		params["coming_id"] = req.ComingId
	}
	if req.Branch != "" {
		filter += ` AND branch_id IN(
			SELECT
				ct.branch_id
			FROM 
				coming_table as ct
			JOIN 
				branch as b ON b.id = ct.branch_id
			WHERE
				b.name ILIKE '%' || @branch || '%'
		)`
		params["branch"] = req.Branch
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
		var comingTable models.ComingTable
		err := rows.Scan()
		if err != nil {
			return resp, err
		}
		resp.ComingTables = append(resp.ComingTables, comingTable)
		resp.Count = len(resp.ComingTables)
	}
	return resp, nil
}

func (s *comingTableRepo) DeleteComingTable(req models.IdRequestComingTable) (string, error) {
	query := `
	DELETE FROM 
		coming_table
	WHERE
		id=$1 `
	resp, err := s.db.Exec(context.Background(), query,
		req.Id,
	)
	if err != nil {
		return "ERROR EXEC", err
	}
	if resp.RowsAffected() == 0 {
		return "error RowsAffected", pgx.ErrNoRows
	}

	return "deleted", nil
}
