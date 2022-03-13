package repo

import (
	"github.com/seed95/OrderManagement/Microservice/product-service/internal/model"
)

type ProductRepo interface {
	CreateProduct(p *model.Product) (*model.Product, error)
	GetAllCarpet(companyId uint) ([]model.Carpet, error)
	DeleteProduct(productId uint) error
}
