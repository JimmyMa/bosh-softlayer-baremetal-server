package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*Info info

swagger:model Info
*/
type Info struct {

	/* data
	 */
	Data *InfoData `json:"data,omitempty"`

	/* status
	 */
	Status *int32 `json:"status,omitempty"`
}

// Validate validates this info
func (m *Info) Validate(formats strfmt.Registry) error {
	return nil
}
