package product

//
//import (
//	"github.com/seed95/product-service/internal/model"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/require"
//	"testing"
//)
//
//func TestCarpetRepo_GetAllCarpet_Ok(t *testing.T) {
//	repo, err := NewProductRepoMock()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	companyId := uint(1)
//
//	p1 := model.NewProduct{
//		CompanyName: "Negin",
//		CompanyId:   companyId,
//		DesignCode:  "105",
//		Colors:      []string{"قرمز", "آبی"},
//		Sizes:  []string{"6", "9"},
//		Description: "توضیحات برای کد ۱۰۵",
//	}
//
//	p, err := repo.CreateProduct(p1)
//	assert.Nil(t, err)
//	assert.NotNil(t, p)
//
//	carpets, err := repo.GetAllCarpet(companyId)
//	assert.Nil(t, err)
//	assert.Equal(t, 4, len(carpets))
//}
//
//func TestCarpetRepo_GetAllCarpet_EmptyColor(t *testing.T) {
//	repo, err := NewProductRepoMock()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	companyId := uint(1)
//
//	p1 := model.NewProduct{
//		CompanyName: "Negin",
//		CompanyId:   companyId,
//		DesignCode:  "105",
//		Colors:      []string{},
//		Sizes:  []string{"6", "9"},
//		Description: "توضیحات برای کد ۱۰۵",
//	}
//
//	p, err := repo.CreateProduct(p1)
//	assert.Nil(t, err)
//	assert.NotNil(t, p)
//
//	carpets, err := repo.GetAllCarpet(companyId)
//	assert.Nil(t, err)
//	assert.Equal(t, 2, len(carpets))
//}
//
//func TestCarpetRepo_GetAllCarpet_Empty(t *testing.T) {
//	repo, err := NewProductRepoMock()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	companyId := uint(1)
//
//	carpets, err := repo.GetAllCarpet(companyId)
//	assert.Nil(t, err)
//	assert.Equal(t, 0, len(carpets))
//}
//
//func TestCarpetRepo_GetAllCarpet_InvalidCompanyId(t *testing.T) {
//	repo, err := NewProductRepoMock()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	companyId := uint(0)
//
//	carpets, err := repo.GetAllCarpet(companyId)
//	assert.NotNil(t, err)
//	assert.Nil(t, carpets)
//
//}
//
//func TestCarpetRepo_GetAllCarpetWithProductId_Ok(t *testing.T) {
//	repo, err := NewProductRepoMock()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	companyId := uint(1)
//
//	p := model.NewProduct{
//		CompanyName: "Negin",
//		CompanyId:   companyId,
//		DesignCode:  "105",
//		Colors:      []string{"قرمز", "آبی"},
//		Sizes:  []string{"6", "9"},
//		Description: "توضیحات برای کد ۱۰۵",
//	}
//	p1, err := repo.CreateProduct(p)
//	require.Nil(t, err)
//	require.NotNil(t, p1)
//
//	p = model.NewProduct{
//		CompanyName: "Negin",
//		CompanyId:   companyId,
//		DesignCode:  "106",
//		Colors:      []string{"قرمز", "آبی"},
//		Sizes:  []string{"6", "12", "9"},
//		Description: "توضیحات برای کد ۱۰۶",
//	}
//	p2, err := repo.CreateProduct(p)
//	require.Nil(t, err)
//	require.NotNil(t, p2)
//
//	carpets, err := repo.GetAllCarpetWithProductId(companyId, p2.Id)
//	require.Nil(t, err)
//	require.Equal(t, len(p.Sizes)*len(p.Colors), len(carpets))
//}
//
//func TestCarpetRepo_GetAllCarpetWithProductId_NotExist(t *testing.T) {
//	repo, err := NewProductRepoMock()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	companyId := uint(1)
//
//	p := model.NewProduct{
//		CompanyName: "Negin",
//		CompanyId:   companyId,
//		DesignCode:  "105",
//		Colors:      []string{"قرمز", "آبی"},
//		Sizes:  []string{"6", "9"},
//		Description: "توضیحات برای کد ۱۰۵",
//	}
//	p1, err := repo.CreateProduct(p)
//	require.Nil(t, err)
//	require.NotNil(t, p1)
//
//	carpets, err := repo.GetAllCarpetWithProductId(companyId, p1.Id+100)
//	require.Nil(t, err)
//	require.Equal(t, 0, len(carpets))
//}
