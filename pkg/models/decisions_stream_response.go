// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DecisionsStreamResponse DecisionsStreamResponse
//
// swagger:model DecisionsStreamResponse
type DecisionsStreamResponse struct {

	// deleted
	Deleted GetDecisionsResponse `json:"deleted,omitempty"`

	// new
	New GetDecisionsResponse `json:"new,omitempty"`
}

// Validate validates this decisions stream response
func (m *DecisionsStreamResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNew(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DecisionsStreamResponse) validateDeleted(formats strfmt.Registry) error {

	if swag.IsZero(m.Deleted) { // not required
		return nil
	}

	if err := m.Deleted.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deleted")
		}
		return err
	}

	return nil
}

func (m *DecisionsStreamResponse) validateNew(formats strfmt.Registry) error {

	if swag.IsZero(m.New) { // not required
		return nil
	}

	if err := m.New.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("new")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DecisionsStreamResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DecisionsStreamResponse) UnmarshalBinary(b []byte) error {
	var res DecisionsStreamResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
