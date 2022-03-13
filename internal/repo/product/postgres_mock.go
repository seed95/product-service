package product

import (
	"github.com/seed95/OrderManagement/Microservice/product-service/internal"
)

func NewProductRepoMock() (*productRepo, error) {

	mock := productRepo{
		config: &internal.PostgresConfig{
			PostgresUri: "host=localhost user=seed password=seed@1400 dbname=db_test port=5432 sslmode=disable",
		},
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

	return &mock, nil
}
