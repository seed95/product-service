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
		Dimensions:  []string{"12", "9"},
		Description: "توضیحات کد ۱۰۵",
	}

	gotP := ProductModelToSchema(p)
	require.Equal(t, p.Id, gotP.ID)
	require.Equal(t, p.CompanyId, gotP.CompanyId)
	require.Equal(t, p.Description, gotP.Description)
	require.Equal(t, p.DesignCode, gotP.DesignCode)
	require.Equal(t, len(p.Dimensions), len(gotP.Dimensions))
	require.Equal(t, len(p.Colors), len(gotP.Themes))
}
