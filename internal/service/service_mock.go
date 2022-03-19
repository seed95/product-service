package service

import (
	"github.com/seed95/product-service/internal/api"
	"github.com/seed95/product-service/internal/repo/product"
	"github.com/stretchr/testify/require"
	"testing"
)

func NewServiceMock(t *testing.T) ProductService {
	// NewProduct repo
	pRepo, err := product.NewProductRepoMock()
	require.Nil(t, err)
	require.NotNil(t, pRepo)

	service, err := New(&Setting{ProductRepo: pRepo})
	require.Nil(t, err)
	require.NotNil(t, service)
	return service
}

func GetProduct1() api.Product {
	return api.Product{
		CompanyId:   1,
		CompanyName: "Negin",
		DesignCode:  "105",
		Description: "توضیحات ۱۰۵",
		Sizes:       []string{"6", "9"},
		Colors:      []string{"قرمز", "آبی"},
	}
}
