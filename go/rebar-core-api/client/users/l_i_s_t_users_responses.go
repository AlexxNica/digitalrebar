package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

// LISTUsersReader is a Reader for the LISTUsers structure.
type LISTUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LISTUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLISTUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLISTUsersOK creates a LISTUsersOK with default headers values
func NewLISTUsersOK() *LISTUsersOK {
	return &LISTUsersOK{}
}

/*LISTUsersOK handles this case with default header values.

LISTUsersOK l i s t users o k
*/
type LISTUsersOK struct {
	Payload []*models.UserOutput
}

func (o *LISTUsersOK) Error() string {
	return fmt.Sprintf("[GET /users][%d] lISTUsersOK  %+v", 200, o.Payload)
}

func (o *LISTUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}