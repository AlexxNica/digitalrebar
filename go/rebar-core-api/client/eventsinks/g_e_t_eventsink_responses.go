package eventsinks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

// GETEventsinkReader is a Reader for the GETEventsink structure.
type GETEventsinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETEventsinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGETEventsinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGETEventsinkOK creates a GETEventsinkOK with default headers values
func NewGETEventsinkOK() *GETEventsinkOK {
	return &GETEventsinkOK{}
}

/*GETEventsinkOK handles this case with default header values.

GETEventsinkOK g e t eventsink o k
*/
type GETEventsinkOK struct {
	Payload *models.EventsinkOutput
}

func (o *GETEventsinkOK) Error() string {
	return fmt.Sprintf("[GET /eventsinks/{id}][%d] gETEventsinkOK  %+v", 200, o.Payload)
}

func (o *GETEventsinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EventsinkOutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}