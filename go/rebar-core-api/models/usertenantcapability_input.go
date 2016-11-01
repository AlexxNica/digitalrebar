package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// UsertenantcapabilityInput usertenantcapability input
// swagger:model usertenantcapability-input
type UsertenantcapabilityInput struct {

	// capability id
	CapabilityID int64 `json:"capability_id,omitempty"`

	// tenant id
	TenantID int64 `json:"tenant_id,omitempty"`

	// user id
	UserID int64 `json:"user_id,omitempty"`
}

// Validate validates this usertenantcapability input
func (m *UsertenantcapabilityInput) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}