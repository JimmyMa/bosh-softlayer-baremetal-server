package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*BaremetalServers baremetal servers

swagger:model BaremetalServers
*/
type BaremetalServers struct {

	/* data
	 */
	Data *BaremetalServer `json:"data,omitempty"`

	/* status
	 */
	Status *int32 `json:"status,omitempty"`
}

// Validate validates this baremetal servers
func (m *BaremetalServers) Validate(formats strfmt.Registry) error {
	return nil
}
