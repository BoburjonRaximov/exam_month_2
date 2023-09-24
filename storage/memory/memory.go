package memory

import (
	"context"
	"fmt"
	"new_project/config"
	"new_project/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type strg struct {
	db                     *pgxpool.Pool
	branch                 *branchRepo
	category               *categoryRepo
	comingTable            *comingTableRepo
	comingTableProduct     *comingTableProductRepo
	product                *productRepo
	remaining              *remainingRepo
}

// Branch implements storage.StorageI.
func (b *strg) Branch() storage.BranchesI {
	if b.branch == nil {
		b.branch = NewBranchRepo(b.db)
	}
	return b.branch
}

// Category implements storage.StorageI.
func (b *strg) Category() storage.CategoriesI {
	if b.category == nil {
		b.category = NewCategoryRepo(b.db)
	}
	return b.category
}

// ComingTable implements storage.StorageI.
func (b *strg) ComingTable() storage.ComingTablesI {
	if b.comingTable == nil {
		b.comingTable = NewComingTableRepo(b.db)
	}
	return b.comingTable
}

// ComingTableProduct implements storage.StorageI.
func (b *strg) ComingTableProduct() storage.ComingTableProductsI {
	if b.comingTableProduct == nil {
		b.comingTableProduct = NewComingTableProductRepo(b.db)
	}
	return b.comingTableProduct
}

// Product implements storage.StorageI.
func (b *strg) Product() storage.ProductsI {
	if b.product == nil {
		b.product = NewProductRepo(b.db)
	}
	return b.product
}

// Remaining implements storage.StorageI.
func (b *strg) Remaining() storage.RemainingsI {
	if b.remaining ==nil {
		b.remaining = NewRemainingRepo(b.db)
	}
	return b.remaining
}

func NewStorage(ctx context.Context, cfg config.ConfigPostgres) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.PostgresUser,
			cfg.PostgresPassword,
			cfg.PostgresHost,
			cfg.PostgresPort,
			cfg.PostgresDatabase,
		),
	)
	if err != nil {
		fmt.Println("ParseConfig:", err.Error())
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections
	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		fmt.Println("ConnectConfig:", err.Error())
		return nil, err
	}
	return &strg{
		db: pool,
	}, nil
}
