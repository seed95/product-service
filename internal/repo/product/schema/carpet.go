package schema

import "github.com/seed95/product-service/internal/model"

type (
	Carpet struct {
		Id          string
		ProductId   uint
		DimensionId uint
		ThemeId     uint
		DesignCode  string
		Size        string
		Color       string
	}
)

func CarpetToModel(c *Carpet, companyId uint) model.Carpet {
	return model.Carpet{
		Id:          c.Id,
		CompanyId:   companyId,
		ProductId:   c.ProductId,
		DimensionId: c.DimensionId,
		ThemeId:     c.ThemeId,
		DesignCode:  c.DesignCode,
		Dimension:   c.Size,
		Color:       c.Color,
	}
}
