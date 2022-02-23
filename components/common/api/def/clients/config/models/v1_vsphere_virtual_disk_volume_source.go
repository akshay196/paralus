// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VsphereVirtualDiskVolumeSource Represents a vSphere volume resource.
//
// swagger:model v1VsphereVirtualDiskVolumeSource
type V1VsphereVirtualDiskVolumeSource struct {

	// Filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// +optional
	FsType string `json:"fsType,omitempty"`

	// Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.
	// +optional
	StoragePolicyID string `json:"storagePolicyID,omitempty"`

	// Storage Policy Based Management (SPBM) profile name.
	// +optional
	StoragePolicyName string `json:"storagePolicyName,omitempty"`

	// Path that identifies vSphere volume vmdk
	VolumePath string `json:"volumePath,omitempty"`
}

// Validate validates this v1 vsphere virtual disk volume source
func (m *V1VsphereVirtualDiskVolumeSource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 vsphere virtual disk volume source based on context it is used
func (m *V1VsphereVirtualDiskVolumeSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1VsphereVirtualDiskVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VsphereVirtualDiskVolumeSource) UnmarshalBinary(b []byte) error {
	var res V1VsphereVirtualDiskVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}