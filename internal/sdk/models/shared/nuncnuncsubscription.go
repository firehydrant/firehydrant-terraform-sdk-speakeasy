// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// NuncNuncSubscription - Nunc_NuncSubscription model
type NuncNuncSubscription struct {
	Response *string `json:"response,omitempty"`
}

func (o *NuncNuncSubscription) GetResponse() *string {
	if o == nil {
		return nil
	}
	return o.Response
}
