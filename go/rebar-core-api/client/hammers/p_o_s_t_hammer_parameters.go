package hammers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

// NewPOSTHammerParams creates a new POSTHammerParams object
// with the default values initialized.
func NewPOSTHammerParams() *POSTHammerParams {
	var ()
	return &POSTHammerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPOSTHammerParamsWithTimeout creates a new POSTHammerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPOSTHammerParamsWithTimeout(timeout time.Duration) *POSTHammerParams {
	var ()
	return &POSTHammerParams{

		timeout: timeout,
	}
}

// NewPOSTHammerParamsWithContext creates a new POSTHammerParams object
// with the default values initialized, and the ability to set a context for a request
func NewPOSTHammerParamsWithContext(ctx context.Context) *POSTHammerParams {
	var ()
	return &POSTHammerParams{

		Context: ctx,
	}
}

/*POSTHammerParams contains all the parameters to send to the API endpoint
for the p o s t hammer operation typically these are written to a http.Request
*/
type POSTHammerParams struct {

	/*Body*/
	Body *models.HammerInput

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the p o s t hammer params
func (o *POSTHammerParams) WithTimeout(timeout time.Duration) *POSTHammerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the p o s t hammer params
func (o *POSTHammerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the p o s t hammer params
func (o *POSTHammerParams) WithContext(ctx context.Context) *POSTHammerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the p o s t hammer params
func (o *POSTHammerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the p o s t hammer params
func (o *POSTHammerParams) WithBody(body *models.HammerInput) *POSTHammerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the p o s t hammer params
func (o *POSTHammerParams) SetBody(body *models.HammerInput) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *POSTHammerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.HammerInput)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}