package api

import (
	"github.com/seed95/product-service/internal/model"
	"github.com/seed95/product-service/internal/repo/product/schema"
)

type Product struct {
	Id          uint     `json:"id"`
	CompanyId   uint     `json:"company_id"`
	CompanyName string   `json:"company_name"`
	DesignCode  string   `json:"design_code"`
	Description string   `json:"description"`
	Sizes       []string `json:"sizes"`
	Colors      []string `json:"colors"`
}

func ProductApiToModel(p Product) model.Product {
	return model.Product{
		Id:          p.Id,
		CompanyName: p.CompanyName,
		CompanyId:   p.CompanyId,
		DesignCode:  p.DesignCode,
		Colors:      p.Colors,
		Sizes:       p.Sizes,
		Description: p.Description,
	}
}

func ProductSchemaToApi(p schema.Product) Product {
	return Product{
		Id:          p.ID,
		CompanyId:   p.CompanyId,
		CompanyName: "",
		DesignCode:  p.DesignCode,
		Description: p.Description,
		Sizes:       schema.GetSizes(p.Dimensions),
		Colors:      schema.GetColors(p.Themes),
	}
}

type (
	CreateNewProductRequest struct {
		NewProduct Product `json:"new_product"`
	}

	CreateNewProductResponse struct {
		Products []Product `json:"products"`
	}
)
