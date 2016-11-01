package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

/*LISTDeploymentsOK l i s t deployments o k

swagger:response lISTDeploymentsOK
*/
type LISTDeploymentsOK struct {

	// In: body
	Payload []*models.DeploymentsOutput `json:"body,omitempty"`
}

// NewLISTDeploymentsOK creates LISTDeploymentsOK with default headers values
func NewLISTDeploymentsOK() *LISTDeploymentsOK {
	return &LISTDeploymentsOK{}
}

// WithPayload adds the payload to the l i s t deployments o k response
func (o *LISTDeploymentsOK) WithPayload(payload []*models.DeploymentsOutput) *LISTDeploymentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the l i s t deployments o k response
func (o *LISTDeploymentsOK) SetPayload(payload []*models.DeploymentsOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LISTDeploymentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}