// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PreheatInfo return detailed information of a preheat task in supernode. An image preheat task may contain multiple downloading
// task because that an image may have more than one layer.
//
// swagger:model PreheatInfo
type PreheatInfo struct {

	// ID of preheat task.
	//
	ID string `json:"ID,omitempty"`

	// the preheat task finish time
	// Format: date-time
	FinishTime strfmt.DateTime `json:"finishTime,omitempty"`

	// the preheat task start time
	// Format: date-time
	StartTime strfmt.DateTime `json:"startTime,omitempty"`

	// The status of preheat task.
	//   WAITING -----> RUNNING -----> SUCCESS
	//                            |--> FAILED
	// The initial status of a created preheat task is WAITING.
	// It's finished when a preheat task's status is FAILED or SUCCESS.
	// A finished preheat task's information can be queried within 24 hours.
	//
	Status PreheatStatus `json:"status,omitempty"`

	// the error message of preheat task when failed
	ErrorMsg string `json:"errorMsg,omitempty"`
}

// Validate validates this preheat info
func (m *PreheatInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFinishTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PreheatInfo) validateFinishTime(formats strfmt.Registry) error {

	if swag.IsZero(m.FinishTime) { // not required
		return nil
	}

	if err := validate.FormatOf("finishTime", "body", "date-time", m.FinishTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PreheatInfo) validateStartTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PreheatInfo) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PreheatInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PreheatInfo) UnmarshalBinary(b []byte) error {
	var res PreheatInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
