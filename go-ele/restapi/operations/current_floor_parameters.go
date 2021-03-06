// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewCurrentFloorParams creates a new CurrentFloorParams object
//
// There are no default values defined in the spec.
func NewCurrentFloorParams() CurrentFloorParams {

	return CurrentFloorParams{}
}

// CurrentFloorParams contains all the bound params for the current floor operation
// typically these are obtained from a http.Request
//
// swagger:parameters currentFloor
type CurrentFloorParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The ID of the car we are investigating
	  Required: true
	  In: path
	*/
	CarID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCurrentFloorParams() beforehand.
func (o *CurrentFloorParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rCarID, rhkCarID, _ := route.Params.GetOK("car_id")
	if err := o.bindCarID(rCarID, rhkCarID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCarID binds and validates parameter CarID from path.
func (o *CurrentFloorParams) bindCarID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.CarID = raw

	return nil
}
