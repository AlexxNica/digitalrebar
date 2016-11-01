package deploymentroles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DELETEDeploymentrolesReader is a Reader for the DELETEDeploymentroles structure.
type DELETEDeploymentrolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DELETEDeploymentrolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDELETEDeploymentrolesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDELETEDeploymentrolesNoContent creates a DELETEDeploymentrolesNoContent with default headers values
func NewDELETEDeploymentrolesNoContent() *DELETEDeploymentrolesNoContent {
	return &DELETEDeploymentrolesNoContent{}
}

/*DELETEDeploymentrolesNoContent handles this case with default header values.

DELETEDeploymentrolesNoContent d e l e t e deploymentroles no content
*/
type DELETEDeploymentrolesNoContent struct {
}

func (o *DELETEDeploymentrolesNoContent) Error() string {
	return fmt.Sprintf("[DELETE /deploymentroles/{id}][%d] dELETEDeploymentrolesNoContent ", 204)
}

func (o *DELETEDeploymentrolesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}