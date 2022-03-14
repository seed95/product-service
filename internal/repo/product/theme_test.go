package product

import (
	"github.com/seed95/product-service/internal/model"
	"github.com/seed95/product-service/internal/repo/product/schema"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestThemeRepo_AddColorToProduct_Ok(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	companyId := uint(1)
	p1 := model.Product{
		CompanyName: "Negin",
		CompanyId:   companyId,
		DesignCode:  "105",
		Colors:      []string{"قرمز", "آبی"},
		Dimensions:  []string{"6", "9"},
		Description: "توضیحات برای کد ۱۰۵",
	}
	p, err := repo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, p)

	t1 := schema.Theme{
		ProductId: p.ID,
		Color:     "سبز",
	}
	gotThemes, err := repo.AddColorToProduct([]schema.Theme{t1})
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
}

func TestThemeRepo_AddColorToProduct_ProductNotExist(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	productId := uint(34)
	t1 := schema.Theme{
		ProductId: productId,
		Color:     "آبی",
	}
	gotThemes, err := repo.AddColorToProduct([]schema.Theme{t1})
	require.Nil(t, gotThemes)
	require.NotNil(t, err)
}

func TestThemeRepo_AddColorToProduct_Duplicate(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	companyId := uint(1)
	p1 := model.Product{
		CompanyName: "Negin",
		CompanyId:   companyId,
		DesignCode:  "105",
		Colors:      []string{"قرمز", "آبی"},
		Dimensions:  []string{"6", "9"},
		Description: "توضیحات برای کد ۱۰۵",
	}
	gotP1, err := repo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, gotP1)

	p2 := model.Product{
		CompanyName: "Negin",
		CompanyId:   companyId,
		DesignCode:  "106",
		Colors:      []string{"قرمز", "آبی"},
		Dimensions:  []string{"6", "9"},
		Description: "توضیحات برای کد ۱۰۶",
	}
	gotP2, err := repo.CreateProduct(p2)
	require.Nil(t, err)
	require.NotNil(t, gotP2)

	t1 := schema.Theme{
		ProductId: gotP1.ID,
		Color:     p1.Colors[0],
	}
	gotThemes, err := repo.AddColorToProduct([]schema.Theme{t1})
	require.Nil(t, gotThemes)
	require.NotNil(t, err)

	t2 := schema.Theme{
		ProductId: gotP1.ID,
		Color:     "سبز",
	}
	gotThemes, err = repo.AddColorToProduct([]schema.Theme{t2, t1})
	require.Nil(t, gotThemes)
	require.NotNil(t, err)

	t3 := schema.Theme{
		ProductId: gotP2.ID,
		Color:     "سبز",
	}
	gotThemes, err = repo.AddColorToProduct([]schema.Theme{t3, t1})
	require.Nil(t, gotThemes)
	require.NotNil(t, err)
}

func TestThemeRepo_AddColorToProduct_MultiProduct(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	companyId := uint(1)
	p1 := model.Product{
		CompanyName: "Negin",
		CompanyId:   companyId,
		DesignCode:  "105",
		Colors:      []string{"قرمز", "آبی"},
		Dimensions:  []string{"6", "9"},
		Description: "توضیحات برای کد ۱۰۵",
	}
	gotP1, err := repo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, gotP1)

	p1 = model.Product{
		CompanyName: "Negin",
		CompanyId:   companyId,
		DesignCode:  "106",
		Colors:      []string{"قرمز"},
		Dimensions:  []string{"12"},
		Description: "توضیحات برای کد ۱۰۶",
	}
	gotP2, err := repo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, gotP2)

	t1 := schema.Theme{
		ProductId: gotP1.ID,
		Color:     "سبز",
	}
	t2 := schema.Theme{
		ProductId: gotP2.ID,
		Color:     "سبز",
	}
	gotThemes, err := repo.AddColorToProduct([]schema.Theme{t1, t2})
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
}
