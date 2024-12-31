// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// PatchV1IntegrationsAwsConnectionsID - Update the AWS connection with the provided data.
type PatchV1IntegrationsAwsConnectionsID struct {
	AwsAccountID     *string `json:"aws_account_id,omitempty"`
	TargetArn        *string `json:"target_arn,omitempty"`
	ConnectionStatus *string `json:"connection_status,omitempty"`
}

func (o *PatchV1IntegrationsAwsConnectionsID) GetAwsAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AwsAccountID
}

func (o *PatchV1IntegrationsAwsConnectionsID) GetTargetArn() *string {
	if o == nil {
		return nil
	}
	return o.TargetArn
}

func (o *PatchV1IntegrationsAwsConnectionsID) GetConnectionStatus() *string {
	if o == nil {
		return nil
	}
	return o.ConnectionStatus
}