package service

import (
	"context"
	"fmt"
	"github.com/seed95/product-service/internal/api"
	"github.com/seed95/product-service/internal/derror"
	"github.com/seed95/product-service/internal/model"
	"github.com/seed95/product-service/internal/repo"
	kitlog "github.com/seed95/product-service/pkg/logger"
	"github.com/seed95/product-service/pkg/logger/keyval"
	"github.com/seed95/product-service/pkg/unique"
)

type ProductService interface {
	CreateNewProduct(ctx context.Context, req api.CreateNewProductRequest) (res *api.GetAllProductsResponse, err error)
	GetAllProducts(ctx context.Context, companyId uint) (res *api.GetAllProductsResponse, err error)
	GetProductWithId(ctx context.Context, productId uint) (res *api.GetProductResponse, err error)
	DeleteProduct(ctx context.Context, productId uint) (err error)
	EditProduct(ctx context.Context, req *api.EditProductRequest) (res *api.EditProductResponse, err error)
}

type (
	gateway struct {
		product repo.ProductRepo
		logger  kitlog.Logger
	}

	Setting struct {
		ProductRepo repo.ProductRepo
		Logger      kitlog.Logger
	}
)

var _ ProductService = (*gateway)(nil)

func New(s *Setting) (ProductService, error) {
	return &gateway{product: s.ProductRepo, logger: s.Logger}, nil

}

func (g *gateway) CreateNewProduct(ctx context.Context, req api.CreateNewProductRequest) (res *api.GetAllProductsResponse, err error) {
	// Log request response
	defer func() {
		commonKeyVal := []keyval.Pair{
			keyval.String("req.Product", fmt.Sprintf("%+v", req.Product)),
			keyval.String("res", fmt.Sprintf("%+v", res)),
		}
		kitlog.LogReqRes(g.logger, "service.CreateNewProduct", err, commonKeyVal...)
	}()

	modelProduct := api.ProductApiToModel(req.Product)
	if err := productIsValid(*modelProduct); err != nil {
		return nil, err
	}

	if modelProduct.Id != 0 {
		return nil, derror.New(derror.InvalidProduct, "invalid product id")
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
		kitlog.LogReqRes(g.logger, "service.GetAllProducts", err, commonKeyVal...)
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
		kitlog.LogReqRes(g.logger, "service.GetProductWithId", err, commonKeyVal...)
	}()

	if productId == 0 {
		return nil, derror.InvalidProduct
	}

	schemaProduct, err := g.product.GetProductWithId(productId)
	if err != nil {
		return nil, err
	}

	res = &api.GetProductResponse{}
	res.Product = *api.ProductSchemaToApi(*schemaProduct)
	return res, nil
}

func (g *gateway) DeleteProduct(ctx context.Context, productId uint) (err error) {
	// Log request response
	defer func() {
		commonKeyVal := []keyval.Pair{
			keyval.String("product_id", fmt.Sprintf("%v", productId)),
		}
		kitlog.LogReqRes(g.logger, "service.DeleteProduct", err, commonKeyVal...)
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
		kitlog.LogReqRes(g.logger, "service.CreateNewProduct", err, commonKeyVal...)
	}()

	modelProduct := api.ProductApiToModel(req.Product)
	if err := productIsValid(*modelProduct); err != nil {
		return nil, err
	}

	if modelProduct.Id == 0 {
		return nil, derror.New(derror.InvalidProduct, "invalid product id")
	}

	editedProduct, err := g.product.EditProduct(*modelProduct)
	if err != nil {
		return nil, err
	}

	res = &api.EditProductResponse{}
	res.Product = *api.ProductSchemaToApi(*editedProduct)
	return res, nil
}

func productIsValid(p model.Product) error {

	// Check empty color
	if len(p.Colors) == 0 {
		return derror.New(derror.InvalidProduct, "empty color")
	}

	// Check empty color
	for _, c := range p.Colors {
		if c == "" {
			return derror.New(derror.InvalidProduct, "empty color")
		}
	}

	// Check empty size
	if len(p.Sizes) == 0 {
		return derror.New(derror.InvalidProduct, "empty size")
	}

	// Check empty size
	for _, s := range p.Sizes {
		if s == "" {
			return derror.New(derror.InvalidProduct, "empty size")
		}
	}

	// Check unique color
	if !unique.StringsAreUnique(p.Colors) {
		return derror.New(derror.InvalidProduct, "not unique color")
	}

	// Check unique size
	if !unique.StringsAreUnique(p.Sizes) {
		return derror.New(derror.InvalidProduct, "not unique size")
	}

	if p.DesignCode != "" {
		return derror.New(derror.InvalidProduct, "empty design code")
	}

	if p.CompanyId != 0 {
		return derror.New(derror.InvalidProduct, "invalid company id")
	}

	return nil
}
