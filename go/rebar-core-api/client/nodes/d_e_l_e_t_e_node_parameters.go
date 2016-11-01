package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDELETENodeParams creates a new DELETENodeParams object
// with the default values initialized.
func NewDELETENodeParams() *DELETENodeParams {
	var ()
	return &DELETENodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDELETENodeParamsWithTimeout creates a new DELETENodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDELETENodeParamsWithTimeout(timeout time.Duration) *DELETENodeParams {
	var ()
	return &DELETENodeParams{

		timeout: timeout,
	}
}

// NewDELETENodeParamsWithContext creates a new DELETENodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewDELETENodeParamsWithContext(ctx context.Context) *DELETENodeParams {
	var ()
	return &DELETENodeParams{

		Context: ctx,
	}
}

/*DELETENodeParams contains all the parameters to send to the API endpoint
for the d e l e t e node operation typically these are written to a http.Request
*/
type DELETENodeParams struct {

	/*ID*/
	ID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the d e l e t e node params
func (o *DELETENodeParams) WithTimeout(timeout time.Duration) *DELETENodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the d e l e t e node params
func (o *DELETENodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the d e l e t e node params
func (o *DELETENodeParams) WithContext(ctx context.Context) *DELETENodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the d e l e t e node params
func (o *DELETENodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the d e l e t e node params
func (o *DELETENodeParams) WithID(id string) *DELETENodeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the d e l e t e node params
func (o *DELETENodeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DELETENodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}