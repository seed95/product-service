package schema

import (
	"fmt"
	"github.com/seed95/product-service/internal/model"
	"gorm.io/gorm"
)

type (
	Theme struct {
		gorm.Model
		ProductId uint   `gorm:"uniqueIndex:theme_unique_id"`
		Color     string `gorm:"uniqueIndex:theme_unique_id"`
	}
)

func GetThemesFromProductModel(product model.Product) []Theme {
	result := make([]Theme, len(product.Colors))
	for i, c := range product.Colors {
		result[i] = Theme{
			ProductId: product.Id,
			Color:     c,
		}
	}
	return result
}

func GetThemes(productId uint, colors []string) []Theme {
	result := make([]Theme, len(colors))
	for i, c := range colors {
		result[i] = Theme{
			ProductId: productId,
			Color:     c,
		}
	}
	return result
}

func (t Theme) String() string {
	return fmt.Sprintf("ID: %v, ProductId: %v, Color: %v", t.ID, t.ProductId, t.Color)
}
