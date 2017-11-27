// Code generated by go-swagger; DO NOT EDIT.

package jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"3_simple/models"
)

// FindJobsOKCode is the HTTP code returned for type FindJobsOK
const FindJobsOKCode int = 200

/*FindJobsOK list all created jobs

swagger:response findJobsOK
*/
type FindJobsOK struct {

	/*
	  In: Body
	*/
	Payload models.FindJobsOKBody `json:"body,omitempty"`
}

// NewFindJobsOK creates FindJobsOK with default headers values
func NewFindJobsOK() *FindJobsOK {
	return &FindJobsOK{}
}

// WithPayload adds the payload to the find jobs o k response
func (o *FindJobsOK) WithPayload(payload models.FindJobsOKBody) *FindJobsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find jobs o k response
func (o *FindJobsOK) SetPayload(payload models.FindJobsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindJobsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.FindJobsOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*FindJobsDefault generic error response

swagger:response findJobsDefault
*/
type FindJobsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewFindJobsDefault creates FindJobsDefault with default headers values
func NewFindJobsDefault(code int) *FindJobsDefault {
	if code <= 0 {
		code = 500
	}

	return &FindJobsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the find jobs default response
func (o *FindJobsDefault) WithStatusCode(code int) *FindJobsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the find jobs default response
func (o *FindJobsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the find jobs default response
func (o *FindJobsDefault) WithPayload(payload *models.Error) *FindJobsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find jobs default response
func (o *FindJobsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindJobsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
