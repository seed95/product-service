package product

import (
	"fmt"
	"github.com/seed95/product-service/internal/derror"
	"github.com/seed95/product-service/internal/model"
	"github.com/seed95/product-service/internal/repo/product/schema"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"testing"
)

func NewThemeRepoMock() ThemeService {
	return NewThemeService()
}

func TestThemeRepo_AddThemesWithColor_Ok1(t *testing.T) {

	// Product repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	companyId := uint(1)
	p1 := model.Product{
		CompanyName: "Negin",
		CompanyId:   companyId,
		DesignCode:  "105",
		Colors:      []string{"قرمز", "آبی"},
		Dimensions:  []string{"6", "9"},
		Description: "توضیحات برای کد ۱۰۵",
	}
	gotP1, err := pRepo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, gotP1)

	// Theme repo
	tRepo := NewThemeRepoMock()

	//colors := []string{"سبز", "نارنجی", "قرمز"}
	//tx := pRepo.db.Session(&gorm.Session{SkipDefaultTransaction: false})
	//gotThemes, err := tRepo.AddThemesWithColor(tx, gotP1.ID, colors)
	//fmt.Println(gotThemes)
	//fmt.Println(err)
	//tx.Commit()

	//err = pRepo.db.Transaction(func(tx *gorm.DB) error {
	//	tx.SkipDefaultTransaction = true
	//	//tx.DryRun = true
	//	colors := []string{"سبز", "نارنجی", "قرمز"}
	//	gotThemes, err := tRepo.AddThemesWithColor(tx, gotP1.ID, colors)
	//	fmt.Println(gotThemes)
	//	fmt.Println(err)
	//	//require.Nil(t, err)
	//	//require.NotNil(t, gotThemes)
	//	//return derror.InvalidColor
	//	return nil
	//})

	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		tx.SkipDefaultTransaction = true
		//tx.DryRun = true
		colors := []uint{gotP1.Themes[0].ID, gotP1.Themes[1].ID + 100}
		err := tRepo.DeleteThemesWithId(tx, gotP1.ID, colors)
		fmt.Println(err)
		//require.Nil(t, err)
		//require.NotNil(t, gotThemes)
		//return derror.InvalidColor
		return nil
	})
	//require.Nil(t, err)
}

func TestThemeRepo_GetThemesWithProductId_Ok(t *testing.T) {
	// Product repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Theme repo
	tRepo := NewThemeRepoMock()

	gotThemes, err := tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
	require.Equal(t, len(gotP1.Themes), len(gotThemes))
	require.Equal(t, gotP1.Themes[0].Color, gotThemes[0].Color)
}

func TestThemeRepo_GetThemesWithProductId_ProductNotExist(t *testing.T) {
	// Product repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Theme repo
	tRepo := NewThemeRepoMock()

	gotThemes, err := tRepo.GetThemesWithProductId(pRepo.db, 100)
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
	require.Equal(t, 0, len(gotThemes))
}

func TestThemeRepo_GetThemesWithProductId_EmptyTheme(t *testing.T) {
	// Product repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	p1 := model.Product{
		CompanyName: "Negin",
		CompanyId:   1,
		DesignCode:  "105",
		Colors:      []string{},
		Dimensions:  []string{"6", "9"},
		Description: "توضیحات برای کد ۱۰۵",
	}
	gotP1, err := pRepo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, gotP1)

	// Theme repo
	tRepo := NewThemeRepoMock()

	gotThemes, err := tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
	require.Equal(t, 0, len(gotThemes))
}

func TestThemeRepo_GetThemesWithProductId_NilTheme(t *testing.T) {
	// Product repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	companyId := uint(1)
	p1 := model.Product{
		CompanyName: "Negin",
		CompanyId:   companyId,
		DesignCode:  "105",
		Colors:      nil,
		Dimensions:  []string{"6", "9"},
		Description: "توضیحات برای کد ۱۰۵",
	}
	gotP1, err := pRepo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, gotP1)

	// Theme repo
	tRepo := NewThemeRepoMock()

	gotThemes, err := tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
	require.Equal(t, 0, len(gotThemes))
}

func TestThemeRepo_AddThemesWithColor_Ok(t *testing.T) {
	// Product repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Theme repo
	tRepo := NewThemeRepoMock()

	var gotThemes []schema.Theme
	colors := []string{"سبز", "نارنجی", "green"}
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		gotThemes, err = tRepo.AddThemesWithColor(tx, gotP1.ID, colors)
		return err
	})
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
	require.Equal(t, len(colors), len(gotThemes))
	gotThemes, err = tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
	require.Equal(t, len(colors)+len(gotP1.Themes), len(gotThemes))
}

func TestThemeRepo_AddThemesWithColor_ProductNotExist(t *testing.T) {
	// Product repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Theme repo
	tRepo := NewThemeRepoMock()

	var gotThemes []schema.Theme
	colors := []string{"سبز"}
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		gotThemes, err = tRepo.AddThemesWithColor(tx, 34, colors)
		return err
	})
	require.Nil(t, gotThemes)
	require.NotNil(t, err)
}

func TestThemeRepo_AddThemesWithColor_Duplicate(t *testing.T) {
	// Product repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Theme repo
	tRepo := NewThemeRepoMock()

	var gotThemes []schema.Theme
	colors := []string{gotP1.Themes[0].Color, "سبز"}
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		gotThemes, err = tRepo.AddThemesWithColor(tx, gotP1.ID, colors)
		return err
	})
	require.Nil(t, gotThemes)
	require.NotNil(t, err)

	t.Run("roll back", func(t *testing.T) {
		colors = []string{"سبز", gotP1.Themes[0].Color}
		err = pRepo.db.Transaction(func(tx *gorm.DB) error {
			gotThemes, err = tRepo.AddThemesWithColor(tx, gotP1.ID, colors)
			return err
		})
		require.Nil(t, gotThemes)
		require.NotNil(t, err)

		gotThemes, err = tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
		require.Nil(t, err)
		require.NotNil(t, gotThemes)
		require.Equal(t, len(gotP1.Themes), len(gotThemes))
	})

}

func TestThemeRepo_AddThemesWithColor_Empty(t *testing.T) {
	// Product repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Theme repo
	tRepo := NewThemeRepoMock()

	var gotThemes []schema.Theme
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		gotThemes, err = tRepo.AddThemesWithColor(tx, gotP1.ID, []string{})
		return err
	})
	require.Nil(t, gotThemes)
	require.Equal(t, derror.InvalidColor, err)
	gotThemes, err = tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
	require.Equal(t, len(gotP1.Themes), len(gotThemes))
}

func TestThemeRepo_AddThemesWithColor_Nil(t *testing.T) {
	// Product repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Theme repo
	tRepo := NewThemeRepoMock()

	var gotThemes []schema.Theme
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		gotThemes, err = tRepo.AddThemesWithColor(tx, gotP1.ID, nil)
		return err
	})
	require.Nil(t, gotThemes)
	require.Equal(t, derror.InvalidColor, err)
	gotThemes, err = tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
	require.Equal(t, len(gotP1.Themes), len(gotThemes))
}

func TestThemeRepo_DeleteThemesWithId_Ok(t *testing.T) {
	// Product repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Theme repo
	tRepo := NewThemeRepoMock()

	themeIds := []uint{gotP1.Themes[0].ID, gotP1.Themes[1].ID}
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		return tRepo.DeleteThemesWithId(tx, gotP1.ID, themeIds)
	})
	require.Nil(t, err)
	gotThemes, err := tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
	require.Equal(t, len(gotP1.Themes)-len(themeIds), len(gotThemes))
}

func TestThemeRepo_DeleteThemesWithId_ProductNotExist(t *testing.T) {
	// Product repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Theme repo
	tRepo := NewThemeRepoMock()

	themeIds := []uint{100}
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		return tRepo.DeleteThemesWithId(tx, 34, themeIds)
	})
	require.Equal(t, derror.ThemeNotFound, err)
}

func TestThemeRepo_DeleteThemesWithId_ThemeNotExist(t *testing.T) {
	// Product repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Theme repo
	tRepo := NewThemeRepoMock()

	themeIds := []uint{gotP1.Themes[0].ID + 100, gotP1.Themes[0].ID}
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		return tRepo.DeleteThemesWithId(tx, gotP1.ID, themeIds)
	})
	require.Equal(t, derror.ThemeNotFound, err)
	gotThemes, err := tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
	require.Equal(t, len(gotP1.Themes), len(gotThemes))

	t.Run("roll back", func(t *testing.T) {
		themeIds = []uint{gotP1.Themes[0].ID, gotP1.Themes[0].ID + 100}
		err = pRepo.db.Transaction(func(tx *gorm.DB) error {
			return tRepo.DeleteThemesWithId(tx, gotP1.ID, themeIds)
		})
		require.Equal(t, derror.ThemeNotFound, err)
		gotThemes, err := tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
		require.Nil(t, err)
		require.NotNil(t, gotThemes)
		require.Equal(t, len(gotP1.Themes), len(gotThemes))
	})
}

func TestThemeRepo_DeleteThemesWithId_Empty(t *testing.T) {
	// Product repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Theme repo
	tRepo := NewThemeRepoMock()

	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		return tRepo.DeleteThemesWithId(tx, gotP1.ID, []uint{})
	})
	require.NotNil(t, derror.InvalidTheme, err)
	gotThemes, err := tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.Equal(t, len(gotP1.Themes), len(gotThemes))
}

func TestThemeRepo_DeleteThemesWithId_Nil(t *testing.T) {
	// Product repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Theme repo
	tRepo := NewThemeRepoMock()

	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		return tRepo.DeleteThemesWithId(tx, gotP1.ID, nil)
	})
	require.NotNil(t, derror.InvalidTheme, err)
	gotThemes, err := tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.Equal(t, len(gotP1.Themes), len(gotThemes))
}

func TestThemeRepo_UpdateThemesWithId_Ok(t *testing.T) {
	// Product repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Theme repo
	tRepo := NewThemeRepoMock()

	themes := []schema.Theme{
		{
			Model: gorm.Model{ID: gotP1.Themes[0].ID}, //قرمز
			Color: "نارنجی",
		},
		{
			Model: gorm.Model{ID: gotP1.Themes[1].ID}, //آبی
			Color: "سبز",
		},
	}
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		return tRepo.UpdateThemesWithId(tx, gotP1.ID, themes)
	})
	require.Nil(t, err)
	gotThemes, err := tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
	require.Equal(t, len(gotP1.Themes), len(gotThemes))
	require.Equal(t, themes[0].Color, gotThemes[0].Color)
}

func TestThemeRepo_UpdateThemesWithId_ProductNotExist(t *testing.T) {
	// Product repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Theme repo
	tRepo := NewThemeRepoMock()

	themes := []schema.Theme{
		{
			Model: gorm.Model{ID: gotP1.Themes[0].ID}, //قرمز
			Color: "نارنجی",
		},
	}
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		return tRepo.UpdateThemesWithId(tx, gotP1.ID+100, themes)
	})
	require.Equal(t, derror.ThemeNotFound, err)
	gotThemes, err := tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
	require.Equal(t, len(gotP1.Themes), len(gotThemes))
	require.NotEqual(t, themes[0].Color, gotThemes[0].Color)
}

func TestThemeRepo_UpdateThemesWithId_ThemeNotExist(t *testing.T) {
	// Product repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Theme repo
	tRepo := NewThemeRepoMock()

	t.Run("one color", func(t *testing.T) {
		// Create product
		gotP1 := CreateProduct1(pRepo, t)

		themes := []schema.Theme{
			{
				Model: gorm.Model{ID: gotP1.Themes[0].ID + 100},
				Color: "نارنجی",
			},
		}
		err = pRepo.db.Transaction(func(tx *gorm.DB) error {
			return tRepo.UpdateThemesWithId(tx, gotP1.ID, themes)
		})
		require.Equal(t, derror.ThemeNotFound, err)
		gotThemes, err := tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
		require.Nil(t, err)
		require.NotNil(t, gotThemes)
		require.Equal(t, len(gotP1.Themes), len(gotThemes))
		require.NotEqual(t, themes[0].Color, gotThemes[0].Color)
	})

	t.Run("two color", func(t *testing.T) {
		// Create product
		gotP2 := CreateProduct2(pRepo, t)

		themes := []schema.Theme{
			{
				Model:     gorm.Model{ID: 100},
				ProductId: gotP2.ID,
				Color:     "صورتی",
			},
			{
				Model:     gorm.Model{ID: gotP2.Themes[0].ID}, //قرمز
				ProductId: gotP2.ID,
				Color:     "نارنجی",
			},
		}
		err = pRepo.db.Transaction(func(tx *gorm.DB) error {
			return tRepo.UpdateThemesWithId(tx, gotP2.ID, themes)
		})
		require.Equal(t, derror.ThemeNotFound, err)
		gotThemes, err := tRepo.GetThemesWithProductId(pRepo.db, gotP2.ID)
		require.Nil(t, err)
		require.NotNil(t, gotThemes)
		require.Equal(t, len(gotP2.Themes), len(gotThemes))
		require.NotEqual(t, themes[0].Color, gotThemes[0].Color)
	})

	t.Run("roll back", func(t *testing.T) {
		// Create product
		gotP3 := CreateProduct3(pRepo, t)

		themes := []schema.Theme{
			{
				Model:     gorm.Model{ID: gotP3.Themes[0].ID}, //قرمز
				ProductId: gotP3.ID,
				Color:     "نارنجی",
			},
			{
				Model:     gorm.Model{ID: 100},
				ProductId: gotP3.ID,
				Color:     "صورتی",
			},
		}
		err = pRepo.db.Transaction(func(tx *gorm.DB) error {
			return tRepo.UpdateThemesWithId(tx, gotP3.ID, themes)
		})
		require.Equal(t, derror.ThemeNotFound, err)
		gotThemes, err := tRepo.GetThemesWithProductId(pRepo.db, gotP3.ID)
		require.Nil(t, err)
		require.NotNil(t, gotThemes)
		require.Equal(t, len(gotP3.Themes), len(gotThemes))
		require.NotEqual(t, themes[0].Color, gotThemes[0].Color)
	})
}

func TestThemeRepo_UpdateThemesWithId_MultiProduct(t *testing.T) {
	// Product repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)
	gotP2 := CreateProduct2(pRepo, t)

	// Theme repo
	tRepo := NewThemeRepoMock()

	themes := []schema.Theme{
		{
			Model: gorm.Model{ID: gotP1.Themes[0].ID}, //قرمز
			Color: "نارنجی",
		},
		{
			Model: gorm.Model{ID: gotP2.Themes[1].ID}, //آبی
			Color: "سبز",
		},
	}
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		return tRepo.UpdateThemesWithId(tx, gotP1.ID, themes)
	})
	require.NotNil(t, derror.ThemeNotFound, err)
	gotTheme1, err := tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.Equal(t, gotP1.Themes[0].Color, gotTheme1[0].Color)
	gotTheme2, err := tRepo.GetThemesWithProductId(pRepo.db, gotP2.ID)
	require.Nil(t, err)
	require.Equal(t, gotP1.Themes[1].Color, gotTheme2[1].Color)
}

func TestThemeRepo_UpdateThemesWithId_Empty(t *testing.T) {
	// Product repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Theme repo
	tRepo := NewThemeRepoMock()

	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		return tRepo.UpdateThemesWithId(tx, gotP1.ID, []schema.Theme{})
	})
	require.Equal(t, derror.InvalidTheme, err)
}

func TestThemeRepo_UpdateThemesWithId_Nil(t *testing.T) {
	// Product repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Theme repo
	tRepo := NewThemeRepoMock()

	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		return tRepo.UpdateThemesWithId(tx, gotP1.ID, nil)
	})
	require.Equal(t, derror.InvalidTheme, err)
}

func TestThemeRepo_EditThemesWithId_Ok(t *testing.T) {
	// Product repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Theme repo
	tRepo := NewThemeRepoMock()

	themes := []schema.Theme{
		{
			Model: gorm.Model{ID: gotP1.Themes[0].ID},
			Color: "نارنجی",
		},
		{
			Color: "سبز",
		},
	}
	var gotThemes []schema.Theme
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		gotThemes, err = tRepo.EditThemesWithId(tx, gotP1.ID, themes)
		return err
	})
	require.Nil(t, err)
	require.Equal(t, len(gotP1.Themes), len(gotThemes))
	require.Equal(t, themes[0].Color, gotThemes[0].Color)
}

func TestThemeRepo_EditThemesWithId_ProductNotExist(t *testing.T) {
	// Product repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Theme repo
	tRepo := NewThemeRepoMock()

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	themes := []schema.Theme{
		{
			Model: gorm.Model{ID: gotP1.Themes[0].ID},
			Color: "آبی",
		},
		{
			Model: gorm.Model{ID: gotP1.Themes[1].ID},
			Color: "آبی",
		},
		{
			Color: "نارنجی",
		},
	}
	var gotThemes []schema.Theme
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		gotThemes, err = tRepo.EditThemesWithId(tx, gotP1.ID+100, themes)
		return err
	})
	require.NotNil(t, err)
	require.Nil(t, gotThemes)
	gotThemes, err = tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
	require.Equal(t, len(gotP1.Themes), len(gotThemes))
	require.Equal(t, gotP1.Themes[0].Color, gotThemes[0].Color)
}

func TestThemeRepo_EditThemesWithId_ThemeNotExist(t *testing.T) {
	// Product repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Theme repo
	tRepo := NewThemeRepoMock()

	t.Run("first", func(t *testing.T) {
		// Create product
		gotP1 := CreateProduct1(pRepo, t)

		themes := []schema.Theme{
			{
				Model: gorm.Model{ID: gotP1.Themes[0].ID + 100},
				Color: "آبی",
			},
			{
				Model: gorm.Model{ID: gotP1.Themes[1].ID},
				Color: "آبی",
			},
			{
				Color: "نارنجی",
			},
		}
		var gotThemes []schema.Theme
		err = pRepo.db.Transaction(func(tx *gorm.DB) error {
			gotThemes, err = tRepo.EditThemesWithId(tx, gotP1.ID, themes)
			return err
		})
		require.NotNil(t, err)
		require.Nil(t, gotThemes)
		gotThemes, err = tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
		require.Equal(t, len(gotP1.Themes), len(gotThemes))
		require.Equal(t, gotP1.Themes[0].Color, gotThemes[0].Color)
	})

	t.Run("second", func(t *testing.T) {
		// Create product
		gotP2 := CreateProduct2(pRepo, t)

		themes := []schema.Theme{
			{
				Model: gorm.Model{ID: gotP2.Themes[0].ID},
				Color: "آبی",
			},
			{
				Model: gorm.Model{ID: gotP2.Themes[1].ID + 100},
				Color: "آبی",
			},
			{
				Color: "نارنجی",
			},
		}
		var gotThemes []schema.Theme
		err = pRepo.db.Transaction(func(tx *gorm.DB) error {
			gotThemes, err = tRepo.EditThemesWithId(tx, gotP2.ID, themes)
			return err
		})
		require.NotNil(t, err)
		require.Nil(t, gotThemes)
		gotThemes, err = tRepo.GetThemesWithProductId(pRepo.db, gotP2.ID)
		require.Equal(t, len(gotP2.Themes), len(gotThemes))
		require.Equal(t, gotP2.Themes[0].Color, gotThemes[0].Color)
	})

}

func TestThemeRepo_EditThemesWithId_Rollback(t *testing.T) {
	// Product repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Theme repo
	tRepo := NewThemeRepoMock()

	t.Run("update", func(t *testing.T) {
		// Create product
		gotP1 := CreateProduct1(pRepo, t)

		themes := []schema.Theme{
			{
				Model: gorm.Model{ID: gotP1.Themes[0].ID},
				Color: "آبی",
			},
			{
				Model: gorm.Model{ID: gotP1.Themes[1].ID},
				Color: "آبی",
			},
			{
				Color: "نارنجی",
			},
		}
		var gotThemes []schema.Theme
		err = pRepo.db.Transaction(func(tx *gorm.DB) error {
			gotThemes, err = tRepo.EditThemesWithId(tx, gotP1.ID, themes)
			return err
		})
		require.NotNil(t, err)
		require.Nil(t, gotThemes)
		gotThemes, err = tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
		require.Equal(t, len(gotP1.Themes), len(gotThemes))
		require.Equal(t, gotP1.Themes[0].Color, gotThemes[0].Color)
	})

	t.Run("add", func(t *testing.T) {
		// Create product
		gotP2 := CreateProduct2(pRepo, t)

		themes := []schema.Theme{
			{
				Model: gorm.Model{ID: gotP2.Themes[0].ID},
				Color: "نارنجی",
			},
			{
				Model: gorm.Model{ID: gotP2.Themes[1].ID},
				Color: "آبی",
			},
			{
				Color: "نارنجی",
			},
		}
		var gotThemes []schema.Theme
		err = pRepo.db.Transaction(func(tx *gorm.DB) error {
			gotThemes, err = tRepo.EditThemesWithId(tx, gotP2.ID, themes)
			return err
		})
		require.NotNil(t, err)
		require.Nil(t, gotThemes)
		gotThemes, err = tRepo.GetThemesWithProductId(pRepo.db, gotP2.ID)
		require.Equal(t, len(gotP2.Themes), len(gotThemes))
		require.Equal(t, gotP2.Themes[0].Color, gotThemes[0].Color)

		themes = []schema.Theme{
			{
				Model: gorm.Model{ID: gotP2.Themes[0].ID},
				Color: "نارنجی",
			},
			{
				Model: gorm.Model{ID: gotP2.Themes[1].ID},
				Color: "آبی",
			},
			{
				Color: "آبی",
			},
		}
		err = pRepo.db.Transaction(func(tx *gorm.DB) error {
			gotThemes, err = tRepo.EditThemesWithId(tx, gotP2.ID, themes)
			return err
		})
		require.NotNil(t, err)
		require.Nil(t, gotThemes)
		gotThemes, err = tRepo.GetThemesWithProductId(pRepo.db, gotP2.ID)
		require.Equal(t, len(gotP2.Themes), len(gotThemes))
		require.Equal(t, gotP2.Themes[0].Color, gotThemes[0].Color)
	})
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
