package dpfm_api_input_reader

import (
	"data-platform-api-tax-code-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToTaxCode() *requests.TaxCode {
	data := sdc.TaxCode
	return &requests.TaxCode{
		Country: data.Country,
		TaxCode: data.TaxCode,
	}
}

func (sdc *SDC) ConvertToTaxCodeText() *requests.TaxCodeText {
	dataTaxCode := sdc.TaxCode
	data := sdc.TaxCode.TaxCodeText
	return &requests.TaxCodeText{
		Country:     dataTaxCode.Country,
		TaxCode:     dataTaxCode.TaxCode,
		Language:    data.Language,
		TaxCodeName: data.TaxCodeName,
	}
}

func (sdc *SDC) ConvertToTaxRate() *requests.TaxRate {
	dataTaxCode := sdc.TaxCode
	data := sdc.TaxCode.TaxRate
	return &requests.TaxRate{
		Country:           dataTaxCode.Country,
		TaxCode:           dataTaxCode.TaxCode,
		ValidityEndDate:   data.ValidityEndDate,
		ValidityStartDate: data.ValidityStartDate,
		TaxRate:           data.TaxRate,
	}
}
