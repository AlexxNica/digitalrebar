package capabilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// LISTCapabilitiesHandlerFunc turns a function with the right signature into a l i s t capabilities handler
type LISTCapabilitiesHandlerFunc func(LISTCapabilitiesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn LISTCapabilitiesHandlerFunc) Handle(params LISTCapabilitiesParams) middleware.Responder {
	return fn(params)
}

// LISTCapabilitiesHandler interface for that can handle valid l i s t capabilities params
type LISTCapabilitiesHandler interface {
	Handle(LISTCapabilitiesParams) middleware.Responder
}

// NewLISTCapabilities creates a new http.Handler for the l i s t capabilities operation
func NewLISTCapabilities(ctx *middleware.Context, handler LISTCapabilitiesHandler) *LISTCapabilities {
	return &LISTCapabilities{Context: ctx, Handler: handler}
}

/*LISTCapabilities swagger:route GET /capabilities Capabilities lISTCapabilities

List Capabilities

*/
type LISTCapabilities struct {
	Context *middleware.Context
	Handler LISTCapabilitiesHandler
}

func (o *LISTCapabilities) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewLISTCapabilitiesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}