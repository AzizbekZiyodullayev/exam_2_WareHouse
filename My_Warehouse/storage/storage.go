package storage

import (
	"context"
	"warehouse/api/models"
)

type StorageI interface {
	Close()
	Product() ProductI
	Category() CategoryI
	Branch() BranchI
	Coming() ComingI
	Remainder() RemainderI
	ComingProducts() ComingProductsI
}

type CategoryI interface {
	Create(context.Context, *models.CreateCategory) (string, error)
	GetByID(context.Context, *models.CategoryPrimaryKey) (*models.Category, error)
	// Search by name
	GetList(context.Context, *models.CategoryGetListRequest) (*models.CategoryGetListResponse, error)
	Update(context.Context, *models.UpdateCategory) (int64, error)
	Delete(context.Context, *models.CategoryPrimaryKey) error
}

type ProductI interface {
	Create(context.Context, *models.CreateProduct) (string, error)
	GetByID(context.Context, *models.ProductPrimaryKey) (*models.Product, error)
	// Search by name,barcode
	GetList(context.Context, *models.ProductGetListRequest) (*models.ProductGetListResponse, error)
	Update(context.Context, *models.UpdateProduct) (int64, error)
	Delete(context.Context, *models.ProductPrimaryKey) error
}

type BranchI interface {
	Create(context.Context, *models.CreateBranch) (string, error)
	GetByID(context.Context, *models.BranchPrimaryKey) (*models.Branch, error)
	// Search by name
	GetList(context.Context, *models.BranchGetListRequest) (*models.BranchGetListResponse, error)
	Update(context.Context, *models.UpdateBranch) (int64, error)
	Delete(context.Context, *models.BranchPrimaryKey) error
}

type ComingI interface {
	Create(context.Context, *models.CreateComing) (string, error)
	GetByID(context.Context, *models.ComingPrimaryKey) (*models.Coming, error)
	// Search by comingId,branch;
	GetList(context.Context, *models.ComingGetListRequest) (*models.ComingGetListResponse, error)
	Update(context.Context, *models.UpdateComing) (int64, error)
	Delete(context.Context, *models.ComingPrimaryKey) error
}

type RemainderI interface {
	Create(context.Context, *models.CreateRemainder) (string, error)
	GetByID(context.Context, *models.RemainderPrimaryKey) (*models.Remainder, error)
	// Search by branch,category,barcode;
	GetList(context.Context, *models.RemainderGetListRequest) (*models.RemainderGetListResponse, error)
	Update(context.Context, *models.UpdateRemainder) (int64, error)
	Delete(context.Context, *models.RemainderPrimaryKey) error
	DoIncomeWareHouse(ctx context.Context, req *models.RespProduct) (int64, error)
}

type ComingProductsI interface {
	Create(context.Context, *models.CreateComingProducts) (string, error)
	GetByID(context.Context, *models.ComingProductsPrimaryKey) (*models.ComingProducts, error)
	// Search by category,barcode;
	GetList(context.Context, *models.ComingProductsGetListRequest) (*models.ComingProductsGetListResponse, error)
	Update(context.Context, *models.UpdateComingProducts) (int64, error)
	Delete(context.Context, *models.ComingProductsPrimaryKey) error
	GetByComingID(ctx context.Context, req *models.ComingPrimaryKey) (*models.ComingProducts, error)
}
