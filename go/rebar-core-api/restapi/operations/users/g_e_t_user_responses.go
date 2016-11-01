package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

/*GETUserOK g e t user o k

swagger:response gETUserOK
*/
type GETUserOK struct {

	// In: body
	Payload *models.UserOutput `json:"body,omitempty"`
}

// NewGETUserOK creates GETUserOK with default headers values
func NewGETUserOK() *GETUserOK {
	return &GETUserOK{}
}

// WithPayload adds the payload to the g e t user o k response
func (o *GETUserOK) WithPayload(payload *models.UserOutput) *GETUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the g e t user o k response
func (o *GETUserOK) SetPayload(payload *models.UserOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GETUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}