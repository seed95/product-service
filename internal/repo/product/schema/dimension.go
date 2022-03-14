package schema

import (
	"github.com/seed95/product-service/internal/model"
	"gorm.io/gorm"
)

type (
	Dimension struct {
		gorm.Model
		ProductId uint   `gorm:"uniqueIndex:dimension_unique_id"`
		Size      string `gorm:"uniqueIndex:dimension_unique_id"`
	}
)

func GetDimensionsFromProductModel(product model.Product) []Dimension {
	result := make([]Dimension, len(product.Dimensions))
	for i, d := range product.Dimensions {
		result[i] = Dimension{
			ProductId: product.Id,
			Size:      d,
		}
	}
	return result
}
