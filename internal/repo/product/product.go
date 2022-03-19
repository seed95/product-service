package product

import (
	"errors"
	"github.com/seed95/product-service/internal/derror"
	"github.com/seed95/product-service/internal/model"
	"github.com/seed95/product-service/internal/repo/product/schema"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CreateProduct create a product with relations(dimension, theme)
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

	if err := r.db.Preload(clause.Associations).First(&product).Error; err != nil {
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

func (r *productRepo) EditProduct(product model.Product) (*schema.Product, error) {
	schemaProduct := schema.ProductModelToSchema(product)

	// Check product exist
	_, err := r.GetProductWithId(schemaProduct.ID)
	if err != nil {
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

func (r *productRepo) GetAllProducts(companyId uint) ([]schema.Product, error) {
	var products []schema.Product

	tx := r.db.Model(&schema.Product{}).Preload(clause.Associations).Where("company_id = ?", companyId).Find(&products)
	if tx.RowsAffected < 1 {
		return nil, derror.ProductNotFound
	} else if err := tx.Error; err != nil {
		return nil, derror.New(derror.InternalServer, err.Error())
	}

	return products, nil
}
