// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// TicketingProjectFieldMapCasesEntityLogic - An unstructured object of key/value pairs describing the logic for applying the rule.
type TicketingProjectFieldMapCasesEntityLogic struct {
}

type TicketingProjectFieldMapCasesEntity struct {
	// An unstructured object of key/value pairs describing the logic for applying the rule.
	Logic         *TicketingProjectFieldMapCasesEntityLogic    `json:"logic,omitempty"`
	ExternalValue *TicketingProjectFieldMapExternalValueEntity `json:"external_value,omitempty"`
}

func (o *TicketingProjectFieldMapCasesEntity) GetLogic() *TicketingProjectFieldMapCasesEntityLogic {
	if o == nil {
		return nil
	}
	return o.Logic
}

func (o *TicketingProjectFieldMapCasesEntity) GetExternalValue() *TicketingProjectFieldMapExternalValueEntity {
	if o == nil {
		return nil
	}
	return o.ExternalValue
}