package requests

type TaxCodeText struct {
	Country     string `json:"Country"`
	TaxCode     string `json:"TaxCode"`
	Language    string `json:"Language"`
	TaxCodeName string `json:"TaxCodeName"`
}
