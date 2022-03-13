package model

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
