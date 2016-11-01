package profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*DELETEProfileNoContent d e l e t e profile no content

swagger:response dELETEProfileNoContent
*/
type DELETEProfileNoContent struct {
}

// NewDELETEProfileNoContent creates DELETEProfileNoContent with default headers values
func NewDELETEProfileNoContent() *DELETEProfileNoContent {
	return &DELETEProfileNoContent{}
}

// WriteResponse to the client
func (o *DELETEProfileNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}