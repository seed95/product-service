package product

import (
	"github.com/seed95/product-service/internal/derror"
	"github.com/seed95/product-service/internal/model"
	"github.com/seed95/product-service/internal/repo/product/schema"
	"strconv"
)

// GetAllCarpet return all carpets for `companyId` in view
func (r *productRepo) GetAllCarpet(companyId uint) ([]model.Carpet, error) {
	var schemaCarpets []schema.Carpet
	viewName := "view_carpet_company_id_" + strconv.FormatUint(uint64(companyId), 10)
	if err := r.db.Table(viewName).Find(&schemaCarpets).Error; err != nil {
		return nil, derror.New(derror.InternalServer, err.Error())
	}

	carpets := make([]model.Carpet, len(schemaCarpets))
	for i, c := range schemaCarpets {
		carpets[i] = schema.CarpetToModel(&c, companyId)
	}

	return carpets, nil
}

// GetAllCarpetWithProductId return all carpets for `companyId` and `productId` in view
func (r *productRepo) GetAllCarpetWithProductId(companyId, productId uint) ([]model.Carpet, error) {

	var schemaCarpets []schema.Carpet
	viewName := "view_carpet_company_id_" + strconv.FormatUint(uint64(companyId), 10)
	if err := r.db.Table(viewName).Where("product_id = ?", productId).Find(&schemaCarpets).Error; err != nil {
		return nil, derror.New(derror.InternalServer, err.Error())
	}

	carpets := make([]model.Carpet, len(schemaCarpets))
	for i, c := range schemaCarpets {
		carpets[i] = schema.CarpetToModel(&c, companyId)
	}
	return carpets, nil
}
