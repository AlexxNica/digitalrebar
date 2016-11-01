package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// BarclampInput barclamp input
// swagger:model barclamp-input
type BarclampInput struct {

	// barclamp id
	BarclampID int64 `json:"barclamp_id,omitempty"`

	// build on
	BuildOn string `json:"build_on,omitempty"`

	// build version
	BuildVersion string `json:"build_version,omitempty"`

	// cfg data
	CfgData *BarclampInputCfgData `json:"cfg_data,omitempty"`

	// commit
	Commit string `json:"commit,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// source path
	SourcePath string `json:"source_path,omitempty"`

	// source url
	SourceURL string `json:"source_url,omitempty"`

	// version
	Version int64 `json:"version,omitempty"`
}

// Validate validates this barclamp input
func (m *BarclampInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCfgData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BarclampInput) validateCfgData(formats strfmt.Registry) error {

	if swag.IsZero(m.CfgData) { // not required
		return nil
	}

	if m.CfgData != nil {

		if err := m.CfgData.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// BarclampInputCfgData barclamp input cfg data
// swagger:model BarclampInputCfgData
type BarclampInputCfgData struct {

	// barclamp
	Barclamp *BarclampInputCfgDataBarclamp `json:"barclamp,omitempty"`

	// jigs
	Jigs []*BarclampInputCfgDataJigsItems0 `json:"jigs,omitempty"`

	// os support
	OsSupport []string `json:"os_support,omitempty"`

	// rebar
	Rebar *BarclampInputCfgDataRebar `json:"rebar,omitempty"`
}

// Validate validates this barclamp input cfg data
func (m *BarclampInputCfgData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBarclamp(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateJigs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOsSupport(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRebar(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BarclampInputCfgData) validateBarclamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Barclamp) { // not required
		return nil
	}

	if m.Barclamp != nil {

		if err := m.Barclamp.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *BarclampInputCfgData) validateJigs(formats strfmt.Registry) error {

	if swag.IsZero(m.Jigs) { // not required
		return nil
	}

	for i := 0; i < len(m.Jigs); i++ {

		if swag.IsZero(m.Jigs[i]) { // not required
			continue
		}

		if m.Jigs[i] != nil {

			if err := m.Jigs[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *BarclampInputCfgData) validateOsSupport(formats strfmt.Registry) error {

	if swag.IsZero(m.OsSupport) { // not required
		return nil
	}

	return nil
}

func (m *BarclampInputCfgData) validateRebar(formats strfmt.Registry) error {

	if swag.IsZero(m.Rebar) { // not required
		return nil
	}

	if m.Rebar != nil {

		if err := m.Rebar.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// BarclampInputCfgDataBarclamp barclamp input cfg data barclamp
// swagger:model BarclampInputCfgDataBarclamp
type BarclampInputCfgDataBarclamp struct {

	// description
	Description string `json:"description,omitempty"`

	// display
	Display string `json:"display,omitempty"`

	// license
	License string `json:"license,omitempty"`

	// license url
	LicenseURL string `json:"license_url,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// source path
	SourcePath string `json:"source_path,omitempty"`

	// source url
	SourceURL string `json:"source_url,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this barclamp input cfg data barclamp
func (m *BarclampInputCfgDataBarclamp) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// BarclampInputCfgDataJigsItems0 barclamp input cfg data jigs items0
// swagger:model BarclampInputCfgDataJigsItems0
type BarclampInputCfgDataJigsItems0 struct {

	// class
	Class string `json:"class,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// implementor
	Implementor string `json:"implementor,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this barclamp input cfg data jigs items0
func (m *BarclampInputCfgDataJigsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// BarclampInputCfgDataRebar barclamp input cfg data rebar
// swagger:model BarclampInputCfgDataRebar
type BarclampInputCfgDataRebar struct {

	// layout
	Layout int64 `json:"layout,omitempty"`
}

// Validate validates this barclamp input cfg data rebar
func (m *BarclampInputCfgDataRebar) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}