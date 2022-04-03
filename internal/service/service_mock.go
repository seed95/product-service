package service

import (
	"context"
	"github.com/seed95/product-service/internal/api"
	"github.com/seed95/product-service/internal/repo/product"
	"github.com/seed95/product-service/pkg/logger/zap"
	"github.com/stretchr/testify/require"
	"testing"
)

func NewServiceMock(t *testing.T) ProductService {
	// NewProduct repo
	pRepo, err := product.NewProductRepoMock()
	require.Nil(t, err)
	require.NotNil(t, pRepo)

	service, err := New(&Setting{ProductRepo: pRepo, Logger: zap.NopLogger})
	require.Nil(t, err)
	require.NotNil(t, service)
	return service
}

func GetProduct1() *api.Product {
	return &api.Product{
		CompanyId:   1,
		CompanyName: "Negin",
		DesignCode:  "105",
		Description: "توضیحات ۱۰۵",
		Sizes:       []string{"6", "9"},
		Colors:      []string{"قرمز", "آبی"},
	}
}

func GetProduct2() *api.Product {
	return &api.Product{
		CompanyId:   1,
		CompanyName: "Negin",
		DesignCode:  "106",
		Description: "توضیحات ۱۰۶",
		Sizes:       []string{"6", "9"},
		Colors:      []string{"قرمز", "آبی"},
	}
}

func GetProduct3() *api.Product {
	return &api.Product{
		CompanyId:   1,
		CompanyName: "Negin",
		DesignCode:  "107",
		Description: "توضیحات ۱۰۷",
		Sizes:       []string{"6", "9"},
		Colors:      []string{"قرمز", "آبی"},
	}
}

func CreateProduct1(service ProductService, t *testing.T) {
	ctx := context.Background()
	req := api.CreateNewProductRequest{NewProduct: GetProduct1()}

	_, err := service.CreateNewProduct(ctx, &req)
	require.Nil(t, err)
}

func CreateProduct2(service ProductService, t *testing.T) {
	ctx := context.Background()
	req := api.CreateNewProductRequest{NewProduct: GetProduct2()}

	_, err := service.CreateNewProduct(ctx, &req)
	require.Nil(t, err)
}

func CreateProduct3(service ProductService, t *testing.T) {
	ctx := context.Background()
	req := api.CreateNewProductRequest{NewProduct: GetProduct3()}

	_, err := service.CreateNewProduct(ctx, &req)
	require.Nil(t, err)
}
