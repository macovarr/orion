// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// InventoryOKCode is the HTTP code returned for type InventoryOK
const InventoryOKCode int = 200

/*InventoryOK count

swagger:response inventoryOK
*/
type InventoryOK struct {

	/*
	  In: Body
	*/
	Payload []*InventoryOKBodyItems0 `json:"body,omitempty"`
}

// NewInventoryOK creates InventoryOK with default headers values
func NewInventoryOK() *InventoryOK {

	return &InventoryOK{}
}

// WithPayload adds the payload to the inventory o k response
func (o *InventoryOK) WithPayload(payload []*InventoryOKBodyItems0) *InventoryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the inventory o k response
func (o *InventoryOK) SetPayload(payload []*InventoryOKBodyItems0) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *InventoryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*InventoryOKBodyItems0, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// InventoryUnauthorizedCode is the HTTP code returned for type InventoryUnauthorized
const InventoryUnauthorizedCode int = 401

/*InventoryUnauthorized not authorized

swagger:response inventoryUnauthorized
*/
type InventoryUnauthorized struct {
}

// NewInventoryUnauthorized creates InventoryUnauthorized with default headers values
func NewInventoryUnauthorized() *InventoryUnauthorized {

	return &InventoryUnauthorized{}
}

// WriteResponse to the client
func (o *InventoryUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
