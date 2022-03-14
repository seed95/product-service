package schema

import (
	"fmt"
	"github.com/seed95/product-service/internal/model"
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
)

func ProductModelToSchema(p model.Product) *Product {
	result := Product{
		Model:       gorm.Model{ID: p.Id},
		CompanyId:   p.CompanyId,
		DesignCode:  p.DesignCode,
		Description: p.Description,
	}

	for _, d := range p.Dimensions {
		result.Dimensions = append(result.Dimensions, Dimension{
			ProductId: p.Id,
			Size:      d,
		})
	}

	for _, c := range p.Colors {
		result.Themes = append(result.Themes, Theme{
			ProductId: p.Id,
			Color:     c,
		})
	}

	return &result
}

func (p Product) String() string {

	var dimension []string
	for _, d := range p.Dimensions {
		dimension = append(dimension, d.Size)
	}

	var theme []string
	for _, t := range p.Themes {
		theme = append(theme, t.Color)
	}

	return fmt.Sprintf("ID: %v,\t CompanyId: %v,\t DesignCode: %v,\t Description: %v,\t Dimensions: %v,\t Theme: %v,\t",
		p.ID, p.CompanyId, p.DesignCode, p.Description, dimension, theme)
}
