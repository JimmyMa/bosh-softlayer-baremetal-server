package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-swagger/go-swagger/httpkit/middleware"
)

// GetSlPackagePackageIDOptionsHandlerFunc turns a function with the right signature into a get sl package package ID options handler
type GetSlPackagePackageIDOptionsHandlerFunc func(GetSlPackagePackageIDOptionsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSlPackagePackageIDOptionsHandlerFunc) Handle(params GetSlPackagePackageIDOptionsParams) middleware.Responder {
	return fn(params)
}

// GetSlPackagePackageIDOptionsHandler interface for that can handle valid get sl package package ID options params
type GetSlPackagePackageIDOptionsHandler interface {
	Handle(GetSlPackagePackageIDOptionsParams) middleware.Responder
}

// NewGetSlPackagePackageIDOptions creates a new http.Handler for the get sl package package ID options operation
func NewGetSlPackagePackageIDOptions(ctx *middleware.Context, handler GetSlPackagePackageIDOptionsHandler) *GetSlPackagePackageIDOptions {
	return &GetSlPackagePackageIDOptions{Context: ctx, Handler: handler}
}

/*GetSlPackagePackageIDOptions swagger:route GET /sl/package/{packageId}/options getSlPackagePackageIdOptions

List SoftLayer package options


*/
type GetSlPackagePackageIDOptions struct {
	Context *middleware.Context
	Handler GetSlPackagePackageIDOptionsHandler
}

func (o *GetSlPackagePackageIDOptions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetSlPackagePackageIDOptionsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
