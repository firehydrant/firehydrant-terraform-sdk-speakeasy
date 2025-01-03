// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// FunctionalityEntityPaginated model
type FunctionalityEntityPaginated struct {
	Data       []FunctionalityEntity `json:"data,omitempty"`
	Pagination *PaginationEntity     `json:"pagination,omitempty"`
}

func (o *FunctionalityEntityPaginated) GetData() []FunctionalityEntity {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *FunctionalityEntityPaginated) GetPagination() *PaginationEntity {
	if o == nil {
		return nil
	}
	return o.Pagination
}
