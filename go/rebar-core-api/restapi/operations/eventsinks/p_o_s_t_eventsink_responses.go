package eventsinks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

/*POSTEventsinkCreated p o s t eventsink created

swagger:response pOSTEventsinkCreated
*/
type POSTEventsinkCreated struct {

	// In: body
	Payload *models.EventsinkOutput `json:"body,omitempty"`
}

// NewPOSTEventsinkCreated creates POSTEventsinkCreated with default headers values
func NewPOSTEventsinkCreated() *POSTEventsinkCreated {
	return &POSTEventsinkCreated{}
}

// WithPayload adds the payload to the p o s t eventsink created response
func (o *POSTEventsinkCreated) WithPayload(payload *models.EventsinkOutput) *POSTEventsinkCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the p o s t eventsink created response
func (o *POSTEventsinkCreated) SetPayload(payload *models.EventsinkOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *POSTEventsinkCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}