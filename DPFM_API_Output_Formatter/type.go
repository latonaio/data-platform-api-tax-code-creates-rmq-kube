package dpfm_api_output_formatter

type TaxCode struct {
	Country     string      `json:"Country"`
	TaxCode     string      `json:"TaxCode"`
	TaxCodeText TaxCodeText `json:"TaxCodeText"`
	TaxRate     TaxRate     `json:"TaxRate"`
}

type TaxCodeText struct {
	Country     string `json:"Country"`
	TaxCode     string `json:"TaxCode"`
	Language    string `json:"Language"`
	TaxCodeName string `json:"TaxCodeName"`
}

type TaxRate struct {
	Country           string  `json:"Country"`
	TaxCode           string  `json:"TaxCode"`
	ValidityEndDate   string  `json:"ValidityEndDate"`
	ValidityStartDate string  `json:"ValidityStartDate"`
	TaxRate           float32 `json:"TaxRate"`
}
