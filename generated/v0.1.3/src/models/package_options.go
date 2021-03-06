package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*PackageOptions package options

swagger:model PackageOptions
*/
type PackageOptions struct {

	/* data
	 */
	Data *PackageOptionsData `json:"data,omitempty"`

	/* status
	 */
	Status *int32 `json:"status,omitempty"`
}

// Validate validates this package options
func (m *PackageOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PackageOptions) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {

		if err := m.Data.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

/*PackageOptionsData package options data

swagger:model PackageOptionsData
*/
type PackageOptionsData struct {

	/* categories
	 */
	Categories []*PackageOption `json:"categories,omitempty"`

	/* datacenters
	 */
	Datacenters []string `json:"datacenters,omitempty"`
}

// Validate validates this package options data
func (m *PackageOptionsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategories(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDatacenters(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PackageOptionsData) validateCategories(formats strfmt.Registry) error {

	if swag.IsZero(m.Categories) { // not required
		return nil
	}

	for i := 0; i < len(m.Categories); i++ {

		if m.Categories[i] != nil {

			if err := m.Categories[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *PackageOptionsData) validateDatacenters(formats strfmt.Registry) error {

	if swag.IsZero(m.Datacenters) { // not required
		return nil
	}

	for i := 0; i < len(m.Datacenters); i++ {

		if err := validate.RequiredString("data"+"."+"datacenters"+"."+strconv.Itoa(i), "body", string(m.Datacenters[i])); err != nil {
			return err
		}

	}

	return nil
}
