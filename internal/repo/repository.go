package repo

import (
	"github.com/seed95/product-service/internal/model"
	"github.com/seed95/product-service/internal/repo/product/schema"
)

type (
	ProductRepo interface {
		CreateProduct(product model.Product) (*schema.Product, error)
		GetProductWithId(productId uint) (*schema.Product, error)
		DeleteProduct(productId uint) error
		EditProduct(product model.Product) (*schema.Product, error)
		GetAllProducts(companyId uint) ([]schema.Product, error)
		CarpetRepo
	}

	CarpetRepo interface {
		GetAllCarpet(companyId uint) ([]model.Carpet, error)
		GetAllCarpetWithProductId(companyId, productId uint) ([]model.Carpet, error)
	}
)
