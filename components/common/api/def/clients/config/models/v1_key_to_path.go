// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1KeyToPath Maps a string key to a path within a volume.
//
// swagger:model v1KeyToPath
type V1KeyToPath struct {

	// The key to project.
	Key string `json:"key,omitempty"`

	// Optional: mode bits used to set permissions on this file.
	// Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.
	// YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.
	// If not specified, the volume defaultMode will be used.
	// This might be in conflict with other options that affect the file
	// mode, like fsGroup, and the result can be other mode bits set.
	// +optional
	Mode int32 `json:"mode,omitempty"`

	// The relative path of the file to map the key to.
	// May not be an absolute path.
	// May not contain the path element '..'.
	// May not start with the string '..'.
	Path string `json:"path,omitempty"`
}

// Validate validates this v1 key to path
func (m *V1KeyToPath) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 key to path based on context it is used
func (m *V1KeyToPath) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1KeyToPath) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1KeyToPath) UnmarshalBinary(b []byte) error {
	var res V1KeyToPath
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}