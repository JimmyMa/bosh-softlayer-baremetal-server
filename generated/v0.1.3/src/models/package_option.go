package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
)

/*PackageOption package option

swagger:model PackageOption
*/
type PackageOption struct {

	/* the code for this package option, e.g., server
	 */
	Code *string `json:"code,omitempty"`

	/* the name for this package option
	 */
	Name *string `json:"name,omitempty"`

	/* options
	 */
	Options []*Option `json:"options,omitempty"`
}

// Validate validates this package option
func (m *PackageOption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOptions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PackageOption) validateOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.Options) { // not required
		return nil
	}

	for i := 0; i < len(m.Options); i++ {

		if m.Options[i] != nil {

			if err := m.Options[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
