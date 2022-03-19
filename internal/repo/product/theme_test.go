package product

import (
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

func TestThemeRepo_GetThemesWithProductId_Ok(t *testing.T) {
	// NewProduct repo
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
	// NewProduct repo
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
	// NewProduct repo
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
		Sizes:       []string{"6", "9"},
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
	// NewProduct repo
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
		Sizes:       []string{"6", "9"},
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

func TestThemeRepo_InsertThemesWithColor_Ok(t *testing.T) {
	// NewProduct repo
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
		gotThemes, err = tRepo.InsertThemesWithColor(tx, gotP1.ID, colors)
		return err
	})
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
	require.Equal(t, len(colors), len(gotThemes))

	gotThemes, err = tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
	require.Equal(t, len(colors)+len(gotP1.Themes), len(gotThemes))
	require.Equal(t, gotP1.Themes[0].Color, gotThemes[0].Color)
}

func TestThemeRepo_InsertThemesWithColor_ProductNotExist(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Theme repo
	tRepo := NewThemeRepoMock()

	var gotThemes []schema.Theme
	colors := []string{"سبز"}
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		gotThemes, err = tRepo.InsertThemesWithColor(tx, 34, colors)
		return err
	})
	require.Nil(t, gotThemes)
	require.NotNil(t, err)
}

func TestThemeRepo_InsertThemesWithColor_Duplicate(t *testing.T) {
	// NewProduct repo
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
		gotThemes, err = tRepo.InsertThemesWithColor(tx, gotP1.ID, colors)
		return err
	})
	require.Nil(t, gotThemes)
	require.NotNil(t, err)

	gotThemes, err = tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
	require.Equal(t, len(gotP1.Dimensions), len(gotThemes))
	require.Equal(t, gotP1.Themes[0].Color, gotThemes[0].Color)

	t.Run("roll back", func(t *testing.T) {
		colors = []string{"سبز", gotP1.Themes[0].Color}
		err = pRepo.db.Transaction(func(tx *gorm.DB) error {
			gotThemes, err = tRepo.InsertThemesWithColor(tx, gotP1.ID, colors)
			return err
		})
		require.Nil(t, gotThemes)
		require.NotNil(t, err)

		gotThemes, err = tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
		require.Nil(t, err)
		require.NotNil(t, gotThemes)
		require.Equal(t, len(gotP1.Themes), len(gotThemes))
		require.Equal(t, gotP1.Themes[0].Color, gotThemes[0].Color)
	})
}

func TestThemeRepo_InsertThemesWithColor_Empty(t *testing.T) {
	// NewProduct repo
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
		gotThemes, err = tRepo.InsertThemesWithColor(tx, gotP1.ID, []string{})
		return err
	})
	require.Nil(t, gotThemes)
	require.Equal(t, derror.InvalidColor, err)

	gotThemes, err = tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
	require.Equal(t, len(gotP1.Themes), len(gotThemes))
}

func TestThemeRepo_InsertThemesWithColor_Nil(t *testing.T) {
	// NewProduct repo
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
		gotThemes, err = tRepo.InsertThemesWithColor(tx, gotP1.ID, nil)
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
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Theme repo
	tRepo := NewThemeRepoMock()

	themes := []schema.Theme{
		{Model: gorm.Model{ID: gotP1.Themes[0].ID}},
		{Model: gorm.Model{ID: gotP1.Themes[1].ID}},
	}
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		return tRepo.DeleteThemesWithId(tx, gotP1.ID, themes)
	})
	require.Nil(t, err)

	gotThemes, err := tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
	require.Equal(t, len(gotP1.Themes)-len(themes), len(gotThemes))
}

func TestThemeRepo_DeleteThemesWithId_ProductNotExist(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Theme repo
	tRepo := NewThemeRepoMock()

	themes := []schema.Theme{
		{Model: gorm.Model{ID: 100}},
	}
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		return tRepo.DeleteThemesWithId(tx, 34, themes)
	})
	require.Equal(t, derror.ThemeNotFound, err)
}

func TestThemeRepo_DeleteThemesWithId_ThemeNotExist(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Theme repo
	tRepo := NewThemeRepoMock()

	themes := []schema.Theme{
		{Model: gorm.Model{ID: gotP1.Themes[0].ID + 100}},
		{Model: gorm.Model{ID: gotP1.Themes[1].ID}},
	}
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		return tRepo.DeleteThemesWithId(tx, gotP1.ID, themes)
	})
	require.Equal(t, derror.ThemeNotFound, err)

	gotThemes, err := tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
	require.Equal(t, len(gotP1.Themes), len(gotThemes))

	t.Run("roll back", func(t *testing.T) {
		themes := []schema.Theme{
			{Model: gorm.Model{ID: gotP1.Themes[0].ID}},
			{Model: gorm.Model{ID: gotP1.Themes[1].ID + 100}},
		}
		err = pRepo.db.Transaction(func(tx *gorm.DB) error {
			return tRepo.DeleteThemesWithId(tx, gotP1.ID, themes)
		})
		require.Equal(t, derror.ThemeNotFound, err)

		gotThemes, err := tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
		require.Nil(t, err)
		require.NotNil(t, gotThemes)
		require.Equal(t, len(gotP1.Themes), len(gotThemes))
	})
}

func TestThemeRepo_DeleteThemesWithId_Empty(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Theme repo
	tRepo := NewThemeRepoMock()

	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		return tRepo.DeleteThemesWithId(tx, gotP1.ID, []schema.Theme{})
	})
	require.Equal(t, derror.InvalidTheme, err)

	gotThemes, err := tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
	require.Equal(t, len(gotP1.Themes), len(gotThemes))
}

func TestThemeRepo_DeleteThemesWithId_Nil(t *testing.T) {
	// NewProduct repo
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
	require.NotNil(t, gotThemes)
	require.Equal(t, len(gotP1.Themes), len(gotThemes))
}

func TestThemeRepo_EditThemes_Ok(t *testing.T) {
	// NewProduct repo
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
			Color: "نارنجی",
		},
		{
			Color: "سبز",
		},
	}
	var gotThemes []schema.Theme
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		gotThemes, err = tRepo.EditThemes(tx, gotP1.ID, themes)
		return err
	})
	require.Nil(t, err)
	require.Equal(t, len(gotP1.Themes), len(gotThemes))
}

func TestThemeRepo_EditThemes_ProductNotExist(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Theme repo
	tRepo := NewThemeRepoMock()

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	themes := []schema.Theme{{Color: "آبی"}, {Color: "آبی"}, {Color: "نارنجی"}}
	var gotThemes []schema.Theme
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		gotThemes, err = tRepo.EditThemes(tx, gotP1.ID+100, themes)
		return err
	})
	require.NotNil(t, err)
	require.Nil(t, gotThemes)

	gotThemes, err = tRepo.GetThemesWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotThemes)
	require.Equal(t, len(gotP1.Themes), len(gotThemes))
	require.Equal(t, gotP1.Themes[0].Color, gotThemes[0].Color)
}

func TestDimensionRepo_EditThemes_NotChange(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Theme repo
	tRepo := NewThemeRepoMock()

	themes := gotP1.Themes
	var gotThemes []schema.Theme
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		gotThemes, err = tRepo.EditThemes(tx, gotP1.ID, themes)
		return err
	})
	require.Nil(t, err)
	require.Equal(t, len(themes), len(gotThemes))
}
