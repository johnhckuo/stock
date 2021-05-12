// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Timeslot timeslot
//
// swagger:model Timeslot
type Timeslot struct {

	// end at
	// Required: true
	EndAt *int64 `json:"endAt"`

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// start at
	// Required: true
	StartAt *int64 `json:"startAt"`

	// user Id
	// Required: true
	UserID *int64 `json:"userId"`
}

// Validate validates this timeslot
func (m *Timeslot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Timeslot) validateEndAt(formats strfmt.Registry) error {

	if err := validate.Required("endAt", "body", m.EndAt); err != nil {
		return err
	}

	return nil
}

func (m *Timeslot) validateStartAt(formats strfmt.Registry) error {

	if err := validate.Required("startAt", "body", m.StartAt); err != nil {
		return err
	}

	return nil
}

func (m *Timeslot) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this timeslot based on the context it is used
func (m *Timeslot) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Timeslot) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Timeslot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Timeslot) UnmarshalBinary(b []byte) error {
	var res Timeslot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
