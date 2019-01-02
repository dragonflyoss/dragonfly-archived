// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PieceInfo Peer's detailed information in supernode.
// swagger:model PieceInfo
type PieceInfo struct {

	// ID of the peer
	ID string `json:"ID,omitempty"`
}

// Validate validates this piece info
func (m *PieceInfo) Validate(formats strfmt.Registry) (err error) {
	return err
}

// MarshalBinary interface implementation
func (m *PieceInfo) MarshalBinary() (b []byte, err error) {
	if m == nil {
		return
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PieceInfo) UnmarshalBinary(b []byte) (err error) {
	var res PieceInfo
	if err = swag.ReadJSON(b, &res); err != nil {
		return
	}
	*m = res
	return
}
