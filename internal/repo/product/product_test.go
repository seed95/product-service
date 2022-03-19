package product

import (
	"github.com/seed95/product-service/internal/derror"
	"github.com/seed95/product-service/internal/model"
	"github.com/seed95/product-service/internal/repo/product/schema"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProductRepo_CreateProduct_ZeroId(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	p1 := model.Product{
		Id:          0,
		CompanyName: "Negin",
		CompanyId:   1,
		DesignCode:  "102",
		Colors:      []string{"آبی", "قرمز"},
		Sizes:       []string{"6", "9"},
		Description: "توضیحات برای کد ۱۰۲",
	}

	p, err := pRepo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, p)
	require.NotEqual(t, p1.Id, p.ID)
	require.Equal(t, p1.Description, p.Description)
	require.Equal(t, p1.DesignCode, p.DesignCode)
	require.Equal(t, len(p1.Colors), len(p.Themes))
	require.Equal(t, len(p1.Sizes), len(p.Dimensions))
}

func TestProductRepo_CreateProduct_NonZeroId(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	p1 := model.Product{
		Id:          100,
		CompanyName: "Negin",
		CompanyId:   1,
		DesignCode:  "102",
		Colors:      []string{"آبی", "قرمز"},
		Sizes:       []string{"6", "9"},
		Description: "توضیحات برای کد ۱۰۲",
	}

	p, err := pRepo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, p)
	require.Equal(t, p1.Id, p.ID)
	require.Equal(t, p1.Description, p.Description)
	require.Equal(t, p1.DesignCode, p.DesignCode)
	require.Equal(t, len(p1.Colors), len(p.Themes))
	require.Equal(t, len(p1.Sizes), len(p.Dimensions))
}

func TestProductRepo_CreateProduct_Duplicate(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
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
			Sizes:       []string{"6", "9"},
			Description: "توضیحات برای کد ۱۰۲",
		}

		p, err := pRepo.CreateProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, p)

		p1 = model.Product{
			CompanyName: "Negin",
			CompanyId:   1,
			DesignCode:  "102",
			Colors:      []string{"سبز"},
			Sizes:       []string{"12"},
			Description: "توضیحات برای کد  تکراری ۱۰۲",
		}
		p, err = pRepo.CreateProduct(p1)
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
			Sizes:       []string{"6", "9", "6"},
			Description: "توضیحات برای کد ۱۰۳",
		}

		p, err := pRepo.CreateProduct(p1)
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
			Sizes:       []string{"6", "9"},
			Description: "توضیحات برای کد ۱۰۵",
		}

		p, err := pRepo.CreateProduct(p1)
		require.NotNil(t, err)
		require.Nil(t, p)
	})
}

func TestProductRepo_CreateProduct_Empty(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Empty product
	t.Run("product", func(t *testing.T) {
		p, err := pRepo.CreateProduct(model.Product{})
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
			Sizes:       []string{},
			Description: "توضیحات برای کد ۱۰۳",
		}

		p, err := pRepo.CreateProduct(p1)
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
			Sizes:       []string{},
			Description: "توضیحات برای کد خالی",
		}

		p, err := pRepo.CreateProduct(p1)
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
			Sizes:       []string{},
			Description: "توضیحات برای کد ۱۰۳",
		}

		p, err := pRepo.CreateProduct(p1)
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
			Sizes:       []string{"6", "9"},
			Description: "توضیحات برای کد ۱۰۵",
		}

		p, err := pRepo.CreateProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, p)
	})
}

func TestProductRepo_CreateProduct_Nil(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
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
			Sizes:       []string{"6", "9"},
			Description: "توضیحات برای کد ۱۰۵",
		}

		p, err := pRepo.CreateProduct(p1)
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
			Sizes:       nil,
			Description: "توضیحات برای کد ۱۰۳",
		}

		p, err := pRepo.CreateProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, p)
	})

}

func TestProductRepo_GetProductWithId_Ok(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)
	gotP2 := CreateProduct2(pRepo, t)

	gotProduct1, err := pRepo.GetProductWithId(gotP1.ID)
	require.Nil(t, err)
	require.NotNil(t, gotProduct1)
	checkEqualProduct(t, gotP1, gotProduct1)

	gotProduct2, err := pRepo.GetProductWithId(gotP2.ID)
	require.Nil(t, err)
	require.NotNil(t, gotProduct2)
	checkEqualProduct(t, gotP2, gotProduct2)
}

func TestProductRepo_GetProductWithId_NotExist(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	gotProduct, err := pRepo.GetProductWithId(10)
	require.Nil(t, gotProduct)
	require.Equal(t, derror.ProductNotFound, err)
}

func TestProductRepo_GetProductWithId_Empty(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
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
			Sizes:       []string{"6", "9"},
			Description: "توضیحات برای کد ۱۰۲",
		}

		gotP1, err := pRepo.CreateProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, gotP1)

		gotProduct, err := pRepo.GetProductWithId(gotP1.ID)
		require.Nil(t, err)
		checkEqualProduct(t, gotP1, gotProduct)
	})

	t.Run("size", func(t *testing.T) {
		p2 := model.Product{
			Id:          0,
			CompanyName: "Negin",
			CompanyId:   1,
			DesignCode:  "103",
			Colors:      []string{"آبی", "قرمز"},
			Sizes:       []string{},
			Description: "توضیحات برای کد ۱۰۳",
		}

		gotP2, err := pRepo.CreateProduct(p2)
		require.Nil(t, err)
		require.NotNil(t, gotP2)

		gotProduct, err := pRepo.GetProductWithId(gotP2.ID)
		require.Nil(t, err)
		checkEqualProduct(t, gotP2, gotProduct)
	})

	t.Run("design code", func(t *testing.T) {
		p3 := model.Product{
			Id:          0,
			CompanyName: "Negin",
			CompanyId:   1,
			DesignCode:  "",
			Colors:      []string{"آبی", "قرمز"},
			Sizes:       []string{"6", "9"},
			Description: "توضیحات برای کد خالی",
		}

		gotP3, err := pRepo.CreateProduct(p3)
		require.Nil(t, err)
		require.NotNil(t, gotP3)

		gotProduct, err := pRepo.GetProductWithId(gotP3.ID)
		require.Nil(t, err)
		checkEqualProduct(t, gotP3, gotProduct)
	})
}

func TestProductRepo_DeleteProduct_Ok(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)

	err = pRepo.DeleteProduct(gotP1.ID)
	require.Nil(t, err)

	gotP1, err = pRepo.GetProductWithId(gotP1.ID)
	require.Nil(t, gotP1)
	require.Equal(t, derror.ProductNotFound, err)
}

func TestProductRepo_DeleteProduct_Empty(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
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
			Sizes:       []string{"6", "9"},
			Description: "توضیحات برای کد ۱۰۵",
		}

		gotP1, err := pRepo.CreateProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, gotP1)

		err = pRepo.DeleteProduct(gotP1.ID)
		require.Nil(t, err)

		gotP1, err = pRepo.GetProductWithId(gotP1.ID)
		require.Nil(t, gotP1)
		require.Equal(t, derror.ProductNotFound, err)
	})

	t.Run("size", func(t *testing.T) {
		p2 := model.Product{
			CompanyName: "Negin",
			CompanyId:   companyId,
			DesignCode:  "106",
			Colors:      []string{"قرمز", "آبی"},
			Sizes:       []string{},
			Description: "توضیحات برای کد ۱۰۶",
		}

		gotP2, err := pRepo.CreateProduct(p2)
		require.Nil(t, err)
		require.NotNil(t, gotP2)

		err = pRepo.DeleteProduct(gotP2.ID)
		require.Nil(t, err)

		gotP2, err = pRepo.GetProductWithId(gotP2.ID)
		require.Nil(t, gotP2)
		require.Equal(t, derror.ProductNotFound, err)
	})

	t.Run("color_size", func(t *testing.T) {
		p3 := model.Product{
			CompanyName: "Negin",
			CompanyId:   companyId,
			DesignCode:  "107",
			Colors:      []string{},
			Sizes:       []string{},
			Description: "توضیحات برای کد ۱۰۷",
		}

		gotP3, err := pRepo.CreateProduct(p3)
		require.Nil(t, err)
		require.NotNil(t, gotP3)

		err = pRepo.DeleteProduct(gotP3.ID)
		require.Nil(t, err)

		gotP3, err = pRepo.GetProductWithId(gotP3.ID)
		require.Nil(t, gotP3)
		require.Equal(t, derror.ProductNotFound, err)
	})

}

func TestProductRepo_DeleteProduct_NotExist(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	err = pRepo.DeleteProduct(100000)
	require.Equal(t, derror.ProductNotFound, err)
}

func TestProductRepo_EditProduct_Ok(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("description", func(t *testing.T) {
		// Create product
		p1 := model.Product{
			CompanyName: "Negin",
			CompanyId:   1,
			DesignCode:  "105",
			Colors:      []string{"قرمز", "آبی"},
			Sizes:       []string{"6", "9"},
			Description: "توضیحات برای کد ۱۰۵",
		}

		gotP1, err := pRepo.CreateProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, gotP1)
		p1.Id = gotP1.ID

		p1.Description = "توضیحات عوض شدن"
		editedProduct, err := pRepo.EditProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, editedProduct)
		require.Equal(t, p1.Description, editedProduct.Description)
		require.Equal(t, p1.DesignCode, editedProduct.DesignCode)
	})

	t.Run("design code", func(t *testing.T) {
		// Create product
		p1 := model.Product{
			CompanyName: "Negin",
			CompanyId:   1,
			DesignCode:  "106",
			Colors:      []string{"قرمز", "آبی"},
			Sizes:       []string{"6", "9"},
			Description: "توضیحات برای کد ۱۰۶",
		}

		gotP1, err := pRepo.CreateProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, gotP1)
		p1.Id = gotP1.ID

		p1.DesignCode = "107"
		editedProduct, err := pRepo.EditProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, editedProduct)
		require.Equal(t, p1.Description, editedProduct.Description)
		require.Equal(t, p1.DesignCode, editedProduct.DesignCode)
	})

	t.Run("color", func(t *testing.T) {
		// Create product
		p1 := model.Product{
			CompanyName: "Negin",
			CompanyId:   1,
			DesignCode:  "108",
			Colors:      []string{"قرمز", "آبی"},
			Sizes:       []string{"6", "9"},
			Description: "توضیحات برای کد ۱۰۸",
		}

		gotP1, err := pRepo.CreateProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, gotP1)
		p1.Id = gotP1.ID

		p1.Colors = []string{"نارنجی", "صورتی", "قرمز"}
		editedProduct, err := pRepo.EditProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, editedProduct)
		require.Equal(t, len(p1.Colors), len(editedProduct.Themes))
	})

	t.Run("size", func(t *testing.T) {
		// Create product
		p1 := model.Product{
			CompanyName: "Negin",
			CompanyId:   1,
			DesignCode:  "109",
			Colors:      []string{"قرمز", "آبی"},
			Sizes:       []string{"6", "9"},
			Description: "توضیحات برای کد ۱۰۸",
		}

		gotP1, err := pRepo.CreateProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, gotP1)
		p1.Id = gotP1.ID

		p1.Sizes = []string{"15", "8", "6"}
		editedProduct, err := pRepo.EditProduct(p1)
		require.Nil(t, err)
		require.NotNil(t, editedProduct)
		require.Equal(t, len(p1.Sizes), len(editedProduct.Dimensions))
	})

}

func TestProductRepo_EditProduct_NotChange(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	p1 := model.Product{
		CompanyName: "Negin",
		CompanyId:   1,
		DesignCode:  "109",
		Colors:      []string{"قرمز", "آبی"},
		Sizes:       []string{"6", "9"},
		Description: "توضیحات برای کد ۱۰۸",
	}

	gotP1, err := pRepo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, gotP1)
	p1.Id = gotP1.ID

	editedProduct, err := pRepo.EditProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, editedProduct)
	require.NotEqual(t, gotP1.UpdatedAt, editedProduct.UpdatedAt)
}

func TestProductRepo_EditProduct_DuplicateDesignCode(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	p1 := model.Product{
		CompanyName: "Negin",
		CompanyId:   1,
		DesignCode:  "109",
		Colors:      []string{"قرمز", "آبی"},
		Sizes:       []string{"6", "9"},
		Description: "توضیحات برای کد ۱۰۸",
	}

	gotP1, err := pRepo.CreateProduct(p1)
	require.Nil(t, err)
	require.NotNil(t, gotP1)
	p1.Id = gotP1.ID

	gotP2 := CreateProduct2(pRepo, t)

	p1.DesignCode = gotP2.DesignCode
	editedProduct, err := pRepo.EditProduct(p1)
	require.NotNil(t, err)
	require.Nil(t, editedProduct)
}

func TestProductRepo_GetAllProducts_Ok(t *testing.T) {
	// NewProduct repo
	pRepo, err := NewProductRepoMock()
	if err != nil {
		t.Fatal(err)
	}

	// Create product
	gotP1 := CreateProduct1(pRepo, t)
	_ = CreateProduct2(pRepo, t)
	_ = CreateProduct3(pRepo, t)

	p := model.Product{
		CompanyName: "Almas",
		CompanyId:   2,
		DesignCode:  "105",
		Colors:      []string{"نارنجی"},
		Sizes:       []string{"12"},
		Description: "توضیحات ۱۰۵ الماس",
	}
	gotP, err := pRepo.CreateProduct(p)
	require.Nil(t, err)
	require.NotNil(t, gotP)

	products, err := pRepo.GetAllProducts(gotP1.CompanyId)
	require.Nil(t, err)
	require.NotNil(t, products)
	require.Equal(t, 3, len(products))
}

func checkEqualProduct(t *testing.T, expectedProduct, gotProduct *schema.Product) {
	require.Equal(t, expectedProduct.ID, gotProduct.ID, "id")
	require.Equal(t, expectedProduct.Description, gotProduct.Description, "description")
	require.Equal(t, expectedProduct.DesignCode, gotProduct.DesignCode, "design code")
	require.Equal(t, expectedProduct.CompanyId, gotProduct.CompanyId, "company id")
	require.Equal(t, len(expectedProduct.Dimensions), len(gotProduct.Dimensions), "dimension")
	require.Equal(t, len(expectedProduct.Themes), len(gotProduct.Themes), "theme")
}
