package product

import (
	"errors"
	"github.com/seed95/product-service/internal/derror"
	"github.com/seed95/product-service/internal/repo/product/schema"
	"gorm.io/gorm"
)

// AddThemesWithColor if a color for `productId` is duplicate, no add any colors
// support roll back if happen error (a color for `productId` is duplicate)
func (r *productRepo) AddThemesWithColor(productId uint, colors []string) ([]schema.Theme, error) {
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

// DeleteThemesWithColor if a color for `productId` not found, no delete any themes and return `derror.ThemeNotFound`
// support roll back if happen error (a color for `productId` not found)
func (r *productRepo) DeleteThemesWithColor(productId uint, colors []string) error {
	if len(colors) == 0 {
		return derror.InvalidColor
	}

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

// DeleteThemesWithId delete specific colors(colorIds) for `productId`
// delete colors it finds. if a `colorId` not found for `productId` do nothing and delete next `colorId`
// don't support roll back if not found a `colorId`
func (r *productRepo) DeleteThemesWithId(productId uint, colorIds []uint) error {

	if len(colorIds) == 0 {
		return derror.InvalidColor
	}

	themes := make([]schema.Theme, len(colorIds))
	for i, cId := range colorIds {
		themes[i] = schema.Theme{Model: gorm.Model{ID: cId}}
	}

	tx := r.db.Where("product_id = ?", productId).Delete(&themes)
	if tx.RowsAffected < 1 {
		return derror.ThemeNotFound

	} else if err := tx.Error; err != nil {
		return derror.New(derror.InternalServer, err.Error())
	}

	//err := r.db.Transaction(func(tx *gorm.DB) error {
	//	for _, c := range colors {
	//		db := tx.Model(&schema.Theme{}).Where("product_id = ? AND color = ?", productId, c).Delete(&schema.Theme{})
	//		if db.RowsAffected < 1 {
	//			return gorm.ErrRecordNotFound
	//		} else if err := db.Error; err != nil {
	//			return err
	//		}
	//	}
	//	return nil
	//})

	//if err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		return derror.ThemeNotFound
	//	}
	//	return derror.New(derror.InternalServer, err.Error())
	//}

	return nil
}

// UpdateThemesWithId update color in theme if exist with theme_id and product_id
// support roll back if not found a theme_id and product_id
func (r *productRepo) UpdateThemesWithId(productId uint, themes []schema.Theme) error {
	if len(themes) == 0 {
		return derror.InvalidTheme
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		for _, t := range themes {
			db := tx.Model(&t).Where("product_id = ?", productId).Update("color", t.Color)
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

// EditThemesWithId edit all themes for `productId`
func (r *productRepo) EditThemesWithId(productId uint, editedThemes []schema.Theme) error {

	if len(editedThemes) == 0 {
		return derror.InvalidTheme
	}

	originalThemes, err := r.GetThemesWithProductId(productId)
	if err != nil {
		return err
	}

	var updatedThemes []schema.Theme
	var deletedThemes []uint

	// Delete and update themes
OriginalLoop:
	for _, ot := range originalThemes {
		for _, et := range editedThemes {
			if ot.ID == et.ID && ot.Color != et.Color { //Update color
				updatedThemes = append(updatedThemes, et)
				continue OriginalLoop
			}
		}
		deletedThemes = append(deletedThemes, ot.ID)
	}

	// Add themes
	var addThemes []string
EditedLoop:
	for _, et := range editedThemes {
		for _, ot := range originalThemes {
			if ot.ID == et.ID && ot.Color != et.Color {
				continue EditedLoop
			}
		}
		addThemes = append(addThemes, et.Color)
	}

	err := r.UpdateThemesWithId(productId, updatedThemes)
	if

	return nil
}
