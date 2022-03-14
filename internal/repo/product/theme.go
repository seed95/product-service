package product

import (
	"errors"
	"github.com/seed95/product-service/internal/derror"
	"github.com/seed95/product-service/internal/repo/product/schema"
	"gorm.io/gorm"
)

func (r *productRepo) AddColorsToProduct(productId uint, colors []string) ([]schema.Theme, error) {
	if len(colors) == 0 {
		return nil, derror.InvalidColor
	}

	themes := make([]schema.Theme, len(colors))
	for i, c := range colors {
		themes[i] = schema.Theme{
			ProductId: productId,
			Color:     c,
		}
	}

	if err := r.db.Create(&themes).Error; err != nil {
		return nil, derror.New(derror.InternalServer, err.Error())
	}
	return themes, nil
}

func (r *productRepo) DeleteColorsInProduct(productId uint, colors []string) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		for _, c := range colors {
			db := tx.Model(&schema.Theme{}).Where("product_id = ? AND color = ?", productId, c).Delete(&schema.Theme{})
			if db.RowsAffected < 1 {
				return gorm.ErrRecordNotFound
			} else if err := db.Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return derror.ThemeNotFound
		}
		return derror.New(derror.InternalServer, err.Error())
	}

	return nil
}

// UpdateColorsWithId update color if exist with color_id and product_id
// can be update colors for multi product
func (r *productRepo) UpdateColorsWithId(themes []schema.Theme) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		for _, t := range themes {
			db := tx.Model(&t).Where("product_id = ?", t.ProductId).Update("color", t.Color)
			if db.RowsAffected < 1 {
				return gorm.ErrRecordNotFound
			} else if err := db.Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return derror.ThemeNotFound
		}
		return derror.New(derror.InternalServer, err.Error())
	}

	return nil
}

// GetThemesWithProductId return all themes for `productId`
func (r *productRepo) GetThemesWithProductId(productId uint) ([]schema.Theme, error) {
	var themes []schema.Theme
	tx := r.db.Order("id ASC").Model(&schema.Theme{}).Where("product_id = ?", productId).Find(&themes)
	if err := tx.Error; err != nil {
		return nil, derror.New(derror.InternalServer, err.Error())
	}
	return themes, nil
}
