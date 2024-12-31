// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// IncidentRoleEntityPaginated model
type IncidentRoleEntityPaginated struct {
	Data       []IncidentRoleEntity `json:"data,omitempty"`
	Pagination *PaginationEntity    `json:"pagination,omitempty"`
}

func (o *IncidentRoleEntityPaginated) GetData() []IncidentRoleEntity {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *IncidentRoleEntityPaginated) GetPagination() *PaginationEntity {
	if o == nil {
		return nil
	}
	return o.Pagination
}