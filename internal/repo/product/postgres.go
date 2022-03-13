package product

import (
	"errors"
	"fmt"
	"github.com/seed95/OrderManagement/Microservice/product-service/internal"
	"github.com/seed95/OrderManagement/Microservice/product-service/internal/derror"
	"github.com/seed95/OrderManagement/Microservice/product-service/internal/model"
	"github.com/seed95/OrderManagement/Microservice/product-service/internal/repo"
	"github.com/seed95/OrderManagement/Microservice/product-service/internal/repo/product/schema"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"strconv"

	gorm_schema "gorm.io/gorm/schema"
)

type (
	productRepo struct {
		db     *gorm.DB
		config *internal.PostgresConfig
	}

	Setting struct {
		Config *internal.PostgresConfig
	}
)

var _ repo.ProductRepo = (*productRepo)(nil)

func New(s *Setting) (repo.ProductRepo, error) {
	productRepo := &productRepo{
		config: s.Config,
	}

	if err := productRepo.connect(); err != nil {
		return nil, err
	}

	if err := productRepo.migration(); err != nil {
		return nil, err
	}

	return productRepo, nil
}

func (r *productRepo) connect() error {

	postgresDB, err := gorm.Open(postgres.Open(r.config.PostgresUri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Disable default gorm log
		NamingStrategy: gorm_schema.NamingStrategy{
			TablePrefix:   "tbl_",
			SingularTable: true,
		},
	})

	if err != nil {
		return errors.New(fmt.Sprintf(derror.CreateProductRepoErrorFormat, err))
	}

	r.db = postgresDB

	return nil
}

func (r *productRepo) migration() error {
	if err := r.db.AutoMigrate(&schema.Product{}, &schema.Dimension{}, &schema.Theme{}); err != nil {
		return errors.New(fmt.Sprintf(derror.CreateProductRepoErrorFormat, err))
	}

	return nil
}

// CreateProduct add one row for each size and color of product
func (r *productRepo) CreateProduct(product *model.Product) (*model.Product, error) {

	if product == nil {
		return nil, derror.NilProduct
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {

		schemaProduct := schema.GetProduct(product)

		if err := tx.Create(schemaProduct).Error; err != nil {
			return derror.New(derror.InternalServer, err.Error())
		}
		product.Id = schemaProduct.ID

		if len(product.Dimensions) != 0 {
			schemaDimension := schema.GetDimensions(product)
			if err := tx.Create(schemaDimension).Error; err != nil {
				return derror.New(derror.InternalServer, err.Error())
			}
		}

		if len(product.Colors) != 0 {
			schemaThemes := schema.GetThemes(product)
			if err := tx.Create(schemaThemes).Error; err != nil {
				return derror.New(derror.InternalServer, err.Error())
			}
		}

		return nil
	})

	if err != nil {
		return nil, derror.New(derror.InternalServer, err.Error())
	}

	return product, nil
}

// GetAllCarpet return all carpets for `companyId` in view
func (r *productRepo) GetAllCarpet(companyId uint) ([]model.Carpet, error) {
	var schemaCarpets []schema.Carpet
	viewName := "view_carpet_company_id_" + strconv.FormatUint(uint64(companyId), 10)
	if err := r.db.Table(viewName).Find(&schemaCarpets).Error; err != nil {
		return nil, derror.New(derror.InternalServer, err.Error())
	}

	carpets := make([]model.Carpet, len(schemaCarpets))
	for i, c := range schemaCarpets {
		carpets[i] = model.Carpet{
			Id:         c.Id,
			CompanyId:  companyId,
			DesignCode: c.DesignCode,
			Size:       c.Size,
			Color:      c.Color,
		}
	}

	return carpets, nil
}

// DeleteProduct soft delete product and relations (associations)
func (r *productRepo) DeleteProduct(productId uint) error {
	if err := r.db.Select(clause.Associations).Delete(&schema.Product{Model: gorm.Model{ID: productId}}, productId).Error; err != nil {
		return derror.New(derror.InternalServer, err.Error())
	}

	return nil
}
