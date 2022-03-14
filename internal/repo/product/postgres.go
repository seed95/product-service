package product

import (
	"errors"
	"fmt"
	"github.com/seed95/product-service/internal"
	"github.com/seed95/product-service/internal/derror"
	"github.com/seed95/product-service/internal/repo"
	"github.com/seed95/product-service/internal/repo/product/schema"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	gormSchema "gorm.io/gorm/schema"
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
		NamingStrategy: gormSchema.NamingStrategy{
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
