// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PeerCreateResponse ID of created peer.
// swagger:model PeerCreateResponse
type PeerCreateResponse struct {

	// Peer ID of the node which dfget locates on.
	// Every peer has a unique ID among peer network.
	// It is generated via host's hostname and IP address.
	//
	ID string `json:"ID,omitempty"`
}

// Validate validates this peer create response
func (m *PeerCreateResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PeerCreateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PeerCreateResponse) UnmarshalBinary(b []byte) error {
	var res PeerCreateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
