package model

import "errors"

var (
	ErrInvalidNumberOfCarpet = errors.New("invalid_number_of_carpet")
	ErrInvalidCarpet         = errors.New("invalid_carpet")
)

type (
	Product struct {
		Id          uint
		CompanyName string
		CompanyId   uint
		DesignCode  string
		Colors      []string
		Dimensions  []string
		Description string
	}
)

func CarpetsToProduct(carpets []Carpet) (*Product, error) {

	result := Product{}

	//if len(carpets) == 0 {
	//	return &result, nil
	//}
	//
	//// Number of carpets should be a multiple of 2
	//if len(carpets)%2 != 0 {
	//	return nil, ErrInvalidNumberOfCarpet
	//}
	//
	//companyId := carpets[0].CompanyId
	//productId := carpets[0].ProductId
	//for _, c := range carpets {
	//	if c.ProductId != productId || c.CompanyId != companyId {
	//		return nil, ErrInvalidCarpet
	//	}
	//	result.Colors = append(result.Colors, c.Color)
	//	result.Dimensions = append(result.Dimensions, c.Dimension)
	//}

	return &result, nil

}
