package service

import (
	"context"
	"github.com/seed95/product-service/internal/api"
	"github.com/seed95/product-service/internal/derror"
	"github.com/seed95/product-service/internal/model"
	"github.com/seed95/product-service/internal/repo/product"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGateway_CreateNewProduct_Ok(t *testing.T) {
	// Service mock
	service := NewServiceMock(t)

	ctx := context.Background()
	req := api.CreateNewProductRequest{Product: GetProduct1()}

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
		req := api.CreateNewProductRequest{Product: p1}

		res, err := service.CreateNewProduct(ctx, &req)
		require.Equal(t, derror.InvalidProduct, err)
		require.Nil(t, res)
	})

	t.Run("color", func(t *testing.T) {
		p1 := GetProduct1()
		p1.Colors = []string{}
		req := api.CreateNewProductRequest{Product: p1}

		res, err := service.CreateNewProduct(ctx, &req)
		require.Equal(t, derror.InvalidProduct, err)
		require.Nil(t, res)
	})

	t.Run("size", func(t *testing.T) {
		p1 := GetProduct1()
		p1.Sizes = []string{}
		req := api.CreateNewProductRequest{Product: p1}

		res, err := service.CreateNewProduct(ctx, &req)
		require.Equal(t, derror.InvalidProduct, err)
		require.Nil(t, res)
	})

	t.Run("description", func(t *testing.T) {
		p1 := GetProduct1()
		p1.Description = ""
		req := api.CreateNewProductRequest{Product: p1}

		res, err := service.CreateNewProduct(ctx, &req)
		require.Nil(t, err)
		require.Equal(t, 1, len(res.Products))
	})
}

func TestGateway_CreateNewProduct_NonZeroId(t *testing.T) {
	// Service mock
	service := NewServiceMock(t)

	ctx := context.Background()
	req := api.CreateNewProductRequest{Product: GetProduct1()}
	req.Product.Id = 10

	res, err := service.CreateNewProduct(ctx, &req)
	require.Equal(t, derror.InvalidProduct, err)
	require.Nil(t, res)
}

func TestGateway_CreateNewProduct_CompanyZeroId(t *testing.T) {
	// Service mock
	service := NewServiceMock(t)

	ctx := context.Background()
	req := api.CreateNewProductRequest{Product: GetProduct1()}
	req.Product.CompanyId = 0

	res, err := service.CreateNewProduct(ctx, &req)
	require.Equal(t, derror.InvalidProduct, err)
	require.Nil(t, res)
}

func TestGateway_GetAllProducts_Ok(t *testing.T) {
	// Service mock
	service := NewServiceMock(t)

	CreateProduct1(service, t)
	CreateProduct2(service, t)
	CreateProduct3(service, t)

	ctx := context.Background()
	res, err := service.GetAllProducts(ctx, 1)
	require.Nil(t, err)
	require.Equal(t, 3, len(res.Products))
}

func TestGateway_GetAllProducts_NoProduct(t *testing.T) {
	// Service mock
	service := NewServiceMock(t)

	CreateProduct1(service, t)
	CreateProduct2(service, t)
	CreateProduct3(service, t)

	ctx := context.Background()
	res, err := service.GetAllProducts(ctx, 2)
	require.Equal(t, derror.ProductNotFound, err)
	require.Nil(t, res)
}

func TestGateway_GetAllProducts_ZeroCompanyId(t *testing.T) {
	// Service mock
	service := NewServiceMock(t)

	CreateProduct1(service, t)
	CreateProduct2(service, t)
	CreateProduct3(service, t)

	ctx := context.Background()
	res, err := service.GetAllProducts(ctx, 0)
	require.Equal(t, derror.InvalidCompany, err)
	require.Nil(t, res)
}

func TestGateway_GetProductWithId_Ok(t *testing.T) {
	// Service mock
	service := NewServiceMock(t)

	// Product repo
	pRepo, err := product.NewProductRepoMock()
	require.Nil(t, err)

	// Create product
	gotP1 := product.CreateProduct1(pRepo, t)
	_ = product.CreateProduct2(pRepo, t)

	ctx := context.Background()
	res, err := service.GetProductWithId(ctx, gotP1.ID)
	require.Nil(t, err)
	require.Equal(t, gotP1.Description, res.Product.Description)
}

func TestGateway_GetProductWithId_ZeroId(t *testing.T) {
	// Service mock
	service := NewServiceMock(t)

	// Product repo
	pRepo, err := product.NewProductRepoMock()
	require.Nil(t, err)

	// Create product
	_ = product.CreateProduct1(pRepo, t)

	ctx := context.Background()
	res, err := service.GetProductWithId(ctx, 0)
	require.Equal(t, derror.InvalidProduct, err)
	require.Nil(t, res)
}

func TestGateway_GetProductWithId_ProductNotExist(t *testing.T) {
	// Service mock
	service := NewServiceMock(t)

	// Product repo
	pRepo, err := product.NewProductRepoMock()
	require.Nil(t, err)

	// Create product
	gotP1 := product.CreateProduct1(pRepo, t)

	ctx := context.Background()
	res, err := service.GetProductWithId(ctx, gotP1.ID+100)
	require.Equal(t, derror.ProductNotFound, err)
	require.Nil(t, res)
}

func TestGateway_DeleteProduct_Ok(t *testing.T) {
	// Service mock
	service := NewServiceMock(t)

	// Product repo
	pRepo, err := product.NewProductRepoMock()
	require.Nil(t, err)

	// Create product
	gotP1 := product.CreateProduct1(pRepo, t)

	ctx := context.Background()
	err = service.DeleteProduct(ctx, gotP1.ID)
	require.Nil(t, err)
}

func TestGateway_DeleteProduct_ZeroId(t *testing.T) {
	// Service mock
	service := NewServiceMock(t)

	// Product repo
	pRepo, err := product.NewProductRepoMock()
	require.Nil(t, err)

	// Create product
	_ = product.CreateProduct1(pRepo, t)

	ctx := context.Background()
	err = service.DeleteProduct(ctx, 0)
	require.Equal(t, derror.InvalidProduct, err)
}

func TestGateway_DeleteProduct_ProductNotExist(t *testing.T) {
	// Service mock
	service := NewServiceMock(t)

	// Product repo
	pRepo, err := product.NewProductRepoMock()
	require.Nil(t, err)

	// Create product
	gotP1 := product.CreateProduct1(pRepo, t)

	ctx := context.Background()
	err = service.DeleteProduct(ctx, gotP1.ID+100)
	require.Equal(t, derror.ProductNotFound, err)
}

func TestGateway_EditProduct_Ok(t *testing.T) {
	// Service mock
	service := NewServiceMock(t)

	// Product repo
	pRepo, err := product.NewProductRepoMock()
	require.Nil(t, err)

	// Create product
	gotP1 := product.CreateProduct1(pRepo, t)
	p1 := GetProduct1()
	p1.Id = gotP1.ID

	ctx := context.Background()

	t.Run("description", func(t *testing.T) {
		p1.Description = "عوض شدن توضیحات"
		req := &api.EditProductRequest{Product: p1}

		res, err := service.EditProduct(ctx, req)
		require.Nil(t, err)
		require.NotNil(t, res)

		getRes, err := service.GetProductWithId(ctx, gotP1.ID)
		require.Nil(t, err)
		require.Equal(t, p1.Description, getRes.Product.Description)
	})

	t.Run("design code", func(t *testing.T) {
		p1.DesignCode = "106"
		req := &api.EditProductRequest{Product: p1}

		res, err := service.EditProduct(ctx, req)
		require.Nil(t, err)
		require.NotNil(t, res)

		getRes, err := service.GetProductWithId(ctx, gotP1.ID)
		require.Nil(t, err)
		require.Equal(t, p1.DesignCode, getRes.Product.DesignCode)
	})

	t.Run("size", func(t *testing.T) {
		p1.Sizes = []string{"12", "8"}
		req := &api.EditProductRequest{Product: p1}

		res, err := service.EditProduct(ctx, req)
		require.Nil(t, err)
		require.NotNil(t, res)

		getRes, err := service.GetProductWithId(ctx, gotP1.ID)
		require.Nil(t, err)
		require.Equal(t, p1.Sizes, getRes.Product.Sizes)
	})

	t.Run("color", func(t *testing.T) {
		p1.Colors = []string{"آبی", "صورتی"}
		req := &api.EditProductRequest{Product: p1}

		res, err := service.EditProduct(ctx, req)
		require.Nil(t, err)
		require.NotNil(t, res)

		getRes, err := service.GetProductWithId(ctx, gotP1.ID)
		require.Nil(t, err)
		require.Equal(t, p1.Colors, getRes.Product.Colors)
	})

}

func TestGateway_EditProduct_NotChange(t *testing.T) {
	// Service mock
	service := NewServiceMock(t)

	// Product repo
	pRepo, err := product.NewProductRepoMock()
	require.Nil(t, err)

	// Create product
	gotP1 := product.CreateProduct1(pRepo, t)
	p1 := GetProduct1()
	p1.Id = gotP1.ID

	ctx := context.Background()
	req := &api.EditProductRequest{Product: p1}

	res, err := service.EditProduct(ctx, req)
	require.Nil(t, err)
	require.NotNil(t, res)

	getRes, err := service.GetProductWithId(ctx, gotP1.ID)
	require.Nil(t, err)
	require.Equal(t, p1.Description, getRes.Product.Description)
}

func TestGateway_EditProduct_ProductNotExist(t *testing.T) {
	// Service mock
	service := NewServiceMock(t)

	// Product repo
	pRepo, err := product.NewProductRepoMock()
	require.Nil(t, err)

	// Create product
	gotP1 := product.CreateProduct1(pRepo, t)
	p1 := GetProduct1()
	p1.Id = gotP1.ID + 100

	ctx := context.Background()
	req := &api.EditProductRequest{Product: p1}

	res, err := service.EditProduct(ctx, req)
	require.Equal(t, derror.ProductNotFound, err)
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
