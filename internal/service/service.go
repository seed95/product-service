package service

import (
	"context"
	"github.com/seed95/product-service/internal/api"
	"github.com/seed95/product-service/internal/derror"
	"github.com/seed95/product-service/internal/model"
	"github.com/seed95/product-service/internal/repo"
	"github.com/seed95/product-service/pkg/unique"
)

type ProductService interface {
	CreateNewProduct(ctx context.Context, req *api.CreateNewProductRequest) (res *api.CreateNewProductResponse, err error)
	GetAllProducts(ctx context.Context, companyId uint) (res *api.GetAllProductsResponse, err error)
}

type (
	gateway struct {
		productRepo repo.ProductRepo
	}

	Setting struct {
		ProductRepo repo.ProductRepo
	}
)

var _ ProductService = (*gateway)(nil)

func New(s *Setting) (ProductService, error) {
	return &gateway{productRepo: s.ProductRepo}, nil

}

func (g gateway) CreateNewProduct(ctx context.Context, req *api.CreateNewProductRequest) (res *api.CreateNewProductResponse, err error) {
	res = &api.CreateNewProductResponse{}

	modelProduct := api.ProductApiToModel(req.NewProduct)
	if !productIsValid(modelProduct) || modelProduct.Id != 0 {
		return nil, derror.InvalidProduct
	}

	_, err = g.productRepo.CreateProduct(modelProduct)
	if err != nil {
		return nil, err
	}

	allProducts, err := g.productRepo.GetAllProducts(modelProduct.CompanyId)
	if err != nil {
		return nil, err
	}
	res.Products = make([]api.Product, len(allProducts))
	for i, p := range allProducts {
		res.Products[i] = api.ProductSchemaToApi(p)
	}
	return res, nil
}

func (g gateway) GetAllProducts(ctx context.Context, companyId uint) (res *api.GetAllProductsResponse, err error) {
	res = &api.GetAllProductsResponse{}

	allProducts, err := g.productRepo.GetAllProducts(companyId)
	if err != nil {
		return nil, err
	}
	res.Products = make([]api.Product, len(allProducts))
	for i, p := range allProducts {
		res.Products[i] = api.ProductSchemaToApi(p)
	}
	return res, nil
}

func productIsValid(p model.Product) bool {
	// Check empty color
	for _, c := range p.Colors {
		if c == "" {
			return false
		}
	}

	// Check empty size
	for _, s := range p.Sizes {
		if s == "" {
			return false
		}
	}

	// Check unique color
	if !unique.StringsAreUnique(p.Colors) {
		return false
	}

	// Check unique size
	if !unique.StringsAreUnique(p.Sizes) {
		return false
	}

	return p.DesignCode != "" && len(p.Colors) != 0 && len(p.Sizes) != 0 && p.CompanyId != 0
}
