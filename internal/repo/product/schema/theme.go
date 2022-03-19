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
	return fmt.Sprintf("ID: %v, Id: %v, Color: %v", t.ID, t.ProductId, t.Color)
}

func GetColors(themes []Theme) []string {
	result := make([]string, len(themes))
	for i, t := range themes {
		result[i] = t.Color
	}
	return result
}
