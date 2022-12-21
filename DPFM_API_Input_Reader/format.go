package dpfm_api_input_reader

import (
	"data-platform-api-doc-type-exconf-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToDocType() *requests.DocType {
	data := sdc.DocType
	return &requests.DocType{
		DocType: data.DocType,
	}
}
