package schema

import (
	"fmt"
	"gorm.io/gorm"
)

type (
	Theme struct {
		gorm.Model
		ProductId uint   `gorm:"uniqueIndex:theme_unique_id"`
		Color     string `gorm:"uniqueIndex:theme_unique_id"`
	}
)

func (t Theme) String() string {
	return fmt.Sprintf("ID: %v, ProductId: %v, Color: %v", t.ID, t.ProductId, t.Color)
}
