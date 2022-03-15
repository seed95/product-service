package product

import (
	"github.com/seed95/product-service/internal/derror"
	"github.com/seed95/product-service/internal/model"
	"github.com/seed95/product-service/internal/repo/product/schema"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"testing"
)

func TestThemeRepo_AddThemesWithColor_Ok(t *testing.T) {
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
	gotThemes, err := repo.AddThemesWithColor(p.ID, colors)
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
}

func TestThemeRepo_AddThemesWithColor_ProductNotExist(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	colors := []string{"سبز"}
	gotThemes, err := repo.AddThemesWithColor(34, colors)
	require.Nil(t, gotThemes)
	require.NotNil(t, err)
}

func TestThemeRepo_AddThemesWithColor_Duplicate(t *testing.T) {
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
	gotThemes, err := repo.AddThemesWithColor(gotP1.ID, colors)
	require.Nil(t, gotThemes)
	require.NotNil(t, err)

	t.Run("roll back", func(t *testing.T) {
		colors = []string{"سبز", p1.Colors[0]}
		gotThemes, err = repo.AddThemesWithColor(gotP1.ID, colors)
		require.Nil(t, gotThemes)
		require.NotNil(t, err)
		gotP1, err = repo.GetProductWithId(gotP1.ID)
		require.Equal(t, p1.Colors[0], gotP1.Themes[0].Color)
	})

}

func TestThemeRepo_AddThemesWithColor_Empty(t *testing.T) {
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

	gotThemes, err := repo.AddThemesWithColor(gotP1.ID, []string{})
	require.Nil(t, gotThemes)
	require.Equal(t, derror.InvalidColor, err)
}

func TestThemeRepo_AddThemesWithColor_Nil(t *testing.T) {
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

	gotThemes, err := repo.AddThemesWithColor(gotP1.ID, nil)
	require.Nil(t, gotThemes)
	require.Equal(t, derror.InvalidColor, err)
}

func TestThemeRepo_DeleteThemesWithColor_Ok(t *testing.T) {
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
	err = repo.DeleteThemesWithColor(p.ID, colors)
	require.Nil(t, err)
}

func TestThemeRepo_DeleteThemesWithColor_ProductNotExist(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	colors := []string{"سبز"}
	err = repo.DeleteThemesWithColor(34, colors)
	require.Equal(t, derror.ThemeNotFound, err)
}

func TestThemeRepo_DeleteThemesWithColor_ColorNotExist(t *testing.T) {
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
	err = repo.DeleteThemesWithColor(gotP1.ID, colors)
	require.Equal(t, derror.ThemeNotFound, err)

	colors = []string{"سبز", "آبی"}
	err = repo.DeleteThemesWithColor(gotP1.ID, colors)
	require.Equal(t, derror.ThemeNotFound, err)
	gotP1, err = repo.GetProductWithId(gotP1.ID)
	require.Equal(t, len(p1.Colors), len(gotP1.Themes))

	colors = []string{"آبی", "سبز"}
	err = repo.DeleteThemesWithColor(gotP1.ID, colors)
	require.Equal(t, derror.ThemeNotFound, err)
	gotP1, err = repo.GetProductWithId(gotP1.ID)
	require.Equal(t, len(p1.Colors), len(gotP1.Themes))
}

func TestThemeRepo_DeleteThemesWithColor_Empty(t *testing.T) {
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

	err = repo.DeleteThemesWithColor(gotP1.ID, []string{})
	require.NotNil(t, err)
	require.Equal(t, derror.InvalidColor, err)
}

func TestThemeRepo_DeleteThemesWithColor_Nil(t *testing.T) {
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

	err = repo.DeleteThemesWithColor(gotP1.ID, nil)
	require.NotNil(t, err)
	require.Equal(t, derror.InvalidColor, err)
}

func TestThemeRepo_DeleteThemesWithId_Ok(t *testing.T) {
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

	colorIds := []uint{gotP1.Themes[0].ID, gotP1.Themes[2].ID}
	err = repo.DeleteThemesWithId(gotP1.ID, colorIds)
	require.Nil(t, err)
	gotThemes, err := repo.GetThemesWithProductId(gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
	require.Equal(t, 1, len(gotThemes))
	require.Equal(t, gotP1.Themes[1].Color, gotThemes[0].Color)
}

func TestThemeRepo_DeleteThemesWithId_ProductNotExist(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	colorIds := []uint{100}
	err = repo.DeleteThemesWithId(34, colorIds)
	require.Equal(t, derror.ThemeNotFound, err)
}

func TestThemeRepo_DeleteThemesWithId_ColorNotExist(t *testing.T) {
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

	colorIds := []uint{gotP1.Themes[0].ID + 100}
	err = repo.DeleteThemesWithId(gotP1.ID, colorIds)
	require.Equal(t, derror.ThemeNotFound, err)

	colorIds = []uint{gotP1.Themes[0].ID + 100, gotP1.Themes[0].ID}
	err = repo.DeleteThemesWithId(gotP1.ID, colorIds)
	require.Nil(t, err)
	gotThemes, err := repo.GetThemesWithProductId(gotP1.ID)
	require.Nil(t, err)
	require.Equal(t, len(p1.Colors)-1, len(gotThemes))

	t.Run("don't support roll back", func(t *testing.T) {
		p1 := model.Product{
			CompanyName: "Negin",
			CompanyId:   companyId,
			DesignCode:  "106",
			Colors:      []string{"قرمز", "آبی"},
			Dimensions:  []string{"6", "9"},
			Description: "توضیحات برای کد ۱۰6",
		}
		gotP1, err := repo.CreateProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, gotP1)

		colorIds = []uint{gotP1.Themes[0].ID, gotP1.Themes[0].ID + 100}
		err = repo.DeleteThemesWithId(gotP1.ID, colorIds)
		require.Nil(t, err)
		gotThemes, err = repo.GetThemesWithProductId(gotP1.ID)
		require.Nil(t, err)
		require.Equal(t, len(p1.Colors)-1, len(gotThemes))
	})
}

func TestThemeRepo_DeleteThemesWithId_Empty(t *testing.T) {
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

	err = repo.DeleteThemesWithId(gotP1.ID, []uint{})
	require.NotNil(t, derror.InvalidColor, err)
	gotThemes, err := repo.GetThemesWithProductId(gotP1.ID)
	require.Nil(t, err)
	require.Equal(t, len(p1.Colors), len(gotThemes))

}

func TestThemeRepo_DeleteThemesWithId_Nil(t *testing.T) {
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

	err = repo.DeleteThemesWithId(gotP1.ID, nil)
	require.NotNil(t, derror.InvalidColor, err)
}

func TestThemeRepo_UpdateThemesWithId_Ok(t *testing.T) {
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

	themes := []schema.Theme{
		{
			Model:     gorm.Model{ID: gotP1.Themes[0].ID}, //قرمز
			ProductId: gotP1.ID,
			Color:     "نارنجی",
		},
		{
			Model:     gorm.Model{ID: gotP1.Themes[2].ID}, //سبز
			ProductId: gotP1.ID,
			Color:     "سبز",
		},
	}
	err = repo.UpdateThemesWithId(gotP1.ID, themes)
	require.Nil(t, err)
	gotThemes, err := repo.GetThemesWithProductId(gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
	require.Equal(t, len(gotP1.Themes), len(gotThemes))
	require.Equal(t, themes[0].Color, gotThemes[0].Color)
}

func TestThemeRepo_UpdateThemesWithId_ProductNotExist(t *testing.T) {
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

	themes := []schema.Theme{
		{
			Model:     gorm.Model{ID: gotP1.Themes[0].ID}, //قرمز
			ProductId: 100,
			Color:     "نارنجی",
		},
	}
	err = repo.UpdateThemesWithId(100, themes)
	require.NotNil(t, err)
	require.Equal(t, derror.ThemeNotFound, err)
}

func TestThemeRepo_UpdateThemesWithId_ColorNotExist(t *testing.T) {
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
		err = repo.UpdateThemesWithId(gotP1.ID, themes)
		require.Equal(t, derror.ThemeNotFound, err)
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
		err = repo.UpdateThemesWithId(gotP1.ID, themes)
		require.Equal(t, derror.ThemeNotFound, err)
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
		err = repo.UpdateThemesWithId(gotP1.ID, themes)
		require.Equal(t, derror.ThemeNotFound, err)
		gotP1, err = repo.GetProductWithId(gotP1.ID)
		require.Equal(t, p1.Colors[0], gotP1.Themes[0].Color)
	})
}

func TestThemeRepo_UpdateThemesWithId_MultiProduct(t *testing.T) {
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
	err = repo.UpdateThemesWithId(gotP1.ID, themes)
	require.NotNil(t, derror.ThemeNotFound, err)
	gotTheme1, err := repo.GetThemesWithProductId(gotP1.ID)
	require.Equal(t, gotP1.Themes[0].Color, gotTheme1[0].Color)
	gotTheme2, err := repo.GetThemesWithProductId(gotP2.ID)
	require.Equal(t, gotP1.Themes[1].Color, gotTheme2[1].Color)
}

func TestThemeRepo_UpdateThemesWithId_Empty(t *testing.T) {
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

	err = repo.UpdateThemesWithId(gotP1.ID, []schema.Theme{})
	require.NotNil(t, err)
	require.Equal(t, derror.InvalidTheme, err)
}

func TestThemeRepo_UpdateThemesWithId_Nil(t *testing.T) {
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

	err = repo.UpdateThemesWithId(gotP1.ID, nil)
	require.NotNil(t, err)
	require.Equal(t, derror.InvalidTheme, err)
}

func TestThemeRepo_GetThemesWithProductId_Ok(t *testing.T) {
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

	gotThemes, err := repo.GetThemesWithProductId(gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
	require.Equal(t, len(gotP1.Themes), len(gotThemes))
	require.Equal(t, gotP1.Themes[0].Color, gotThemes[0].Color)
}

func TestThemeRepo_GetThemesWithProductId_ProductNotExist(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	gotThemes, err := repo.GetThemesWithProductId(100)
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
	require.Equal(t, 0, len(gotThemes))
}

func TestThemeRepo_GetThemesWithProductId_EmptyTheme(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	companyId := uint(1)
	p1 := model.Product{
		CompanyName: "Negin",
		CompanyId:   companyId,
		DesignCode:  "105",
		Colors:      []string{},
		Dimensions:  []string{"6", "9"},
		Description: "توضیحات برای کد ۱۰۵",
	}
	gotP1, err := repo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, gotP1)

	gotThemes, err := repo.GetThemesWithProductId(gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
	require.Equal(t, 0, len(gotThemes))
}

func TestThemeRepo_GetThemesWithProductId_NilTheme(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	companyId := uint(1)
	p1 := model.Product{
		CompanyName: "Negin",
		CompanyId:   companyId,
		DesignCode:  "105",
		Colors:      nil,
		Dimensions:  []string{"6", "9"},
		Description: "توضیحات برای کد ۱۰۵",
	}
	gotP1, err := repo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, gotP1)

	gotThemes, err := repo.GetThemesWithProductId(gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
	require.Equal(t, 0, len(gotThemes))
}
