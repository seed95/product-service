package product

import (
	"fmt"
	"github.com/seed95/product-service/internal/derror"
	"github.com/seed95/product-service/internal/repo/product/schema"
	"github.com/seed95/product-service/pkg/logger"
	"github.com/seed95/product-service/pkg/logger/keyval"
	"gorm.io/gorm"
)

type (
	themeRepo struct {
		logger logger.Logger
	}

	ThemeService interface {
		GetThemesWithProductId(db *gorm.DB, productId uint) ([]schema.Theme, error)
		InsertThemesWithColor(tx *gorm.DB, productId uint, colors []string) ([]schema.Theme, error)
		DeleteThemesWithId(tx *gorm.DB, productId uint, themes []schema.Theme) error
		EditThemes(tx *gorm.DB, productId uint, editedThemes []schema.Theme) ([]schema.Theme, error)
	}
)

var _ ThemeService = (*themeRepo)(nil)

func NewThemeService(l logger.Logger) ThemeService {
	return &themeRepo{logger: l}
}

// GetThemesWithProductId return all themes for `productId`
func (r *themeRepo) GetThemesWithProductId(db *gorm.DB, productId uint) (themes []schema.Theme, err error) {
	// Log request response
	defer func() {
		commonKeyVal := []keyval.Pair{
			keyval.String("product_id", fmt.Sprintf("%v", productId)),
			keyval.String("themes", fmt.Sprintf("%+v", themes)),
		}
		logger.LogReqRes(r.logger, "theme.GetThemesWithProductId", err, commonKeyVal...)
	}()

	tx := db.Order("id ASC").Model(&schema.Theme{}).Where("product_id = ?", productId).Find(&themes)
	if err := tx.Error; err != nil {
		return nil, derror.New(derror.InternalServer, err.Error())
	}
	return themes, nil
}

// InsertThemesWithColor if a color for `productId` is duplicate, no add any themes
// support roll back
func (r *themeRepo) InsertThemesWithColor(tx *gorm.DB, productId uint, colors []string) (themes []schema.Theme, err error) {
	// Log request response
	defer func() {
		commonKeyVal := []keyval.Pair{
			keyval.String("product_id", fmt.Sprintf("%v", productId)),
			keyval.String("colors", fmt.Sprintf("%v", colors)),
			keyval.String("themes", fmt.Sprintf("%+v", themes)),
		}
		logger.LogReqRes(r.logger, "theme.InsertThemesWithColor", err, commonKeyVal...)
	}()

	if len(colors) == 0 {
		return nil, derror.InvalidColor
	}

	themes = make([]schema.Theme, len(colors))
	for i, c := range colors {
		themes[i] = schema.Theme{
			ProductId: productId,
			Color:     c,
		}
	}

	if err := tx.Create(&themes).Error; err != nil {
		return nil, derror.New(derror.InternalServer, err.Error())
	}
	return themes, nil
}

// DeleteThemesWithId delete specific colors(colorIds) for `productId`
// delete theme it finds. if a `themeId` not found for `productId` do nothing and delete next `themeId`
// if a `themeId` not found return derror.ThemeNotFound
// don't support roll back if not found a `themeId`
func (r *themeRepo) DeleteThemesWithId(tx *gorm.DB, productId uint, themes []schema.Theme) (err error) {
	// Log request response
	defer func() {
		commonKeyVal := []keyval.Pair{
			keyval.String("product_id", fmt.Sprintf("%v", productId)),
			keyval.String("themes", fmt.Sprintf("%+v", themes)),
		}
		logger.LogReqRes(r.logger, "theme.DeleteThemesWithId", err, commonKeyVal...)
	}()

	if len(themes) == 0 {
		return derror.InvalidTheme
	}

	db := tx.Where("product_id = ?", productId).Delete(&themes)
	if db.RowsAffected != int64(len(themes)) {
		return derror.ThemeNotFound
	} else if err := db.Error; err != nil {
		return derror.New(derror.InternalServer, err.Error())
	}

	return nil
}

func (r *themeRepo) EditThemes(tx *gorm.DB, productId uint, editedThemes []schema.Theme) (themes []schema.Theme, err error) {
	// Log request response
	defer func() {
		commonKeyVal := []keyval.Pair{
			keyval.String("product_id", fmt.Sprintf("%v", productId)),
			keyval.String("edited_themes", fmt.Sprintf("%+v", editedThemes)),
			keyval.String("themes", fmt.Sprintf("%+v", themes)),
		}
		logger.LogReqRes(r.logger, "theme.EditThemes", err, commonKeyVal...)
	}()

	if len(editedThemes) == 0 {
		return nil, derror.InvalidTheme
	}

	originalThemes, err := r.GetThemesWithProductId(tx, productId)
	if err != nil {
		return nil, err
	}

	var deletedThemes []schema.Theme

	// Delete themes
OriginalLoop:
	for _, ot := range originalThemes {
		for _, et := range editedThemes {
			if ot.Color == et.Color {
				themes = append(themes, ot)
				continue OriginalLoop
			}
		}
		deletedThemes = append(deletedThemes, ot)
	}

	// New themes
	var newColors []string
EditedLoop:
	for _, et := range editedThemes {
		for _, ot := range originalThemes {
			if ot.Color == et.Color {
				continue EditedLoop
			}
		}
		newColors = append(newColors, et.Color)
	}

	if len(deletedThemes) != 0 {
		err = r.DeleteThemesWithId(tx, productId, deletedThemes)
		if err != nil {
			return nil, err
		}
	}

	if len(newColors) != 0 {
		newThemes, err := r.InsertThemesWithColor(tx, productId, newColors)
		if err != nil {
			return nil, err
		}
		themes = append(themes, newThemes...)
	}

	return themes, nil
}
