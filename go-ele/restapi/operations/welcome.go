// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WelcomeHandlerFunc turns a function with the right signature into a welcome handler
type WelcomeHandlerFunc func(WelcomeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WelcomeHandlerFunc) Handle(params WelcomeParams) middleware.Responder {
	return fn(params)
}

// WelcomeHandler interface for that can handle valid welcome params
type WelcomeHandler interface {
	Handle(WelcomeParams) middleware.Responder
}

// NewWelcome creates a new http.Handler for the welcome operation
func NewWelcome(ctx *middleware.Context, handler WelcomeHandler) *Welcome {
	return &Welcome{Context: ctx, Handler: handler}
}

/* Welcome swagger:route GET /welcome welcome

welcome

Simple welcome message to verify service is up

*/
type Welcome struct {
	Context *middleware.Context
	Handler WelcomeHandler
}

func (o *Welcome) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWelcomeParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// WelcomeOKBody welcome o k body
//
// swagger:model WelcomeOKBody
type WelcomeOKBody struct {

	// msg
	// Required: true
	Msg *string `json:"msg"`
}

// Validate validates this welcome o k body
func (o *WelcomeOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMsg(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WelcomeOKBody) validateMsg(formats strfmt.Registry) error {

	if err := validate.Required("welcomeOK"+"."+"msg", "body", o.Msg); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this welcome o k body based on context it is used
func (o *WelcomeOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *WelcomeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *WelcomeOKBody) UnmarshalBinary(b []byte) error {
	var res WelcomeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
