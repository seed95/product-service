package model

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCarpetsToProduct(t *testing.T) {

	c1 := Carpet{
		Id:          "P12D14C13",
		CompanyId:   1,
		ProductId:   12,
		DimensionId: 14,
		ThemeId:     13,
		DesignCode:  "106",
		Dimension:   "6",
		Color:       "آبی",
	}

	c2 := Carpet{
		Id:          "P12D14C14",
		CompanyId:   1,
		ProductId:   12,
		DimensionId: 14,
		ThemeId:     14,
		DesignCode:  "106",
		Dimension:   "6",
		Color:       "قرمز",
	}

	c3 := Carpet{
		Id:          "P12D13C14",
		CompanyId:   1,
		ProductId:   12,
		DimensionId: 13,
		ThemeId:     14,
		DesignCode:  "106",
		Dimension:   "9",
		Color:       "قرمز",
	}

	c4 := Carpet{
		Id:          "P12D13C13",
		CompanyId:   1,
		ProductId:   12,
		DimensionId: 13,
		ThemeId:     13,
		DesignCode:  "106",
		Dimension:   "9",
		Color:       "آبی",
	}

	carpets := []Carpet{c1, c2, c3, c4}

	product, err := CarpetsToProduct(carpets)
	require.Nil(t, err)
	require.NotNil(t, product)

}
