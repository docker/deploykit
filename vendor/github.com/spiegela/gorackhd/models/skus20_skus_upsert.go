package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Skus20SkusUpsert A sku for RackHD
// swagger:model skus.2.0_SkusUpsert
type Skus20SkusUpsert struct {

	// discovery graph name
	// Min Length: 1
	DiscoveryGraphName string `json:"discoveryGraphName,omitempty"`

	// discovery graph options
	DiscoveryGraphOptions interface{} `json:"discoveryGraphOptions,omitempty"`

	// name
	// Min Length: 1
	Name string `json:"name,omitempty"`

	// Possible Rules a Sku can use
	Rules []*TagRule `json:"rules"`
}

// Validate validates this skus 2 0 skus upsert
func (m *Skus20SkusUpsert) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiscoveryGraphName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRules(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Skus20SkusUpsert) validateDiscoveryGraphName(formats strfmt.Registry) error {

	if swag.IsZero(m.DiscoveryGraphName) { // not required
		return nil
	}

	if err := validate.MinLength("discoveryGraphName", "body", string(m.DiscoveryGraphName), 1); err != nil {
		return err
	}

	return nil
}

func (m *Skus20SkusUpsert) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", string(m.Name), 1); err != nil {
		return err
	}

	return nil
}

func (m *Skus20SkusUpsert) validateRules(formats strfmt.Registry) error {

	if swag.IsZero(m.Rules) { // not required
		return nil
	}

	for i := 0; i < len(m.Rules); i++ {

		if swag.IsZero(m.Rules[i]) { // not required
			continue
		}

		if m.Rules[i] != nil {

			if err := m.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}
