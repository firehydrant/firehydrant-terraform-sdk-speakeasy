// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// PatchV1IncidentRolesIncidentRoleID - Update a single incident role from its ID
type PatchV1IncidentRolesIncidentRoleID struct {
	Name        *string `json:"name,omitempty"`
	Summary     *string `json:"summary,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (o *PatchV1IncidentRolesIncidentRoleID) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PatchV1IncidentRolesIncidentRoleID) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *PatchV1IncidentRolesIncidentRoleID) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}