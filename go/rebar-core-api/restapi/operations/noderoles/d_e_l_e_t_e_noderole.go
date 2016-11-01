package noderoles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DELETENoderoleHandlerFunc turns a function with the right signature into a d e l e t e noderole handler
type DELETENoderoleHandlerFunc func(DELETENoderoleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DELETENoderoleHandlerFunc) Handle(params DELETENoderoleParams) middleware.Responder {
	return fn(params)
}

// DELETENoderoleHandler interface for that can handle valid d e l e t e noderole params
type DELETENoderoleHandler interface {
	Handle(DELETENoderoleParams) middleware.Responder
}

// NewDELETENoderole creates a new http.Handler for the d e l e t e noderole operation
func NewDELETENoderole(ctx *middleware.Context, handler DELETENoderoleHandler) *DELETENoderole {
	return &DELETENoderole{Context: ctx, Handler: handler}
}

/*DELETENoderole swagger:route DELETE /noderoles/{id} Noderoles dELETENoderole

Delete NodeRole

*/
type DELETENoderole struct {
	Context *middleware.Context
	Handler DELETENoderoleHandler
}

func (o *DELETENoderole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewDELETENoderoleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}