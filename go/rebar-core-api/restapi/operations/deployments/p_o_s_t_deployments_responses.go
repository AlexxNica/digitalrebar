package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

/*POSTDeploymentsCreated p o s t deployments created

swagger:response pOSTDeploymentsCreated
*/
type POSTDeploymentsCreated struct {

	// In: body
	Payload *models.DeploymentsOutput `json:"body,omitempty"`
}

// NewPOSTDeploymentsCreated creates POSTDeploymentsCreated with default headers values
func NewPOSTDeploymentsCreated() *POSTDeploymentsCreated {
	return &POSTDeploymentsCreated{}
}

// WithPayload adds the payload to the p o s t deployments created response
func (o *POSTDeploymentsCreated) WithPayload(payload *models.DeploymentsOutput) *POSTDeploymentsCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the p o s t deployments created response
func (o *POSTDeploymentsCreated) SetPayload(payload *models.DeploymentsOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *POSTDeploymentsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}