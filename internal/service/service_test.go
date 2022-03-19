package service

import (
	"context"
	"github.com/seed95/product-service/internal/api"
	"github.com/seed95/product-service/internal/derror"
	"github.com/seed95/product-service/internal/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGateway_CreateNewProduct_Ok(t *testing.T) {
	// Service mock
	service := NewServiceMock(t)

	ctx := context.Background()
	req := api.CreateNewProductRequest{NewProduct: GetProduct1()}

	res, err := service.CreateNewProduct(ctx, &req)
	require.Nil(t, err)
	require.Equal(t, 1, len(res.Products))
}

func TestGateway_CreateNewProduct_Empty(t *testing.T) {
	// Service mock
	service := NewServiceMock(t)

	ctx := context.Background()
	t.Run("design code", func(t *testing.T) {
		p1 := GetProduct1()
		p1.DesignCode = ""
		req := api.CreateNewProductRequest{NewProduct: p1}

		res, err := service.CreateNewProduct(ctx, &req)
		require.Equal(t, derror.InvalidProduct, err)
		require.Nil(t, res)
	})

	t.Run("color", func(t *testing.T) {
		p1 := GetProduct1()
		p1.Colors = []string{}
		req := api.CreateNewProductRequest{NewProduct: p1}

		res, err := service.CreateNewProduct(ctx, &req)
		require.Equal(t, derror.InvalidProduct, err)
		require.Nil(t, res)
	})

	t.Run("size", func(t *testing.T) {
		p1 := GetProduct1()
		p1.Sizes = []string{}
		req := api.CreateNewProductRequest{NewProduct: p1}

		res, err := service.CreateNewProduct(ctx, &req)
		require.Equal(t, derror.InvalidProduct, err)
		require.Nil(t, res)
	})

	t.Run("description", func(t *testing.T) {
		p1 := GetProduct1()
		p1.Description = ""
		req := api.CreateNewProductRequest{NewProduct: p1}

		res, err := service.CreateNewProduct(ctx, &req)
		require.Nil(t, err)
		require.Equal(t, 1, len(res.Products))
	})
}

func TestGateway_CreateNewProduct_NonZeroId(t *testing.T) {
	// Service mock
	service := NewServiceMock(t)

	ctx := context.Background()
	req := api.CreateNewProductRequest{NewProduct: GetProduct1()}
	req.NewProduct.Id = 10

	res, err := service.CreateNewProduct(ctx, &req)
	require.Equal(t, derror.InvalidProduct, err)
	require.Nil(t, res)
}

func TestGateway_CreateNewProduct_CompanyZeroId(t *testing.T) {
	// Service mock
	service := NewServiceMock(t)

	ctx := context.Background()
	req := api.CreateNewProductRequest{NewProduct: GetProduct1()}
	req.NewProduct.CompanyId = 0

	res, err := service.CreateNewProduct(ctx, &req)
	require.Equal(t, derror.InvalidProduct, err)
	require.Nil(t, res)
}

func TestProductIsValid(t *testing.T) {

	tests := []struct {
		Name    string
		Product model.Product
		Valid   bool
	}{
		{
			Name: "Ok",
			Product: model.Product{
				Id:          0,
				CompanyName: "Negin",
				CompanyId:   1,
				DesignCode:  "105",
				Colors:      []string{"آبی"},
				Sizes:       []string{"12"},
				Description: "",
			},
			Valid: true,
		},
		{
			Name: "two color",
			Product: model.Product{
				Id:          0,
				CompanyName: "Negin",
				CompanyId:   1,
				DesignCode:  "105",
				Colors:      []string{"قرمز", "آبی"},
				Sizes:       []string{"12"},
				Description: "",
			},
			Valid: true,
		},
		{
			Name: "two size",
			Product: model.Product{
				Id:          0,
				CompanyName: "Negin",
				CompanyId:   1,
				DesignCode:  "105",
				Colors:      []string{"آبی"},
				Sizes:       []string{"12", "6"},
				Description: "",
			},
			Valid: true,
		},
		{
			Name: "non zero id",
			Product: model.Product{
				Id:          100,
				CompanyName: "Negin",
				CompanyId:   1,
				DesignCode:  "105",
				Colors:      []string{"آبی"},
				Sizes:       []string{"12"},
				Description: "",
			},
			Valid: true,
		},
		{
			Name: "empty company name",
			Product: model.Product{
				Id:          0,
				CompanyName: "",
				CompanyId:   1,
				DesignCode:  "105",
				Colors:      []string{"آبی"},
				Sizes:       []string{"12"},
				Description: "",
			},
			Valid: true,
		},
		{
			Name: "zero company id",
			Product: model.Product{
				Id:          0,
				CompanyName: "Negin",
				CompanyId:   0,
				DesignCode:  "105",
				Colors:      []string{"آبی"},
				Sizes:       []string{"12"},
				Description: "",
			},
			Valid: false,
		},
		{
			Name: "empty design code",
			Product: model.Product{
				Id:          0,
				CompanyName: "Negin",
				CompanyId:   1,
				DesignCode:  "",
				Colors:      []string{"آبی"},
				Sizes:       []string{"12"},
				Description: "",
			},
			Valid: false,
		},
		{
			Name: "empty color",
			Product: model.Product{
				Id:          0,
				CompanyName: "Negin",
				CompanyId:   1,
				DesignCode:  "105",
				Colors:      []string{},
				Sizes:       []string{"12"},
				Description: "",
			},
			Valid: false,
		},
		{
			Name: "empty size",
			Product: model.Product{
				Id:          0,
				CompanyName: "Negin",
				CompanyId:   1,
				DesignCode:  "105",
				Colors:      []string{"آبی"},
				Sizes:       []string{},
				Description: "",
			},
			Valid: false,
		},
		{
			Name: "color with empty value",
			Product: model.Product{
				Id:          0,
				CompanyName: "Negin",
				CompanyId:   1,
				DesignCode:  "105",
				Colors:      []string{""},
				Sizes:       []string{"12"},
				Description: "",
			},
			Valid: false,
		},
		{
			Name: "size with empty value",
			Product: model.Product{
				Id:          0,
				CompanyName: "Negin",
				CompanyId:   1,
				DesignCode:  "105",
				Colors:      []string{"آبی"},
				Sizes:       []string{""},
				Description: "",
			},
			Valid: false,
		},
		{
			Name: "duplicate size",
			Product: model.Product{
				Id:          0,
				CompanyName: "Negin",
				CompanyId:   1,
				DesignCode:  "105",
				Colors:      []string{"آبی"},
				Sizes:       []string{"12", "12"},
				Description: "",
			},
			Valid: false,
		},
		{
			Name: "duplicate color",
			Product: model.Product{
				Id:          0,
				CompanyName: "Negin",
				CompanyId:   1,
				DesignCode:  "105",
				Colors:      []string{"آبی", "آبی"},
				Sizes:       []string{"12"},
				Description: "",
			},
			Valid: false,
		},
		{
			Name: "color with value and empty",
			Product: model.Product{
				Id:          0,
				CompanyName: "Negin",
				CompanyId:   1,
				DesignCode:  "105",
				Colors:      []string{"آبی", ""},
				Sizes:       []string{"12"},
				Description: "",
			},
			Valid: false,
		},
		{
			Name: "size with value and empty",
			Product: model.Product{
				Id:          0,
				CompanyName: "Negin",
				CompanyId:   1,
				DesignCode:  "105",
				Colors:      []string{"آبی"},
				Sizes:       []string{"12", ""},
				Description: "",
			},
			Valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			valid := productIsValid(tt.Product)
			require.Equal(t, tt.Valid, valid, tt.Product)
		})
	}

}
