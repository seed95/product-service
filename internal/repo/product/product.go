package product

import (
	"errors"
	"fmt"
	"github.com/seed95/product-service/internal/derror"
	"github.com/seed95/product-service/internal/model"
	"github.com/seed95/product-service/internal/repo/product/schema"
	"github.com/seed95/product-service/pkg/logger"
	"github.com/seed95/product-service/pkg/logger/keyval"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CreateProduct create a product with relations(dimension, theme)
func (r *productRepo) CreateProduct(product model.Product) (schemaProduct *schema.Product, err error) {
	// Log request response
	defer func() {
		commonKeyVal := []keyval.Pair{
			keyval.String("model_product", fmt.Sprintf("%+v", product)),
			keyval.String("schema_product", fmt.Sprintf("%+v", schemaProduct)),
		}
		logger.LogReqRes(r.logger, "product.CreateProduct", err, commonKeyVal...)
	}()

	schemaProduct = schema.ProductModelToSchema(product)
	if err := r.db.Create(schemaProduct).Error; err != nil {
		return nil, derror.New(derror.InternalServer, err.Error())
	}

	return schemaProduct, nil
}

func (r *productRepo) GetProductWithId(productId uint) (product *schema.Product, err error) {
	// Log request response
	defer func() {
		commonKeyVal := []keyval.Pair{
			keyval.String("id", fmt.Sprintf("%v", productId)),
			keyval.String("product", fmt.Sprintf("%+v", product)),
		}
		logger.LogReqRes(r.logger, "product.GetProductWithId", err, commonKeyVal...)
	}()

	product = &schema.Product{
		Model: gorm.Model{ID: productId},
	}
	if err := r.db.Preload(clause.Associations).First(product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, derror.ProductNotFound
		}
		return nil, derror.New(derror.InternalServer, err.Error())
	}
	return product, nil
}

// DeleteProduct soft delete product and relations (associations)
func (r *productRepo) DeleteProduct(productId uint) (err error) {
	// Log request response
	defer func() {
		commonKeyVal := []keyval.Pair{
			keyval.String("id", fmt.Sprintf("%v", productId)),
		}
		logger.LogReqRes(r.logger, "product.DeleteProduct", err, commonKeyVal...)
	}()

	tx := r.db.Select(clause.Associations).Delete(&schema.Product{Model: gorm.Model{ID: productId}})
	if tx.RowsAffected < 1 {
		return derror.ProductNotFound
	} else if err := tx.Error; err != nil {
		return derror.New(derror.InternalServer, err.Error())
	}

	return nil
}

func (r *productRepo) EditProduct(product model.Product) (schemaProduct *schema.Product, err error) {
	// Log request response
	defer func() {
		commonKeyVal := []keyval.Pair{
			keyval.String("model_product", fmt.Sprintf("%+v", product)),
			keyval.String("schema_product", fmt.Sprintf("%+v", schemaProduct)),
		}
		logger.LogReqRes(r.logger, "product.EditProduct", err, commonKeyVal...)
	}()

	schemaProduct = schema.ProductModelToSchema(product)

	// Check product exist
	if _, err = r.GetProductWithId(schemaProduct.ID); err != nil {
		return nil, err
	}

	err = r.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(schema.Product{Model: gorm.Model{ID: schemaProduct.ID}}).
			Updates(schema.Product{DesignCode: schemaProduct.DesignCode, Description: schemaProduct.Description})
		if err := result.Error; err != nil {
			return err
		}

		themes, err := r.theme.EditThemes(tx, schemaProduct.ID, schemaProduct.Themes)
		if err != nil {
			return err
		}
		schemaProduct.Themes = themes

		dimensions, err := r.dimension.EditDimensions(tx, schemaProduct.ID, schemaProduct.Dimensions)
		if err != nil {
			return err
		}
		schemaProduct.Dimensions = dimensions

		return nil
	})

	if err != nil {
		return nil, derror.New(derror.InternalServer, err.Error())
	}

	return schemaProduct, nil
}

func (r *productRepo) GetAllProducts(companyId uint) (products []schema.Product, err error) {
	// Log request response
	defer func() {
		commonKeyVal := []keyval.Pair{
			keyval.String("company_id", fmt.Sprintf("%v", companyId)),
			keyval.String("products", fmt.Sprintf("%+v", products)),
		}
		logger.LogReqRes(r.logger, "product.GetAllProducts", err, commonKeyVal...)
	}()

	tx := r.db.Model(&schema.Product{}).Preload(clause.Associations).Where("company_id = ?", companyId).Find(&products)
	if tx.RowsAffected < 1 {
		return nil, derror.ProductNotFound
	} else if err := tx.Error; err != nil {
		return nil, derror.New(derror.InternalServer, err.Error())
	}

	return products, nil
}
