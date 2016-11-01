package deploymentroles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

/*POSTDeploymentrolesCreated p o s t deploymentroles created

swagger:response pOSTDeploymentrolesCreated
*/
type POSTDeploymentrolesCreated struct {

	// In: body
	Payload *models.DeploymentrolesOutput `json:"body,omitempty"`
}

// NewPOSTDeploymentrolesCreated creates POSTDeploymentrolesCreated with default headers values
func NewPOSTDeploymentrolesCreated() *POSTDeploymentrolesCreated {
	return &POSTDeploymentrolesCreated{}
}

// WithPayload adds the payload to the p o s t deploymentroles created response
func (o *POSTDeploymentrolesCreated) WithPayload(payload *models.DeploymentrolesOutput) *POSTDeploymentrolesCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the p o s t deploymentroles created response
func (o *POSTDeploymentrolesCreated) SetPayload(payload *models.DeploymentrolesOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *POSTDeploymentrolesCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}