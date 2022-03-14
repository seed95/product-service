package product

import (
	"errors"
	"github.com/seed95/product-service/internal/derror"
	"github.com/seed95/product-service/internal/repo/product/schema"
	"gorm.io/gorm"
)

func (r *productRepo) AddDimensionsToProduct(productId uint, sizes []string) ([]schema.Dimension, error) {
	if len(sizes) == 0 {
		return nil, derror.InvalidDimension
	}

	dimensions := make([]schema.Dimension, len(sizes))
	for i, s := range sizes {
		dimensions[i] = schema.Dimension{
			ProductId: productId,
			Size:      s,
		}
	}

	if err := r.db.Create(&dimensions).Error; err != nil {
		return nil, derror.New(derror.InternalServer, err.Error())
	}
	return dimensions, nil
}

func (r *productRepo) DeleteDimensionsInProduct(productId uint, sizes []string) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		for _, s := range sizes {
			db := tx.Model(&schema.Dimension{}).Where("product_id = ? AND size = ?", productId, s).Delete(&schema.Dimension{})
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
			return derror.DimensionNotFound
		}
		return derror.New(derror.InternalServer, err.Error())
	}

	return nil
}

// UpdateDimensionsWithId update dimension if exist with dimension_id and product_id
// can be update dimension for multi product
func (r *productRepo) UpdateDimensionsWithId(dimensions []schema.Dimension) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		for _, d := range dimensions {
			db := tx.Model(&d).Where("product_id = ?", d.ProductId).Update("size", d.Size)
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
			return derror.DimensionNotFound
		}
		return derror.New(derror.InternalServer, err.Error())
	}

	return nil
}
