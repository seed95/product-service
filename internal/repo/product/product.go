package product

import (
	"errors"
	"github.com/seed95/product-service/internal/derror"
	"github.com/seed95/product-service/internal/model"
	"github.com/seed95/product-service/internal/repo/product/schema"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CreateProduct add one row for each size and color of product
func (r *productRepo) CreateProduct(product model.Product) (*schema.Product, error) {
	schemaProduct := schema.ProductModelToSchema(product)

	if err := r.db.Create(schemaProduct).Error; err != nil {
		return nil, derror.New(derror.InternalServer, err.Error())
	}

	return schemaProduct, nil
}

func (r *productRepo) GetProductWithId(productId uint) (*schema.Product, error) {
	product := schema.Product{
		Model: gorm.Model{ID: productId},
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.First(&product).Error; err != nil {
			return err
		}

		if err := tx.Order("id ASC").Where("product_id = ?", productId).Find(&product.Dimensions).Error; err != nil {
			return err
		}

		if err := tx.Order("id ASC").Where("product_id = ?", productId).Find(&product.Themes).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, derror.ProductNotFound
		}
		return nil, derror.New(derror.InternalServer, err.Error())
	}

	return &product, nil
}

// DeleteProduct soft delete product and relations (associations)
func (r *productRepo) DeleteProduct(productId uint) error {
	tx := r.db.Select(clause.Associations).Delete(&schema.Product{Model: gorm.Model{ID: productId}})
	if tx.RowsAffected < 1 {
		return derror.ProductNotFound
	} else if err := tx.Error; err != nil {
		return derror.New(derror.InternalServer, err.Error())
	}

	return nil
}

func (r *productRepo) EditProduct(editedProduct schema.Product) (*schema.Product, error) {

	originalProduct, err := r.GetProductWithId(editedProduct.ID)
	_ = originalProduct
	if err != nil {
		return nil, derror.New(derror.InternalServer, err.Error())
	}

	err = r.db.Transaction(func(tx *gorm.DB) error {

		productFields := map[string]interface{}{"design_code": editedProduct.DesignCode, "description": editedProduct.Description}
		if err := tx.Model(&schema.Product{}).Where("product_id = ?", editedProduct.ID).Updates(productFields).Error; err != nil {
			return err
		}

		//if len(product.Dimensions) != 0 {
		//	schemaDimension := schema.GetDimensions(&product)
		//	if err := tx.Create(schemaDimension).Error; err != nil {
		//		return derror.New(derror.InternalServer, err.Error())
		//	}
		//}
		//
		//if len(product.Colors) != 0 {
		//	schemaThemes := schema.GetThemesFromProductModel(&product)
		//	if err := tx.Create(schemaThemes).Error; err != nil {
		//		return derror.New(derror.InternalServer, err.Error())
		//	}
		//}

		return nil
	})

	if err != nil {
		return nil, derror.New(derror.InternalServer, err.Error())
	}

	return nil, nil
}
