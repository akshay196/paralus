// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RPCLookupUserResponse rpc lookup user response
//
// swagger:model rpcLookupUserResponse
type RPCLookupUserResponse struct {

	// account ID
	AccountID string `json:"accountID,omitempty"`

	// is s s o
	IsSSO string `json:"isSSO,omitempty"`

	// organization ID
	OrganizationID string `json:"organizationID,omitempty"`

	// partner ID
	PartnerID string `json:"partnerID,omitempty"`

	// session type
	SessionType string `json:"sessionType,omitempty"`

	// user name
	UserName string `json:"userName,omitempty"`
}

// Validate validates this rpc lookup user response
func (m *RPCLookupUserResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this rpc lookup user response based on context it is used
func (m *RPCLookupUserResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RPCLookupUserResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RPCLookupUserResponse) UnmarshalBinary(b []byte) error {
	var res RPCLookupUserResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}