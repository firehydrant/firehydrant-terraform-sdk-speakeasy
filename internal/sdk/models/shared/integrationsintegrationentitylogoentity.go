// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type IntegrationsIntegrationEntityLogoEntity struct {
	LogoURL *string `json:"logo_url,omitempty"`
}

func (o *IntegrationsIntegrationEntityLogoEntity) GetLogoURL() *string {
	if o == nil {
		return nil
	}
	return o.LogoURL
}