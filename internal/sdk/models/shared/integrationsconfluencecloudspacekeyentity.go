// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// IntegrationsConfluenceCloudSpaceKeyEntity - Integrations_ConfluenceCloud_SpaceKeyEntity model
type IntegrationsConfluenceCloudSpaceKeyEntity struct {
	Key  *string `json:"key,omitempty"`
	Name *string `json:"name,omitempty"`
}

func (o *IntegrationsConfluenceCloudSpaceKeyEntity) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *IntegrationsConfluenceCloudSpaceKeyEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}