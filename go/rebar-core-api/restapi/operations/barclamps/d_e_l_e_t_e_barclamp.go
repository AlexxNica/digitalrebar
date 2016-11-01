package barclamps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DELETEBarclampHandlerFunc turns a function with the right signature into a d e l e t e barclamp handler
type DELETEBarclampHandlerFunc func(DELETEBarclampParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DELETEBarclampHandlerFunc) Handle(params DELETEBarclampParams) middleware.Responder {
	return fn(params)
}

// DELETEBarclampHandler interface for that can handle valid d e l e t e barclamp params
type DELETEBarclampHandler interface {
	Handle(DELETEBarclampParams) middleware.Responder
}

// NewDELETEBarclamp creates a new http.Handler for the d e l e t e barclamp operation
func NewDELETEBarclamp(ctx *middleware.Context, handler DELETEBarclampHandler) *DELETEBarclamp {
	return &DELETEBarclamp{Context: ctx, Handler: handler}
}

/*DELETEBarclamp swagger:route DELETE /barclamps/{id} Barclamps dELETEBarclamp

Delete Barclamp

*/
type DELETEBarclamp struct {
	Context *middleware.Context
	Handler DELETEBarclampHandler
}

func (o *DELETEBarclamp) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewDELETEBarclampParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}