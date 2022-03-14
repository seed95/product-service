package product

import (
	"github.com/seed95/product-service/internal/derror"
	"github.com/seed95/product-service/internal/model"
	"github.com/seed95/product-service/internal/repo/product/schema"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"testing"
)

func TestDimensionRepo_AddDimensionsToProduct_Ok(t *testing.T) {
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

	dimensions := []string{"12", "15"}
	gotDimensions, err := repo.AddDimensionsToProduct(p.ID, dimensions)
	require.Nil(t, err)
	require.NotNil(t, gotDimensions)
}

func TestDimensionRepo_AddDimensionsToProduct_ProductNotExist(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	dimensions := []string{"12"}
	gotDimensions, err := repo.AddDimensionsToProduct(34, dimensions)
	require.Nil(t, gotDimensions)
	require.NotNil(t, err)
}

func TestDimensionRepo_AddDimensionsToProduct_Duplicate(t *testing.T) {
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

	dimensions := []string{p1.Dimensions[0]}
	gotDimensions, err := repo.AddDimensionsToProduct(gotP1.ID, dimensions)
	require.Nil(t, gotDimensions)
	require.NotNil(t, err)

	t.Run("roll back", func(t *testing.T) {
		dimensions = []string{"12", p1.Dimensions[0]}
		gotDimensions, err = repo.AddDimensionsToProduct(gotP1.ID, dimensions)
		require.Nil(t, gotDimensions)
		require.NotNil(t, err)
		gotP1, err = repo.GetProductWithId(gotP1.ID)
		require.Equal(t, p1.Dimensions[0], gotP1.Dimensions[0].Size)
	})

}

func TestDimensionRepo_AddDimensionsToProduct_Empty(t *testing.T) {
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

	gotDimensions, err := repo.AddDimensionsToProduct(gotP1.ID, []string{})
	require.Nil(t, gotDimensions)
	require.Equal(t, derror.InvalidDimension, err)
}

func TestDimensionRepo_AddDimensionsToProduct_Nil(t *testing.T) {
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

	gotDimensions, err := repo.AddDimensionsToProduct(gotP1.ID, nil)
	require.Nil(t, gotDimensions)
	require.Equal(t, derror.InvalidDimension, err)
}

func TestDimensionRepo_DeleteDimensionsInProduct_Ok(t *testing.T) {
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
		Dimensions:  []string{"6", "9", "12"},
		Description: "توضیحات برای کد ۱۰۵",
	}
	gotP1, err := repo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, gotP1)

	dimensions := []string{"6", "9"}
	err = repo.DeleteDimensionsInProduct(gotP1.ID, dimensions)
	require.Nil(t, err)
	gotP1, err = repo.GetProductWithId(gotP1.ID)
	require.Equal(t, 1, len(gotP1.Dimensions))
	require.Equal(t, p1.Dimensions[2], gotP1.Dimensions[0].Size)
}

func TestDimensionRepo_DeleteDimensionsInProduct_ProductNotExist(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	dimensions := []string{"12"}
	err = repo.DeleteDimensionsInProduct(34, dimensions)
	require.Equal(t, err, derror.DimensionNotFound)
}

func TestDimensionRepo_DeleteDimensionsInProduct_DimensionNotExist(t *testing.T) {
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

	dimensions := []string{"12"}
	err = repo.DeleteDimensionsInProduct(gotP1.ID, dimensions)
	require.Equal(t, err, derror.DimensionNotFound)

	dimensions = []string{"12", "6"}
	err = repo.DeleteDimensionsInProduct(gotP1.ID, dimensions)
	require.Equal(t, err, derror.DimensionNotFound)
	gotP1, err = repo.GetProductWithId(gotP1.ID)
	require.Equal(t, len(p1.Dimensions), len(gotP1.Dimensions))

	t.Run("roll back", func(t *testing.T) {
		dimensions = []string{"6", "12"}
		err = repo.DeleteDimensionsInProduct(gotP1.ID, dimensions)
		require.Equal(t, err, derror.DimensionNotFound)
		gotP1, err = repo.GetProductWithId(gotP1.ID)
		require.Equal(t, len(p1.Dimensions), len(gotP1.Dimensions))
	})
}

func TestDimensionRepo_DeleteDimensionsInProduct_Empty(t *testing.T) {
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

	err = repo.DeleteDimensionsInProduct(gotP1.ID, []string{})
	require.Nil(t, err)
}

func TestDimensionRepo_DeleteDimensionsInProduct_Nil(t *testing.T) {
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

	err = repo.DeleteDimensionsInProduct(gotP1.ID, nil)
	require.Nil(t, err)
}

func TestDimensionRepo_UpdateDimensionsWithId_Ok(t *testing.T) {
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

	dimensions := []schema.Dimension{
		{
			Model:     gorm.Model{ID: gotP1.Dimensions[0].ID}, //6
			ProductId: gotP1.ID,
			Size:      "12",
		},
		{
			Model:     gorm.Model{ID: gotP1.Dimensions[1].ID}, //9
			ProductId: gotP1.ID,
			Size:      "15",
		},
	}
	err = repo.UpdateDimensionsWithId(dimensions)
	require.Nil(t, err)
	gotP1, err = repo.GetProductWithId(gotP1.ID)
	require.Equal(t, len(p1.Dimensions), len(gotP1.Dimensions))
	require.Equal(t, dimensions[0].Size, gotP1.Dimensions[0].Size)
	require.Equal(t, dimensions[1].Size, gotP1.Dimensions[1].Size)
}

func TestDimensionRepo_UpdateDimensionsWithId_ProductNotExist(t *testing.T) {
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

	dimensions := []schema.Dimension{
		{
			Model:     gorm.Model{ID: p.Dimensions[0].ID}, //6
			ProductId: 100,
			Size:      "12",
		},
	}
	err = repo.UpdateDimensionsWithId(dimensions)
	require.Equal(t, derror.DimensionNotFound, err)
}

func TestDimensionRepo_UpdateDimensionsWithId_DimensionNotExist(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("one dimension", func(t *testing.T) {
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

		dimensions := []schema.Dimension{
			{
				Model:     gorm.Model{ID: 100},
				ProductId: gotP1.ID,
				Size:      "12",
			},
		}
		err = repo.UpdateDimensionsWithId(dimensions)
		require.Equal(t, derror.DimensionNotFound, err)
	})

	t.Run("two dimension", func(t *testing.T) {
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

		dimensions := []schema.Dimension{
			{
				Model:     gorm.Model{ID: 100},
				ProductId: gotP1.ID,
				Size:      "5",
			},
			{
				Model:     gorm.Model{ID: gotP1.Dimensions[0].ID}, //6
				ProductId: gotP1.ID,
				Size:      "15",
			},
		}
		err = repo.UpdateDimensionsWithId(dimensions)
		require.Equal(t, derror.DimensionNotFound, err)
		gotP1, err = repo.GetProductWithId(gotP1.ID)
		require.Equal(t, p1.Dimensions[0], gotP1.Dimensions[0].Size)
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

		dimensions := []schema.Dimension{
			{
				Model:     gorm.Model{ID: gotP1.Dimensions[0].ID}, //6
				ProductId: gotP1.ID,
				Size:      "12",
			},
			{
				Model:     gorm.Model{ID: 100},
				ProductId: gotP1.ID,
				Size:      "15",
			},
		}
		err = repo.UpdateDimensionsWithId(dimensions)
		require.Equal(t, derror.DimensionNotFound, err)
		gotP1, err = repo.GetProductWithId(gotP1.ID)
		require.Equal(t, p1.Dimensions[0], gotP1.Dimensions[0].Size)
	})
}

func TestDimensionRepo_UpdateDimensionsWithId_MultiProduct(t *testing.T) {
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

	dimensions := []schema.Dimension{
		{
			Model:     gorm.Model{ID: gotP1.Dimensions[0].ID}, //6
			ProductId: gotP1.ID,
			Size:      "نارنجی",
		},
		{
			Model:     gorm.Model{ID: gotP2.Dimensions[1].ID}, //9
			ProductId: gotP2.ID,
			Size:      "سبز",
		},
	}
	err = repo.UpdateDimensionsWithId(dimensions)
	require.Nil(t, err)
	gotP1, err = repo.GetProductWithId(gotP1.ID)
	require.Equal(t, dimensions[0].Size, gotP1.Dimensions[0].Size)
	gotP2, err = repo.GetProductWithId(gotP2.ID)
	require.Equal(t, dimensions[1].Size, gotP2.Dimensions[1].Size)
}

func TestDimensionRepo_UpdateDimensionsWithId_Empty(t *testing.T) {
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

	err = repo.UpdateDimensionsWithId([]schema.Dimension{})
	require.Nil(t, err)
}

func TestDimensionRepo_UpdateDimensionsWithId_Nil(t *testing.T) {
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

	err = repo.UpdateDimensionsWithId(nil)
	require.Nil(t, err)
}

func TestThemeRepo_GetDimensionsWithProductId_Ok(t *testing.T) {
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

	gotDimensions, err := repo.GetDimensionsWithProductId(gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotDimensions)
	require.Equal(t, len(gotP1.Dimensions), len(gotDimensions))
	require.Equal(t, gotP1.Dimensions[0].Size, gotDimensions[0].Size)
}

func TestThemeRepo_GetDimensionsWithProductId_ProductNotExist(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	gotDimensions, err := repo.GetDimensionsWithProductId(100)
	require.Nil(t, err)
	require.NotNil(t, gotDimensions)
	require.Equal(t, 0, len(gotDimensions))
}

func TestThemeRepo_GetDimensionsWithProductId_EmptyDimension(t *testing.T) {
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
		Dimensions:  []string{},
		Description: "توضیحات برای کد ۱۰۵",
	}
	gotP1, err := repo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, gotP1)

	gotDimensions, err := repo.GetDimensionsWithProductId(gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotDimensions)
	require.Equal(t, 0, len(gotDimensions))
}

func TestThemeRepo_GetDimensionsWithProductId_NilDimension(t *testing.T) {
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
		Dimensions:  nil,
		Description: "توضیحات برای کد ۱۰۵",
	}
	gotP1, err := repo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, gotP1)

	gotDimensions, err := repo.GetDimensionsWithProductId(gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotDimensions)
	require.Equal(t, 0, len(gotDimensions))
}
