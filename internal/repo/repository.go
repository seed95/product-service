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
		EditProduct(editedProduct schema.Product) (*schema.Product, error)
		CarpetRepo
		ThemeRepo
		DimensionRepo
	}

	CarpetRepo interface {
		GetAllCarpet(companyId uint) ([]model.Carpet, error)
		GetAllCarpetWithProductId(companyId, productId uint) ([]model.Carpet, error)
	}

	ThemeRepo interface {
		AddThemesWithColor(productId uint, colors []string) ([]schema.Theme, error)
		DeleteThemesWithColor(productId uint, colors []string) error
		DeleteThemesWithId(productId uint, themeIds []uint) error
		UpdateThemesWithId(productId uint, themes []schema.Theme) error
		GetThemesWithProductId(productId uint) ([]schema.Theme, error)
		EditThemesWithId(productId uint, editedThemes []schema.Theme) error
	}

	DimensionRepo interface {
		AddDimensionsToProduct(productId uint, sizes []string) ([]schema.Dimension, error)
		DeleteDimensionsInProduct(productId uint, sizes []string) error
		UpdateDimensionsWithId(dimensions []schema.Dimension) error
		GetDimensionsWithProductId(productId uint) ([]schema.Dimension, error)
	}
)
