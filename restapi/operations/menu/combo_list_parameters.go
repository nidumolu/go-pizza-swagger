// Code generated by go-swagger; DO NOT EDIT.

package menu

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewComboListParams creates a new ComboListParams object
//
// There are no default values defined in the spec.
func NewComboListParams() ComboListParams {

	return ComboListParams{}
}

// ComboListParams contains all the bound params for the combo list operation
// typically these are obtained from a http.Request
//
// swagger:parameters ComboList
type ComboListParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewComboListParams() beforehand.
func (o *ComboListParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
