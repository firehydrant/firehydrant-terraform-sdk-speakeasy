// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// PostV1IncidentsIncidentIDTasks - Create a task
type PostV1IncidentsIncidentIDTasks struct {
	// The title of the task.
	Title string `json:"title"`
	// The state of the task. One of: open, in_progress, cancelled, done
	State *string `json:"state,omitempty"`
	// A description of what the task is for.
	Description *string `json:"description,omitempty"`
	// The ID of the user assigned to the task.
	AssigneeID *string `json:"assignee_id,omitempty"`
	// The due date of the task. Relative values are supported such as '5m'
	DueAt *string `json:"due_at,omitempty"`
}

func (o *PostV1IncidentsIncidentIDTasks) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *PostV1IncidentsIncidentIDTasks) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *PostV1IncidentsIncidentIDTasks) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PostV1IncidentsIncidentIDTasks) GetAssigneeID() *string {
	if o == nil {
		return nil
	}
	return o.AssigneeID
}

func (o *PostV1IncidentsIncidentIDTasks) GetDueAt() *string {
	if o == nil {
		return nil
	}
	return o.DueAt
}