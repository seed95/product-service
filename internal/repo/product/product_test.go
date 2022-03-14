package product

import (
	"github.com/seed95/product-service/internal/derror"
	"github.com/seed95/product-service/internal/model"
	"github.com/seed95/product-service/internal/repo/product/schema"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProductRepo_CreateProduct_ZeroId(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	p1 := model.Product{
		Id:          0,
		CompanyName: "Negin",
		CompanyId:   1,
		DesignCode:  "102",
		Colors:      []string{"آبی", "قرمز"},
		Dimensions:  []string{"6", "9"},
		Description: "توضیحات برای کد ۱۰۲",
	}

	p, err := repo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, p)
}

func TestProductRepo_CreateProduct_Duplicate(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Duplicate design code
	t.Run("design code", func(t *testing.T) {
		p1 := model.Product{
			CompanyName: "Negin",
			CompanyId:   1,
			DesignCode:  "102",
			Colors:      []string{"آبی", "قرمز"},
			Dimensions:  []string{"6", "9"},
			Description: "توضیحات برای کد ۱۰۲",
		}

		p, err := repo.CreateProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, p)

		p, err = repo.CreateProduct(p1)
		require.NotNil(t, err)
		require.Nil(t, p)
	})

	// Duplicate size
	t.Run("size", func(t *testing.T) {

		p1 := model.Product{
			CompanyName: "Negin",
			CompanyId:   1,
			DesignCode:  "103",
			Colors:      []string{"آبی", "قرمز"},
			Dimensions:  []string{"6", "9", "6"},
			Description: "توضیحات برای کد ۱۰۳",
		}

		p, err := repo.CreateProduct(p1)
		require.NotNil(t, err)
		require.Nil(t, p)
	})

	// Duplicate color
	t.Run("color", func(t *testing.T) {

		p1 := model.Product{
			CompanyName: "Negin",
			CompanyId:   1,
			DesignCode:  "105",
			Colors:      []string{"آبی", "آبی"},
			Dimensions:  []string{"6", "9"},
			Description: "توضیحات برای کد ۱۰۵",
		}

		p, err := repo.CreateProduct(p1)
		require.NotNil(t, err)
		require.Nil(t, p)
	})
}

func TestProductRepo_CreateProduct_Empty(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Empty product
	t.Run("product", func(t *testing.T) {
		p, err := repo.CreateProduct(model.Product{})
		require.Nil(t, err)
		require.NotNil(t, p)
	})

	// Empty company
	t.Run("company", func(t *testing.T) {

		p1 := model.Product{
			CompanyName: "",
			CompanyId:   0,
			DesignCode:  "103",
			Colors:      []string{"آبی", "قرمز"},
			Dimensions:  []string{},
			Description: "توضیحات برای کد ۱۰۳",
		}

		p, err := repo.CreateProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, p)
	})

	// Empty design code
	t.Run("design code", func(t *testing.T) {

		p1 := model.Product{
			CompanyName: "Negin",
			CompanyId:   1,
			DesignCode:  "",
			Colors:      []string{"آبی", "قرمز"},
			Dimensions:  []string{},
			Description: "توضیحات برای کد خالی",
		}

		p, err := repo.CreateProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, p)
	})

	// Empty size
	t.Run("size", func(t *testing.T) {

		p1 := model.Product{
			CompanyName: "Negin",
			CompanyId:   1,
			DesignCode:  "103",
			Colors:      []string{"آبی", "قرمز"},
			Dimensions:  []string{},
			Description: "توضیحات برای کد ۱۰۳",
		}

		p, err := repo.CreateProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, p)
	})

	// Empty color
	t.Run("color", func(t *testing.T) {

		p1 := model.Product{
			CompanyName: "Negin",
			CompanyId:   1,
			DesignCode:  "105",
			Colors:      []string{},
			Dimensions:  []string{"6", "9"},
			Description: "توضیحات برای کد ۱۰۵",
		}

		p, err := repo.CreateProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, p)
	})
}

func TestProductRepo_CreateProduct_Nil(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Nil color
	t.Run("color", func(t *testing.T) {

		p1 := model.Product{
			CompanyName: "Negin",
			CompanyId:   1,
			DesignCode:  "105",
			Colors:      nil,
			Dimensions:  []string{"6", "9"},
			Description: "توضیحات برای کد ۱۰۵",
		}

		p, err := repo.CreateProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, p)
	})

	// Nil size
	t.Run("size", func(t *testing.T) {

		p1 := model.Product{
			CompanyName: "Negin",
			CompanyId:   1,
			DesignCode:  "103",
			Colors:      []string{"آبی", "قرمز"},
			Dimensions:  nil,
			Description: "توضیحات برای کد ۱۰۳",
		}

		p, err := repo.CreateProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, p)
	})

}

func TestProductRepo_GetProductWithId_Ok(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	p1 := model.Product{
		Id:          0,
		CompanyName: "Negin",
		CompanyId:   1,
		DesignCode:  "102",
		Colors:      []string{"آبی", "قرمز"},
		Dimensions:  []string{"6", "9"},
		Description: "توضیحات برای کد ۱۰۲",
	}

	p, err := repo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, p)

	p1 = model.Product{
		Id:          0,
		CompanyName: "Negin",
		CompanyId:   1,
		DesignCode:  "103",
		Colors:      []string{"سبز"},
		Dimensions:  []string{"6", "12", "9"},
		Description: "توضیحات برای کد ۱۰۳",
	}

	p, err = repo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, p)

	gotProduct, err := repo.GetProductWithId(p.ID)
	checkEqualProduct(t, p, gotProduct)
}

func TestProductRepo_GetProductWithId_NotExist(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	gotProduct, err := repo.GetProductWithId(10)
	require.Nil(t, gotProduct)
	require.Equal(t, derror.ProductNotFound, err)
}

func TestProductRepo_GetProductWithId_Empty(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("color", func(t *testing.T) {
		p1 := model.Product{
			Id:          0,
			CompanyName: "Negin",
			CompanyId:   1,
			DesignCode:  "102",
			Colors:      []string{},
			Dimensions:  []string{"6", "9"},
			Description: "توضیحات برای کد ۱۰۲",
		}

		p, err := repo.CreateProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, p)

		gotProduct, err := repo.GetProductWithId(p.ID)
		require.Nil(t, err)
		checkEqualProduct(t, p, gotProduct)
	})

	t.Run("size", func(t *testing.T) {
		p1 := model.Product{
			Id:          0,
			CompanyName: "Negin",
			CompanyId:   1,
			DesignCode:  "103",
			Colors:      []string{"آبی", "قرمز"},
			Dimensions:  []string{},
			Description: "توضیحات برای کد ۱۰۳",
		}

		p, err := repo.CreateProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, p)

		gotProduct, err := repo.GetProductWithId(p.ID)
		checkEqualProduct(t, p, gotProduct)
	})

	t.Run("design code", func(t *testing.T) {
		p1 := model.Product{
			Id:          0,
			CompanyName: "Negin",
			CompanyId:   1,
			DesignCode:  "",
			Colors:      []string{"آبی", "قرمز"},
			Dimensions:  []string{"6", "9"},
			Description: "توضیحات برای کد خالی",
		}

		p, err := repo.CreateProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, p)

		gotProduct, err := repo.GetProductWithId(p.ID)
		require.Nil(t, err)
		checkEqualProduct(t, p, gotProduct)
	})
}

func TestProductRepo_DeleteProduct_Ok(t *testing.T) {
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
	err = repo.DeleteProduct(p.ID)
	require.Nil(t, err)
}

func TestProductRepo_DeleteProduct_Empty(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	companyId := uint(1)

	t.Run("color", func(t *testing.T) {
		p1 := model.Product{
			CompanyName: "Negin",
			CompanyId:   companyId,
			DesignCode:  "105",
			Colors:      []string{},
			Dimensions:  []string{"6", "9"},
			Description: "توضیحات برای کد ۱۰۵",
		}

		p, err := repo.CreateProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, p)
		err = repo.DeleteProduct(p.ID)
		require.Nil(t, err)
	})

	t.Run("size", func(t *testing.T) {
		p1 := model.Product{
			CompanyName: "Negin",
			CompanyId:   companyId,
			DesignCode:  "106",
			Colors:      []string{"قرمز", "آبی"},
			Dimensions:  []string{},
			Description: "توضیحات برای کد ۱۰۶",
		}

		p, err := repo.CreateProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, p)
		err = repo.DeleteProduct(p.ID)
		require.Nil(t, err)

	})

	t.Run("color_size", func(t *testing.T) {
		p1 := model.Product{
			CompanyName: "Negin",
			CompanyId:   companyId,
			DesignCode:  "107",
			Colors:      []string{},
			Dimensions:  []string{},
			Description: "توضیحات برای کد ۱۰۶",
		}

		p, err := repo.CreateProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, p)
		err = repo.DeleteProduct(p.ID)
		require.Nil(t, err)

	})

}

func TestProductRepo_DeleteProduct_NotExist(t *testing.T) {
	repo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	err = repo.DeleteProduct(100000)
	require.Equal(t, derror.ProductNotFound, err)
}

func TestProductRepo_EditProduct_Ok(t *testing.T) {
	//repo, err := NewProductRepoMock()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//companyId := uint(1)
	//
	//p1 := model.Product{
	//	CompanyName: "Negin",
	//	CompanyId:   companyId,
	//	DesignCode:  "105",
	//	Colors:      []string{"قرمز", "آبی"},
	//	Dimensions:  []string{"6", "9"},
	//	Description: "توضیحات برای کد ۱۰۵",
	//}
	//
	//p, err := repo.CreateProduct(p1)
	//assert.Nil(t, err)
	//assert.NotNil(t, p)
	//
	//p1.DesignCode = "107"
	//_, err = repo.EditProduct(&p1)
	//assert.Nil(t, err)
}

func checkEqualProduct(t *testing.T, expectedProduct, gotProduct *schema.Product) {
	require.Equal(t, expectedProduct.ID, gotProduct.ID, "id")
	require.Equal(t, expectedProduct.Description, gotProduct.Description, "description")
	require.Equal(t, expectedProduct.DesignCode, gotProduct.DesignCode, "design code")
	require.Equal(t, expectedProduct.CompanyId, gotProduct.CompanyId, "company id")
	require.Equal(t, len(expectedProduct.Dimensions), len(gotProduct.Dimensions), "dimension")
	require.Equal(t, len(expectedProduct.Themes), len(gotProduct.Themes), "theme")
}
