package product

import (
	"github.com/seed95/product-service/internal/derror"
	"github.com/seed95/product-service/internal/model"
	"github.com/seed95/product-service/internal/repo/product/schema"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"testing"
)

func NewDimensionRepoMock() DimensionService {
	return NewDimensionService()
}

func TestDimensionRepo_GetDimensionsWithProductId_Ok(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Dimension repo
	dRepo := NewDimensionRepoMock()

	gotDimensions, err := dRepo.GetDimensionsWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotDimensions)
	require.Equal(t, len(gotP1.Dimensions), len(gotDimensions))
	require.Equal(t, gotP1.Dimensions[0].Size, gotDimensions[0].Size)
}

func TestDimensionRepo_GetDimensionsWithProductId_ProductNotExist(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Dimension repo
	dRepo := NewDimensionRepoMock()

	gotDimensions, err := dRepo.GetDimensionsWithProductId(pRepo.db, 100)
	require.Nil(t, err)
	require.NotNil(t, gotDimensions)
	require.Equal(t, 0, len(gotDimensions))
}

func TestDimensionRepo_GetDimensionsWithProductId_EmptyDimension(t *testing.T) {
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
		Colors:      []string{"قرمز", "آبی"},
		Sizes:       []string{},
		Description: "توضیحات برای کد ۱۰۵",
	}
	gotP1, err := pRepo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, gotP1)

	// Dimension repo
	dRepo := NewDimensionRepoMock()

	gotDimensions, err := dRepo.GetDimensionsWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotDimensions)
	require.Equal(t, 0, len(gotDimensions))
}

func TestDimensionRepo_GetDimensionsWithProductId_NilDimension(t *testing.T) {
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
		Colors:      []string{"قرمز", "آبی"},
		Sizes:       nil,
		Description: "توضیحات برای کد ۱۰۵",
	}
	gotP1, err := pRepo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, gotP1)

	// Dimension repo
	dRepo := NewDimensionRepoMock()

	gotDimensions, err := dRepo.GetDimensionsWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotDimensions)
	require.Equal(t, 0, len(gotDimensions))
}

func TestDimensionRepo_InsertDimensions_Ok(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Dimension repo
	dRepo := NewDimensionRepoMock()

	var gotDimensions []schema.Dimension
	sizes := []string{"12", "15"}
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		gotDimensions, err = dRepo.InsertDimensions(tx, gotP1.ID, sizes)
		return err
	})
	require.Nil(t, err)
	require.NotNil(t, gotDimensions)
	require.Equal(t, len(sizes), len(gotDimensions))

	gotDimensions, err = dRepo.GetDimensionsWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotDimensions)
	require.Equal(t, len(sizes)+len(gotP1.Dimensions), len(gotDimensions))
	require.Equal(t, gotP1.Dimensions[0].Size, gotDimensions[0].Size)
}

func TestDimensionRepo_InsertDimensions_ProductNotExist(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Dimension repo
	dRepo := NewDimensionRepoMock()

	var gotDimensions []schema.Dimension
	sizes := []string{"12", "15"}
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		gotDimensions, err = dRepo.InsertDimensions(tx, 34, sizes)
		return err
	})
	require.Nil(t, gotDimensions)
	require.NotNil(t, err)
}

func TestDimensionRepo_InsertDimensions_Duplicate(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Dimension repo
	dRepo := NewDimensionRepoMock()

	var gotDimensions []schema.Dimension
	sizes := []string{gotP1.Dimensions[0].Size, "15"}
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		gotDimensions, err = dRepo.InsertDimensions(tx, gotP1.ID, sizes)
		return err
	})
	require.Nil(t, gotDimensions)
	require.NotNil(t, err)

	gotDimensions, err = dRepo.GetDimensionsWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotDimensions)
	require.Equal(t, len(gotP1.Dimensions), len(gotDimensions))
	require.Equal(t, gotP1.Dimensions[0].Size, gotDimensions[0].Size)

	t.Run("roll back", func(t *testing.T) {
		sizes = []string{"12", gotP1.Dimensions[0].Size}
		err = pRepo.db.Transaction(func(tx *gorm.DB) error {
			gotDimensions, err = dRepo.InsertDimensions(pRepo.db, gotP1.ID, sizes)
			return err
		})
		require.Nil(t, gotDimensions)
		require.NotNil(t, err)

		gotDimensions, err = dRepo.GetDimensionsWithProductId(pRepo.db, gotP1.ID)
		require.Nil(t, err)
		require.NotNil(t, gotDimensions)
		require.Equal(t, len(gotP1.Dimensions), len(gotDimensions))
		require.Equal(t, gotP1.Dimensions[0].Size, gotDimensions[0].Size)
	})
}

func TestDimensionRepo_InsertDimensions_Empty(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Dimension repo
	dRepo := NewDimensionRepoMock()

	var gotDimensions []schema.Dimension
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		gotDimensions, err = dRepo.InsertDimensions(tx, gotP1.ID, []string{})
		return err
	})
	require.Nil(t, gotDimensions)
	require.Equal(t, derror.InvalidDimension, err)

	gotDimensions, err = dRepo.GetDimensionsWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotDimensions)
	require.Equal(t, len(gotP1.Dimensions), len(gotDimensions))
}

func TestDimensionRepo_InsertDimensions_Nil(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Dimension repo
	dRepo := NewDimensionRepoMock()

	var gotDimensions []schema.Dimension
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		gotDimensions, err = dRepo.InsertDimensions(tx, gotP1.ID, nil)
		return err
	})
	require.Nil(t, gotDimensions)
	require.Equal(t, derror.InvalidDimension, err)

	gotDimensions, err = dRepo.GetDimensionsWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotDimensions)
	require.Equal(t, len(gotP1.Dimensions), len(gotDimensions))
}

func TestDimensionRepo_DeleteDimensionsWithId_Ok(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Dimension repo
	dRepo := NewDimensionRepoMock()

	dimensions := []schema.Dimension{
		{Model: gorm.Model{ID: gotP1.Dimensions[0].ID}},
		{Model: gorm.Model{ID: gotP1.Dimensions[1].ID}},
	}
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		return dRepo.DeleteDimensionsWithId(tx, gotP1.ID, dimensions)
	})
	require.Nil(t, err)

	gotDimensions, err := dRepo.GetDimensionsWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotDimensions)
	require.Equal(t, len(gotP1.Dimensions)-len(dimensions), len(gotDimensions))
}

func TestDimensionRepo_DeleteDimensionsWithId_ProductNotExist(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Dimension repo
	dRepo := NewDimensionRepoMock()

	dimensions := []schema.Dimension{{Model: gorm.Model{ID: 100}}}
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		return dRepo.DeleteDimensionsWithId(tx, 34, dimensions)
	})
	require.Equal(t, derror.DimensionNotFound, err)
}

func TestDimensionRepo_DeleteDimensionsWithId_DimensionNotExist(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Dimension repo
	dRepo := NewDimensionRepoMock()

	dimensions := []schema.Dimension{
		{Model: gorm.Model{ID: gotP1.Dimensions[0].ID + 100}},
		{Model: gorm.Model{ID: gotP1.Dimensions[1].ID}},
	}
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		return dRepo.DeleteDimensionsWithId(tx, gotP1.ID, dimensions)
	})
	require.Equal(t, derror.DimensionNotFound, err)

	gotDimensions, err := dRepo.GetDimensionsWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotDimensions)
	require.Equal(t, len(gotP1.Dimensions), len(gotDimensions))

	t.Run("roll back", func(t *testing.T) {
		dimensions = []schema.Dimension{
			{Model: gorm.Model{ID: gotP1.Dimensions[0].ID}},
			{Model: gorm.Model{ID: gotP1.Dimensions[1].ID + 100}},
		}
		err = pRepo.db.Transaction(func(tx *gorm.DB) error {
			return dRepo.DeleteDimensionsWithId(tx, gotP1.ID, dimensions)
		})
		require.Equal(t, derror.DimensionNotFound, err)

		gotDimensions, err := dRepo.GetDimensionsWithProductId(pRepo.db, gotP1.ID)
		require.Nil(t, err)
		require.NotNil(t, gotDimensions)
		require.Equal(t, len(gotP1.Dimensions), len(gotDimensions))
	})
}

func TestDimensionRepo_DeleteDimensionsWithId_Empty(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Dimension repo
	dRepo := NewDimensionRepoMock()

	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		return dRepo.DeleteDimensionsWithId(tx, gotP1.ID, []schema.Dimension{})
	})
	require.Equal(t, derror.InvalidDimension, err)

	gotDimensions, err := dRepo.GetDimensionsWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotDimensions)
	require.Equal(t, len(gotP1.Dimensions), len(gotDimensions))
}

func TestDimensionRepo_DeleteDimensionsWithId_Nil(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Dimension repo
	dRepo := NewDimensionRepoMock()

	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		return dRepo.DeleteDimensionsWithId(tx, gotP1.ID, nil)
	})
	require.Equal(t, derror.InvalidDimension, err)

	gotDimensions, err := dRepo.GetDimensionsWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotDimensions)
	require.Equal(t, len(gotP1.Dimensions), len(gotDimensions))
}

func TestDimensionRepo_EditDimensions_Ok(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Dimension repo
	dRepo := NewDimensionRepoMock()

	dimensions := []schema.Dimension{{Size: "12"}, {Size: "6"}, {Size: "15"}}
	var gotDimensions []schema.Dimension
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		gotDimensions, err = dRepo.EditDimensions(tx, gotP1.ID, dimensions)
		return err
	})
	require.Nil(t, err)
	require.Equal(t, len(dimensions), len(gotDimensions))
}

func TestDimensionRepo_EditDimensions_ProductNotExist(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Dimension repo
	dRepo := NewDimensionRepoMock()

	dimensions := []schema.Dimension{{Size: "12"}, {Size: "6"}, {Size: "15"}}
	var gotDimensions []schema.Dimension
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		gotDimensions, err = dRepo.EditDimensions(tx, gotP1.ID+100, dimensions)
		return err
	})
	require.NotNil(t, err)
	require.Nil(t, gotDimensions)

	gotDimensions, err = dRepo.GetDimensionsWithProductId(pRepo.db, gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotDimensions)
	require.Equal(t, len(gotP1.Dimensions), len(gotDimensions))
	require.Equal(t, gotP1.Dimensions[0].Size, gotDimensions[0].Size)
}

func TestDimensionRepo_EditDimensions_NotChange(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	// Dimension repo
	dRepo := NewDimensionRepoMock()

	dimensions := gotP1.Dimensions
	var gotDimensions []schema.Dimension
	err = pRepo.db.Transaction(func(tx *gorm.DB) error {
		gotDimensions, err = dRepo.EditDimensions(tx, gotP1.ID, dimensions)
		return err
	})
	require.Nil(t, err)
	require.Equal(t, len(dimensions), len(gotDimensions))
}
