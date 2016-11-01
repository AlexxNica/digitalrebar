package networkranges

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewLISTNetworkrangesParams creates a new LISTNetworkrangesParams object
// with the default values initialized.
func NewLISTNetworkrangesParams() LISTNetworkrangesParams {
	var ()
	return LISTNetworkrangesParams{}
}

// LISTNetworkrangesParams contains all the bound params for the l i s t networkranges operation
// typically these are obtained from a http.Request
//
// swagger:parameters LIST-networkranges
type LISTNetworkrangesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *LISTNetworkrangesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}