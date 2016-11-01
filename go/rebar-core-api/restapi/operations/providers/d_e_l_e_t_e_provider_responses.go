package providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*DELETEProviderNoContent d e l e t e provider no content

swagger:response dELETEProviderNoContent
*/
type DELETEProviderNoContent struct {
}

// NewDELETEProviderNoContent creates DELETEProviderNoContent with default headers values
func NewDELETEProviderNoContent() *DELETEProviderNoContent {
	return &DELETEProviderNoContent{}
}

// WriteResponse to the client
func (o *DELETEProviderNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}