package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// DeploymentsOutput deployments output
// swagger:model deployments-output
type DeploymentsOutput struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// state
	State int64 `json:"state,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`

	DeploymentsInput
}

// Validate validates this deployments output
func (m *DeploymentsOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.DeploymentsInput.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}