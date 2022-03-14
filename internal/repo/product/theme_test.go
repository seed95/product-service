package product

import (
	"github.com/seed95/product-service/internal/derror"
	"github.com/seed95/product-service/internal/model"
	"github.com/seed95/product-service/internal/repo/product/schema"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"testing"
)

func TestThemeRepo_AddColorsToProduct_Ok(t *testing.T) {
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

	colors := []string{"سبز", "نارنجی", "green"}
	gotThemes, err := repo.AddColorsToProduct(p.ID, colors)
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
}

func TestThemeRepo_AddColorsToProduct_ProductNotExist(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	colors := []string{"سبز"}
	gotThemes, err := repo.AddColorsToProduct(34, colors)
	require.Nil(t, gotThemes)
	require.NotNil(t, err)
}

func TestThemeRepo_AddColorsToProduct_Duplicate(t *testing.T) {
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

	colors := []string{p1.Colors[0]}
	gotThemes, err := repo.AddColorsToProduct(gotP1.ID, colors)
	require.Nil(t, gotThemes)
	require.NotNil(t, err)

	t.Run("roll back", func(t *testing.T) {
		colors = []string{"سبز", p1.Colors[0]}
		gotThemes, err = repo.AddColorsToProduct(gotP1.ID, colors)
		require.Nil(t, gotThemes)
		require.NotNil(t, err)
		gotP1, err = repo.GetProductWithId(gotP1.ID)
		require.Equal(t, p1.Colors[0], gotP1.Themes[0].Color)
	})

}

func TestThemeRepo_AddColorsToProduct_Empty(t *testing.T) {
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

	gotThemes, err := repo.AddColorsToProduct(gotP1.ID, []string{})
	require.Nil(t, gotThemes)
	require.Equal(t, derror.InvalidColor, err)
}

func TestThemeRepo_AddColorsToProduct_Nil(t *testing.T) {
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

	gotThemes, err := repo.AddColorsToProduct(gotP1.ID, nil)
	require.Nil(t, gotThemes)
	require.Equal(t, derror.InvalidColor, err)
}

func TestThemeRepo_DeleteColorsInProduct_Ok(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	companyId := uint(1)
	p1 := model.Product{
		CompanyName: "Negin",
		CompanyId:   companyId,
		DesignCode:  "105",
		Colors:      []string{"قرمز", "آبی", "سبز"},
		Dimensions:  []string{"6", "9"},
		Description: "توضیحات برای کد ۱۰۵",
	}
	p, err := repo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, p)

	colors := []string{"قرمز", "آبی"}
	err = repo.DeleteColorsInProduct(p.ID, colors)
	require.Nil(t, err)
}

func TestThemeRepo_DeleteColorsInProduct_ProductNotExist(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	colors := []string{"سبز"}
	err = repo.DeleteColorsInProduct(34, colors)
	require.Equal(t, err, derror.ColorNotFound)
}

func TestThemeRepo_DeleteColorsInProduct_ColorNotExist(t *testing.T) {
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

	colors := []string{"سبز"}
	err = repo.DeleteColorsInProduct(gotP1.ID, colors)
	require.Equal(t, err, derror.ColorNotFound)

	colors = []string{"سبز", "آبی"}
	err = repo.DeleteColorsInProduct(gotP1.ID, colors)
	require.Equal(t, err, derror.ColorNotFound)
	gotP1, err = repo.GetProductWithId(gotP1.ID)
	require.Equal(t, len(p1.Colors), len(gotP1.Themes))

	colors = []string{"آبی", "سبز"}
	err = repo.DeleteColorsInProduct(gotP1.ID, colors)
	require.Equal(t, err, derror.ColorNotFound)
	gotP1, err = repo.GetProductWithId(gotP1.ID)
	require.Equal(t, len(p1.Colors), len(gotP1.Themes))
}

func TestThemeRepo_DeleteColorsInProduct_Empty(t *testing.T) {
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

	err = repo.DeleteColorsInProduct(gotP1.ID, []string{})
	require.Nil(t, err)
}

func TestThemeRepo_DeleteColorsInProduct_Nil(t *testing.T) {
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

	err = repo.DeleteColorsInProduct(gotP1.ID, nil)
	require.Nil(t, err)
}

func TestThemeRepo_UpdateColorsWithId_Ok(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	companyId := uint(1)
	p1 := model.Product{
		CompanyName: "Negin",
		CompanyId:   companyId,
		DesignCode:  "105",
		Colors:      []string{"قرمز", "آبی", "سبز"},
		Dimensions:  []string{"6", "9"},
		Description: "توضیحات برای کد ۱۰۵",
	}
	p, err := repo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, p)

	themes := []schema.Theme{
		{
			Model:     gorm.Model{ID: p.Themes[0].ID}, //قرمز
			ProductId: p.ID,
			Color:     "نارنجی",
		},
		{
			Model:     gorm.Model{ID: p.Themes[2].ID}, //سبز
			ProductId: p.ID,
			Color:     "سبز",
		},
	}
	err = repo.UpdateColorsWithId(themes)
	require.Nil(t, err)
}

func TestThemeRepo_UpdateColorsWithId_ProductNotExist(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	companyId := uint(1)
	p1 := model.Product{
		CompanyName: "Negin",
		CompanyId:   companyId,
		DesignCode:  "105",
		Colors:      []string{"قرمز", "آبی", "سبز"},
		Dimensions:  []string{"6", "9"},
		Description: "توضیحات برای کد ۱۰۵",
	}
	p, err := repo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, p)

	themes := []schema.Theme{
		{
			Model:     gorm.Model{ID: p.Themes[0].ID}, //قرمز
			ProductId: 100,
			Color:     "نارنجی",
		},
	}
	err = repo.UpdateColorsWithId(themes)
	require.NotNil(t, err)
}

func TestThemeRepo_UpdateColorsWithId_ColorNotExist(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("one color", func(t *testing.T) {
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

		themes := []schema.Theme{
			{
				Model:     gorm.Model{ID: 100},
				ProductId: gotP1.ID,
				Color:     "نارنجی",
			},
		}
		err = repo.UpdateColorsWithId(themes)
		require.Equal(t, err, derror.ColorNotFound)
	})

	t.Run("two color", func(t *testing.T) {
		companyId := uint(1)
		p1 := model.Product{
			CompanyName: "Negin",
			CompanyId:   companyId,
			DesignCode:  "106",
			Colors:      []string{"قرمز", "آبی"},
			Dimensions:  []string{"6", "9"},
			Description: "توضیحات برای کد ۱۰۶",
		}
		gotP1, err := repo.CreateProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, gotP1)

		themes := []schema.Theme{
			{
				Model:     gorm.Model{ID: 100},
				ProductId: gotP1.ID,
				Color:     "صورتی",
			},
			{
				Model:     gorm.Model{ID: gotP1.Themes[0].ID}, //قرمز
				ProductId: gotP1.ID,
				Color:     "نارنجی",
			},
		}
		err = repo.UpdateColorsWithId(themes)
		require.Equal(t, err, derror.ColorNotFound)
		gotP1, err = repo.GetProductWithId(gotP1.ID)
		require.Equal(t, p1.Colors[0], gotP1.Themes[0].Color)
	})

	t.Run("roll back", func(t *testing.T) {
		companyId := uint(1)
		p1 := model.Product{
			CompanyName: "Negin",
			CompanyId:   companyId,
			DesignCode:  "107",
			Colors:      []string{"قرمز", "آبی"},
			Dimensions:  []string{"6", "9"},
			Description: "توضیحات برای کد ۱۰۷",
		}
		gotP1, err := repo.CreateProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, gotP1)

		themes := []schema.Theme{
			{
				Model:     gorm.Model{ID: gotP1.Themes[0].ID}, //قرمز
				ProductId: gotP1.ID,
				Color:     "نارنجی",
			},
			{
				Model:     gorm.Model{ID: 100},
				ProductId: gotP1.ID,
				Color:     "صورتی",
			},
		}
		err = repo.UpdateColorsWithId(themes)
		require.Equal(t, err, derror.ColorNotFound)
		gotP1, err = repo.GetProductWithId(gotP1.ID)
		require.Equal(t, p1.Colors[0], gotP1.Themes[0].Color)
	})
}

func TestThemeRepo_UpdateColorsWithId_MultiProduct(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	companyId := uint(1)
	p1 := model.Product{
		CompanyName: "Negin",
		CompanyId:   companyId,
		DesignCode:  "105",
		Colors:      []string{"قرمز", "آبی", "سبز"},
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

	themes := []schema.Theme{
		{
			Model:     gorm.Model{ID: gotP1.Themes[0].ID}, //قرمز
			ProductId: gotP1.ID,
			Color:     "نارنجی",
		},
		{
			Model:     gorm.Model{ID: gotP2.Themes[1].ID}, //آبی
			ProductId: gotP2.ID,
			Color:     "سبز",
		},
	}
	err = repo.UpdateColorsWithId(themes)
	require.Nil(t, err)
	gotP1, err = repo.GetProductWithId(gotP1.ID)
	require.Equal(t, themes[0].Color, gotP1.Themes[0].Color)
	gotP2, err = repo.GetProductWithId(gotP2.ID)
	require.Equal(t, themes[1].Color, gotP2.Themes[1].Color)
}

func TestThemeRepo_UpdateColorsWithId_Empty(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	companyId := uint(1)
	p1 := model.Product{
		CompanyName: "Negin",
		CompanyId:   companyId,
		DesignCode:  "105",
		Colors:      []string{"قرمز", "آبی", "سبز"},
		Dimensions:  []string{"6", "9"},
		Description: "توضیحات برای کد ۱۰۵",
	}
	p, err := repo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, p)

	err = repo.UpdateColorsWithId([]schema.Theme{})
	require.Nil(t, err)
}

func TestThemeRepo_UpdateColorsWithId_Nil(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	companyId := uint(1)
	p1 := model.Product{
		CompanyName: "Negin",
		CompanyId:   companyId,
		DesignCode:  "105",
		Colors:      []string{"قرمز", "آبی", "سبز"},
		Dimensions:  []string{"6", "9"},
		Description: "توضیحات برای کد ۱۰۵",
	}
	p, err := repo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, p)

	err = repo.UpdateColorsWithId(nil)
	require.Nil(t, err)
}
