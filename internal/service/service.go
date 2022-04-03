package service

import (
	"context"
	"fmt"
	"github.com/seed95/product-service/internal/api"
	"github.com/seed95/product-service/internal/derror"
	"github.com/seed95/product-service/internal/model"
	"github.com/seed95/product-service/internal/repo"
	"github.com/seed95/product-service/pkg/logger"
	"github.com/seed95/product-service/pkg/logger/keyval"
	"github.com/seed95/product-service/pkg/unique"
)

type ProductService interface {
	CreateNewProduct(ctx context.Context, req *api.CreateNewProductRequest) (res *api.GetAllProductsResponse, err error)
	GetAllProducts(ctx context.Context, companyId uint) (res *api.GetAllProductsResponse, err error)
	GetProductWithId(ctx context.Context, productId uint) (res *api.GetProductResponse, err error)
	DeleteProduct(ctx context.Context, productId uint) (err error)
	EditProduct(ctx context.Context, req *api.EditProductRequest) (res *api.EditProductResponse, err error)
}

type (
	gateway struct {
		product repo.ProductRepo
		logger  logger.Logger
	}

	Setting struct {
		ProductRepo repo.ProductRepo
		Logger      logger.Logger
	}
)

var _ ProductService = (*gateway)(nil)

func New(s *Setting) (ProductService, error) {
	return &gateway{product: s.ProductRepo, logger: s.Logger}, nil

}

func (g *gateway) CreateNewProduct(ctx context.Context, req *api.CreateNewProductRequest) (res *api.GetAllProductsResponse, err error) {
	// Log request response
	defer func() {
		commonKeyVal := []keyval.Pair{
			keyval.String("req", fmt.Sprintf("%+v", req)),
			keyval.String("res", fmt.Sprintf("%+v", res)),
		}
		logger.LogReqRes(g.logger, "service.CreateNewProduct", err, commonKeyVal...)
	}()

	modelProduct := api.ProductApiToModel(*req.NewProduct)
	if !productIsValid(*modelProduct) || modelProduct.Id != 0 {
		return nil, derror.InvalidProduct
	}

	_, err = g.product.CreateProduct(*modelProduct)
	if err != nil {
		return nil, err
	}

	return g.GetAllProducts(ctx, modelProduct.CompanyId)
}

func (g *gateway) GetAllProducts(ctx context.Context, companyId uint) (res *api.GetAllProductsResponse, err error) {
	// Log request response
	defer func() {
		commonKeyVal := []keyval.Pair{
			keyval.String("company_id", fmt.Sprintf("%v", companyId)),
			keyval.String("res", fmt.Sprintf("%+v", res)),
		}
		logger.LogReqRes(g.logger, "service.GetAllProducts", err, commonKeyVal...)
	}()

	if companyId == 0 {
		return nil, derror.InvalidCompany
	}

	allProducts, err := g.product.GetAllProducts(companyId)
	if err != nil {
		return nil, err
	}

	res = &api.GetAllProductsResponse{}
	res.Products = make([]api.Product, len(allProducts))
	for i, p := range allProducts {
		res.Products[i] = *api.ProductSchemaToApi(p)
	}
	return res, nil
}

func (g *gateway) GetProductWithId(ctx context.Context, productId uint) (res *api.GetProductResponse, err error) {
	// Log request response
	defer func() {
		commonKeyVal := []keyval.Pair{
			keyval.String("product_id", fmt.Sprintf("%v", productId)),
			keyval.String("res", fmt.Sprintf("%+v", res)),
		}
		logger.LogReqRes(g.logger, "service.GetProductWithId", err, commonKeyVal...)
	}()

	if productId == 0 {
		return nil, derror.InvalidProduct
	}

	schemaProduct, err := g.product.GetProductWithId(productId)
	if err != nil {
		return nil, err
	}

	res = &api.GetProductResponse{}
	res.Product = api.ProductSchemaToApi(*schemaProduct)
	return res, nil
}

func (g *gateway) DeleteProduct(ctx context.Context, productId uint) (err error) {
	// Log request response
	defer func() {
		commonKeyVal := []keyval.Pair{
			keyval.String("product_id", fmt.Sprintf("%v", productId)),
		}
		logger.LogReqRes(g.logger, "service.DeleteProduct", err, commonKeyVal...)
	}()

	if productId == 0 {
		return derror.InvalidProduct
	}
	return g.product.DeleteProduct(productId)
}

func (g *gateway) EditProduct(ctx context.Context, req *api.EditProductRequest) (res *api.EditProductResponse, err error) {
	// Log request response
	defer func() {
		commonKeyVal := []keyval.Pair{
			keyval.String("req", fmt.Sprintf("%+v", req)),
			keyval.String("res", fmt.Sprintf("%+v", res)),
		}
		logger.LogReqRes(g.logger, "service.CreateNewProduct", err, commonKeyVal...)
	}()

	modelProduct := api.ProductApiToModel(*req.EditedProduct)
	if !productIsValid(*modelProduct) || modelProduct.Id == 0 {
		return nil, derror.InvalidProduct
	}

	editedProduct, err := g.product.EditProduct(*modelProduct)
	if err != nil {
		return nil, err
	}

	res = &api.EditProductResponse{}
	res.NewProduct = api.ProductSchemaToApi(*editedProduct)
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
