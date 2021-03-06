// Code generated by go-swagger; DO NOT EDIT.

package menu

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"go-pizza-swagger/models"
)

// ComboListOKCode is the HTTP code returned for type ComboListOK
const ComboListOKCode int = 200

/*ComboListOK List of Offers

swagger:response comboListOK
*/
type ComboListOK struct {

	/*
	  In: Body
	*/
	Payload models.Combos `json:"body,omitempty"`
}

// NewComboListOK creates ComboListOK with default headers values
func NewComboListOK() *ComboListOK {

	return &ComboListOK{}
}

// WithPayload adds the payload to the combo list o k response
func (o *ComboListOK) WithPayload(payload models.Combos) *ComboListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the combo list o k response
func (o *ComboListOK) SetPayload(payload models.Combos) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ComboListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Combos{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ComboListBadRequestCode is the HTTP code returned for type ComboListBadRequest
const ComboListBadRequestCode int = 400

/*ComboListBadRequest Bad Request

swagger:response comboListBadRequest
*/
type ComboListBadRequest struct {
}

// NewComboListBadRequest creates ComboListBadRequest with default headers values
func NewComboListBadRequest() *ComboListBadRequest {

	return &ComboListBadRequest{}
}

// WriteResponse to the client
func (o *ComboListBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// ComboListNotFoundCode is the HTTP code returned for type ComboListNotFound
const ComboListNotFoundCode int = 404

/*ComboListNotFound Categories Not Found

swagger:response comboListNotFound
*/
type ComboListNotFound struct {
}

// NewComboListNotFound creates ComboListNotFound with default headers values
func NewComboListNotFound() *ComboListNotFound {

	return &ComboListNotFound{}
}

// WriteResponse to the client
func (o *ComboListNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ComboListInternalServerErrorCode is the HTTP code returned for type ComboListInternalServerError
const ComboListInternalServerErrorCode int = 500

/*ComboListInternalServerError Server Error

swagger:response comboListInternalServerError
*/
type ComboListInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewComboListInternalServerError creates ComboListInternalServerError with default headers values
func NewComboListInternalServerError() *ComboListInternalServerError {

	return &ComboListInternalServerError{}
}

// WithPayload adds the payload to the combo list internal server error response
func (o *ComboListInternalServerError) WithPayload(payload string) *ComboListInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the combo list internal server error response
func (o *ComboListInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ComboListInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
