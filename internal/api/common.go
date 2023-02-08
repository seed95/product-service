package api

type Common struct {
	Language    string `json:"language"`
	Username    string `json:"username"`
	CompanyId   int64  `json:"company_id"`
	CompanyName string `json:"company_name"`
}
