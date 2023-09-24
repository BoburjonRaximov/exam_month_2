package storage

import "new_project/models"

type StorageI interface {
	Branch() BranchesI
	Category() CategoriesI
	ComingTable() ComingTablesI
	ComingTableProduct() ComingTableProductsI
	Product() ProductsI
	Remaining() RemainingsI
}

type BranchesI interface {
	CreateBranch(models.CreateBranch) (string, error)
	UpdateBranch(models.Branch) (string, error)
	GetBranch(models.IdRequest) (models.Branch, error)
	GetAllBranch(models.GetAllBranchRequest) (models.GetAllBranch, error)
	DeleteBranch(models.IdRequest) (string, error)
}

type CategoriesI interface {
	CreateCategory(models.CreateCategory) (string, error)
	UpdateCategory(models.Category) (string, error)
	GetCategory(models.IdRequestCategory) (models.Category, error)
	GetAllCategory(models.GetAllCategoryRequest) (models.GetAllCategory, error)
	DeleteCategory(models.IdRequestCategory) (string, error)
}
type ComingTablesI interface {
	CreateComingTable(models.CreateComingTable) (string, error)
	UpdateComingTable(models.ComingTable) (string, error)
	GetComingTable(models.IdRequestComingTable) (models.ComingTable, error)
	GetAllComingTable(models.GetAllComingTableRequest) (models.GetAllComingTable, error)
	DeleteComingTable(models.IdRequestComingTable) (string, error)
}
type ComingTableProductsI interface {
	//
	CreateComingTableProduct(models.CreateComingTableProduct) (string, error)
	UpdateComingTableProduct(models.ComingTableProduct) (string, error)
	GetComingTableProduct(models.IdRequestComingTableProduct) (models.ComingTableProduct, error)
	GetAllComingTableProduct(models.GetAllComingTableProductRequest) (models.GetAllComingTableProduct, error)
	DeleteComingTableProduct(models.IdRequestComingTableProduct) (string, error)
}

type ProductsI interface {
	CreateProduct(models.CreateProduct) (string, error)
	UpdateProduct(models.Product) (string, error)
	GetProduct(models.IdRequestProduct) (models.Product, error)
	GetAllProduct(models.GetAllProductRequest) (models.GetAllProduct, error)
	DeleteProduct(models.IdRequestProduct) (string, error)
}

type RemainingsI interface {
	CreateRemaining(models.CreateRemaining) (string, error)
	UpdateRemaining(models.Remaining) (string, error)
	GetRemaining(models.IdRequestRemaining) (models.Remaining, error)
	GetAllRemaining(models.GetAllRemainingRequest) (models.GetAllRemaining, error)
	DeleteRemaining(models.IdRequestRemaining) (string, error)
}
