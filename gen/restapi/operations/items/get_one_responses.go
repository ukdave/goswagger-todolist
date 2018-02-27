// Code generated by go-swagger; DO NOT EDIT.

package items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ukdave/goswagger-todolist/gen/models"
)

// GetOneOKCode is the HTTP code returned for type GetOneOK
const GetOneOKCode int = 200

/*GetOneOK OK

swagger:response getOneOK
*/
type GetOneOK struct {

	/*
	  In: Body
	*/
	Payload *models.Item `json:"body,omitempty"`
}

// NewGetOneOK creates GetOneOK with default headers values
func NewGetOneOK() *GetOneOK {
	return &GetOneOK{}
}

// WithPayload adds the payload to the get one o k response
func (o *GetOneOK) WithPayload(payload *models.Item) *GetOneOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get one o k response
func (o *GetOneOK) SetPayload(payload *models.Item) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOneOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetOneDefault error

swagger:response getOneDefault
*/
type GetOneDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetOneDefault creates GetOneDefault with default headers values
func NewGetOneDefault(code int) *GetOneDefault {
	if code <= 0 {
		code = 500
	}

	return &GetOneDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get one default response
func (o *GetOneDefault) WithStatusCode(code int) *GetOneDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get one default response
func (o *GetOneDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get one default response
func (o *GetOneDefault) WithPayload(payload *models.Error) *GetOneDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get one default response
func (o *GetOneDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOneDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
