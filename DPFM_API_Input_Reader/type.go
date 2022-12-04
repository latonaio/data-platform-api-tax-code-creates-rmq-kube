package dpfm_api_input_reader

type EC_MC struct {
}

type SDC struct {
	ConnectionKey       string   `json:"connection_key"`
	Result              bool     `json:"result"`
	RedisKey            string   `json:"redis_key"`
	Filepath            string   `json:"filepath"`
	APIStatusCode       int      `json:"api_status_code"`
	RuntimeSessionID    string   `json:"runtime_session_id"`
	BusinessPartnerID   *int     `json:"business_partner"`
	ServiceLabel        string   `json:"service_label"`
	TaxCode             TaxCode  `json:"TaxCode"`
	APISchema           string   `json:"api_schema"`
	Accepter            []string `json:"accepter"`
	OrderID             *int     `json:"order_id"`
	Deleted             bool     `json:"deleted"`
	SQLUpdateResult     *bool    `json:"sql_update_result"`
	SQLUpdateError      string   `json:"sql_update_error"`
	SubfuncResult       *bool    `json:"subfunc_result"`
	SubfuncError        string   `json:"subfunc_error"`
	ExconfResult        *bool    `json:"exconf_result"`
	ExconfError         string   `json:"exconf_error"`
	APIProcessingResult *bool    `json:"api_processing_result"`
	APIProcessingError  string   `json:"api_processing_error"`
}

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
	Country           string   `json:"Country"`
	TaxCode           string   `json:"TaxCode"`
	ValidityEndDate   *string  `json:"ValidityEndDate"`
	ValidityStartDate *string  `json:"ValidityStartDate"`
	TaxRate           *float32 `json:"TaxRate"`
}
