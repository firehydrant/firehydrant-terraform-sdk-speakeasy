// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type TicketingProjectFieldMapCasesElseEntity struct {
	ExternalValue *TicketingProjectFieldMapExternalValueEntity `json:"external_value,omitempty"`
}

func (o *TicketingProjectFieldMapCasesElseEntity) GetExternalValue() *TicketingProjectFieldMapExternalValueEntity {
	if o == nil {
		return nil
	}
	return o.ExternalValue
}