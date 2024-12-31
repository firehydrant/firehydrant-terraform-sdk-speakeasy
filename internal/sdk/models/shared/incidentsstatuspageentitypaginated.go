// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// IncidentsStatusPageEntityPaginated - Incidents_StatusPageEntityPaginated model
type IncidentsStatusPageEntityPaginated struct {
	Data       []IncidentsStatusPageEntity `json:"data,omitempty"`
	Pagination *PaginationEntity           `json:"pagination,omitempty"`
}

func (o *IncidentsStatusPageEntityPaginated) GetData() []IncidentsStatusPageEntity {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *IncidentsStatusPageEntityPaginated) GetPagination() *PaginationEntity {
	if o == nil {
		return nil
	}
	return o.Pagination
}