package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*DELETETenantNoContent d e l e t e tenant no content

swagger:response dELETETenantNoContent
*/
type DELETETenantNoContent struct {
}

// NewDELETETenantNoContent creates DELETETenantNoContent with default headers values
func NewDELETETenantNoContent() *DELETETenantNoContent {
	return &DELETETenantNoContent{}
}

// WriteResponse to the client
func (o *DELETETenantNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}