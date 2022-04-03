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
	dimensionRepo struct {
		logger logger.Logger
	}

	DimensionService interface {
		GetDimensionsWithProductId(db *gorm.DB, productId uint) ([]schema.Dimension, error)
		InsertDimensions(tx *gorm.DB, productId uint, sizes []string) ([]schema.Dimension, error)
		DeleteDimensionsWithId(tx *gorm.DB, productId uint, dimensions []schema.Dimension) error
		EditDimensions(tx *gorm.DB, productId uint, editedDimensions []schema.Dimension) ([]schema.Dimension, error)
	}
)

var _ DimensionService = (*dimensionRepo)(nil)

func NewDimensionService(l logger.Logger) DimensionService {
	return &dimensionRepo{logger: l}
}

// GetDimensionsWithProductId return all dimensions for `productId`
func (r *dimensionRepo) GetDimensionsWithProductId(db *gorm.DB, productId uint) (dimensions []schema.Dimension, err error) {
	// Log request response
	defer func() {
		commonKeyVal := []keyval.Pair{
			keyval.String("product_id", fmt.Sprintf("%v", productId)),
			keyval.String("dimensions", fmt.Sprintf("%+v", dimensions)),
		}
		logger.LogReqRes(r.logger, "dimension.GetDimensionsWithProductId", err, commonKeyVal...)
	}()

	tx := db.Order("id ASC").Model(&schema.Dimension{}).Where("product_id = ?", productId).Find(&dimensions)
	if err := tx.Error; err != nil {
		return nil, derror.New(derror.InternalServer, err.Error())
	}
	return dimensions, nil
}

// InsertDimensions if a size for `productId` is duplicate, no add any dimensions
// support roll back
func (r *dimensionRepo) InsertDimensions(tx *gorm.DB, productId uint, sizes []string) (dimensions []schema.Dimension, err error) {
	// Log request response
	defer func() {
		commonKeyVal := []keyval.Pair{
			keyval.String("product_id", fmt.Sprintf("%v", productId)),
			keyval.String("sizes", fmt.Sprintf("%v", sizes)),
			keyval.String("dimensions", fmt.Sprintf("%+v", dimensions)),
		}
		logger.LogReqRes(r.logger, "dimension.InsertDimensions", err, commonKeyVal...)
	}()

	if len(sizes) == 0 {
		return nil, derror.InvalidDimension
	}

	dimensions = make([]schema.Dimension, len(sizes))
	for i, s := range sizes {
		dimensions[i] = schema.Dimension{
			ProductId: productId,
			Size:      s,
		}
	}

	if err := tx.Create(&dimensions).Error; err != nil {
		return nil, derror.New(derror.InternalServer, err.Error())
	}
	return dimensions, nil
}

// DeleteDimensionsWithId delete specific dimension(dimensionIds) for `productId`
// delete dimension it finds. if a `dimensionId` not found for `productId` do nothing and delete next `dimensionId`
// if a `dimensionId` not found return derror.DimensionNotFound
// don't support roll back if not found a `dimensionId`
func (r *dimensionRepo) DeleteDimensionsWithId(tx *gorm.DB, productId uint, dimensions []schema.Dimension) (err error) {
	// Log request response
	defer func() {
		commonKeyVal := []keyval.Pair{
			keyval.String("product_id", fmt.Sprintf("%v", productId)),
			keyval.String("dimensions", fmt.Sprintf("%+v", dimensions)),
		}
		logger.LogReqRes(r.logger, "dimension.DeleteDimensionsWithId", err, commonKeyVal...)
	}()

	if len(dimensions) == 0 {
		return derror.InvalidDimension
	}

	db := tx.Where("product_id = ?", productId).Delete(&dimensions)
	if db.RowsAffected != int64(len(dimensions)) {
		return derror.DimensionNotFound
	} else if err := db.Error; err != nil {
		return derror.New(derror.InternalServer, err.Error())
	}

	return nil
}

func (r *dimensionRepo) EditDimensions(tx *gorm.DB, productId uint, editedDimensions []schema.Dimension) (dimensions []schema.Dimension, err error) {
	// Log request response
	defer func() {
		commonKeyVal := []keyval.Pair{
			keyval.String("product_id", fmt.Sprintf("%v", productId)),
			keyval.String("edited_dimensions", fmt.Sprintf("%+v", editedDimensions)),
			keyval.String("dimensions", fmt.Sprintf("%+v", dimensions)),
		}
		logger.LogReqRes(r.logger, "dimension.EditDimensions", err, commonKeyVal...)
	}()

	if len(editedDimensions) == 0 {
		return nil, derror.InvalidDimension
	}

	originalDimensions, err := r.GetDimensionsWithProductId(tx, productId)
	if err != nil {
		return nil, err
	}

	var deletedDimensions []schema.Dimension

	// Delete dimensions
OriginalLoop:
	for _, od := range originalDimensions {
		for _, ed := range editedDimensions {
			if od.Size == ed.Size {
				dimensions = append(dimensions, od)
				continue OriginalLoop
			}
		}
		deletedDimensions = append(deletedDimensions, od)
	}

	// New dimensions
	var newSizes []string
EditedLoop:
	for _, ed := range editedDimensions {
		for _, od := range originalDimensions {
			if od.Size == ed.Size {
				continue EditedLoop
			}
		}
		newSizes = append(newSizes, ed.Size)
	}

	if len(deletedDimensions) != 0 {
		err = r.DeleteDimensionsWithId(tx, productId, deletedDimensions)
		if err != nil {
			return nil, err
		}
	}

	if len(newSizes) != 0 {
		newDimensions, err := r.InsertDimensions(tx, productId, newSizes)
		if err != nil {
			return nil, err
		}
		dimensions = append(dimensions, newDimensions...)
	}

	return dimensions, nil

}
