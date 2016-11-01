package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

/*LISTUsersOK l i s t users o k

swagger:response lISTUsersOK
*/
type LISTUsersOK struct {

	// In: body
	Payload []*models.UserOutput `json:"body,omitempty"`
}

// NewLISTUsersOK creates LISTUsersOK with default headers values
func NewLISTUsersOK() *LISTUsersOK {
	return &LISTUsersOK{}
}

// WithPayload adds the payload to the l i s t users o k response
func (o *LISTUsersOK) WithPayload(payload []*models.UserOutput) *LISTUsersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the l i s t users o k response
func (o *LISTUsersOK) SetPayload(payload []*models.UserOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LISTUsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}