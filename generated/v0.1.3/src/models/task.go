package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*Task task

swagger:model Task
*/
type Task struct {

	/* the description string for this task
	 */
	Description *string `json:"description,omitempty"`

	/* id
	 */
	ID *int32 `json:"id,omitempty"`

	/* the start time for task, e.g., 2016-07-05T19:38:08+08:00
	 */
	StartTime *string `json:"start_time,omitempty"`

	/* the status string for this task
	 */
	Status *string `json:"status,omitempty"`
}

// Validate validates this task
func (m *Task) Validate(formats strfmt.Registry) error {
	return nil
}