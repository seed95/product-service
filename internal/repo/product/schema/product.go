package schema

import (
	"github.com/seed95/OrderManagement/Microservice/product-service/internal/model"
	"gorm.io/gorm"
)

type (
	Product struct {
		gorm.Model
		CompanyId   uint   `gorm:"uniqueIndex:product_unique_id"`
		DesignCode  string `gorm:"uniqueIndex:product_unique_id"`
		Description string
		Dimensions  []Dimension
		Themes      []Theme
	}

	Dimension struct {
		gorm.Model
		ProductId uint   `gorm:"uniqueIndex:dimension_unique_id"`
		Size      string `gorm:"uniqueIndex:dimension_unique_id"`
	}

	Theme struct {
		gorm.Model
		ProductId uint   `gorm:"uniqueIndex:theme_unique_id"`
		Color     string `gorm:"uniqueIndex:theme_unique_id"`
	}
)

func GetProduct(p *model.Product) *Product {
	return &Product{
		Model: gorm.Model{
			ID: p.Id,
		},
		CompanyId:   p.CompanyId,
		DesignCode:  p.DesignCode,
		Description: p.Description,
	}
}

func GetDimensions(product *model.Product) []Dimension {

	result := make([]Dimension, len(product.Dimensions))

	for i, d := range product.Dimensions {
		result[i] = Dimension{
			ProductId: product.Id,
			Size:      d,
		}
	}

	return result
}

func GetThemes(product *model.Product) []Theme {

	result := make([]Theme, len(product.Colors))

	for i, c := range product.Colors {
		result[i] = Theme{
			ProductId: product.Id,
			Color:     c,
		}
	}

	return result
}
