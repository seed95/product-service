package product

import (
	"github.com/seed95/product-service/internal/derror"
	"github.com/seed95/product-service/internal/repo/product/schema"
)

func (r *productRepo) AddColorToProduct(themes []schema.Theme) ([]schema.Theme, error) {
	//TODO copy slice for safe input slice from changed
	if err := r.db.Create(themes).Error; err != nil {
		return nil, derror.New(derror.InternalServer, err.Error())
	}
	return themes, nil
}

func (r *productRepo) DeleteColorInProduct(themes []schema.Theme) ([]schema.Theme, error) {

	panic("implement me")
	//schemaColors := schema.GetThemes(productId, colors)
	//
	//if err := r.db.Create(schemaColors).Error; err != nil {
	//	return nil, derror.New(derror.InternalServer, err.Error())
	//}
	//return nil, nil
}
