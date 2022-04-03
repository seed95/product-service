package product

import (
	"github.com/seed95/product-service/internal"
	"github.com/seed95/product-service/internal/model"
	"github.com/seed95/product-service/internal/repo/product/schema"
	"github.com/seed95/product-service/pkg/logger/zap"
	"github.com/stretchr/testify/require"
	"testing"
)

func NewProductRepoMock() (*productRepo, error) {

	mock := productRepo{
		config: &internal.PostgresConfig{
			DSN: "host=localhost user=seed password=seed@1400 dbname=db_test port=5432 sslmode=disable",
		},
		logger: zap.NopLogger,
	}

	if err := mock.connect(); err != nil {
		return nil, err
	}

	if err := mock.migration(); err != nil {
		return nil, err
	}

	if err := mock.db.Exec("TRUNCATE tbl_theme,tbl_dimension,tbl_product;").Error; err != nil {
		return nil, err
	}

	mock.theme = NewThemeService(mock.logger)
	mock.dimension = NewDimensionService(mock.logger)

	return &mock, nil
}

func CreateProduct1(repo *productRepo, t *testing.T) *schema.Product {
	p1 := model.Product{
		CompanyName: "Negin",
		CompanyId:   1,
		DesignCode:  "105",
		Colors:      []string{"قرمز", "آبی"},
		Sizes:       []string{"6", "9"},
		Description: "توضیحات برای کد ۱۰۵",
	}
	p, err := repo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, p)
	return p
}

func CreateProduct2(repo *productRepo, t *testing.T) *schema.Product {
	p2 := model.Product{
		CompanyName: "Negin",
		CompanyId:   1,
		DesignCode:  "106",
		Colors:      []string{"قرمز", "آبی"},
		Sizes:       []string{"6", "9"},
		Description: "توضیحات برای کد ۱۰۶",
	}
	p, err := repo.CreateProduct(p2)
	require.Nil(t, err)
	require.NotNil(t, p)
	return p
}

func CreateProduct3(repo *productRepo, t *testing.T) *schema.Product {
	p3 := model.Product{
		CompanyName: "Negin",
		CompanyId:   1,
		DesignCode:  "107",
		Colors:      []string{"قرمز", "آبی"},
		Sizes:       []string{"6", "9"},
		Description: "توضیحات برای کد ۱۰۷",
	}
	p, err := repo.CreateProduct(p3)
	require.Nil(t, err)
	require.NotNil(t, p)
	return p
}
