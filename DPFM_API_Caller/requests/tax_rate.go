package requests

type TaxRate struct {
	Country           string   `json:"Country"`
	TaxCode           string   `json:"TaxCode"`
	ValidityEndDate   *string  `json:"ValidityEndDate"`
	ValidityStartDate *string  `json:"ValidityStartDate"`
	TaxRate           *float32 `json:"TaxRate"`
}
