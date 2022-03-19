package schema

import (
	"github.com/seed95/product-service/internal/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProductModelToSchema(t *testing.T) {
	p := model.Product{
		Id:          101,
		CompanyName: "Negin",
		CompanyId:   1,
		DesignCode:  "105",
		Colors:      []string{""},
		Sizes:       []string{"12", "9"},
		Description: "توضیحات کد ۱۰۵",
	}

	gotP := ProductModelToSchema(p)
	require.Equal(t, p.Id, gotP.ID)
	require.Equal(t, p.CompanyId, gotP.CompanyId)
	require.Equal(t, p.Description, gotP.Description)
	require.Equal(t, p.DesignCode, gotP.DesignCode)
	require.Equal(t, len(p.Sizes), len(gotP.Dimensions))
	require.Equal(t, len(p.Colors), len(gotP.Themes))
}

func TestGetSizes(t *testing.T) {

	tests := []struct {
		Name       string
		Dimensions []Dimension
		Sizes      []string
	}{
		{
			Name:       "Ok",
			Dimensions: []Dimension{{Size: "12"}, {Size: "6"}},
			Sizes:      []string{"12", "6"},
		},
		{
			Name:       "OneDimension",
			Dimensions: []Dimension{{Size: "12"}},
			Sizes:      []string{"12"},
		},
		{
			Name:       "EmptySize",
			Dimensions: []Dimension{{Size: ""}},
			Sizes:      []string{""},
		},
		{
			Name:       "EmptySlice",
			Dimensions: []Dimension{},
			Sizes:      []string{},
		},
		{
			Name:       "DuplicateSize",
			Dimensions: []Dimension{{Size: "12"}, {Size: "12"}},
			Sizes:      []string{"12", "12"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			sizes := GetSizes(tt.Dimensions)
			require.Equal(t, tt.Sizes, sizes)
		})
	}
}

func TestGetColors(t *testing.T) {

	tests := []struct {
		Name   string
		Themes []Theme
		Colors []string
	}{
		{
			Name:   "Ok",
			Themes: []Theme{{Color: "آبی"}, {Color: "قرمز"}},
			Colors: []string{"آبی", "قرمز"},
		},
		{
			Name:   "OneTheme",
			Themes: []Theme{{Color: "آبی"}},
			Colors: []string{"آبی"},
		},
		{
			Name:   "EmptyColor",
			Themes: []Theme{{Color: ""}},
			Colors: []string{""},
		},
		{
			Name:   "EmptySlice",
			Themes: []Theme{},
			Colors: []string{},
		},
		{
			Name:   "DuplicateColor",
			Themes: []Theme{{Color: "آبی"}, {Color: "آبی"}},
			Colors: []string{"آبی", "آبی"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			sizes := GetColors(tt.Themes)
			require.Equal(t, tt.Colors, sizes)
		})
	}
}
