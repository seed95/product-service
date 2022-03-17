package schema

import (
	"fmt"
	"gorm.io/gorm"
)

type (
	Dimension struct {
		gorm.Model
		ProductId uint   `gorm:"uniqueIndex:dimension_unique_id"`
		Size      string `gorm:"uniqueIndex:dimension_unique_id"`
	}
)

func (d Dimension) String() string {
	return fmt.Sprintf("ID: %v, ProductId: %v, Size: %v", d.ID, d.ProductId, d.Size)
}
