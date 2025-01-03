// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type UpdateAwsConnectionRequest struct {
	ID                                  string                                     `pathParam:"style=simple,explode=false,name=id"`
	PatchV1IntegrationsAwsConnectionsID shared.PatchV1IntegrationsAwsConnectionsID `request:"mediaType=application/json"`
}

func (o *UpdateAwsConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateAwsConnectionRequest) GetPatchV1IntegrationsAwsConnectionsID() shared.PatchV1IntegrationsAwsConnectionsID {
	if o == nil {
		return shared.PatchV1IntegrationsAwsConnectionsID{}
	}
	return o.PatchV1IntegrationsAwsConnectionsID
}

type UpdateAwsConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Update the AWS connection with the provided data.
	IntegrationsAwsConnectionEntity *shared.IntegrationsAwsConnectionEntity
	// A collection of codes that generally means the end user got something wrong in making the request
	BadRequest *shared.BadRequest
	// A collection of codes that generally means the client was not authenticated correctly for the request they want to make
	Unauthorized *shared.Unauthorized
	// Status codes relating to the resource/entity they are requesting not being found or endpoints/routes not existing
	NotFound *shared.NotFound
	// Status codes relating to the client being rate limited by the server
	RateLimited *shared.RateLimited
	// A collection of status codes that generally mean the server failed in an unexpected way
	InternalServerError *shared.InternalServerError
	// Timeouts occurred with the request
	Timeout *shared.Timeout
}

func (o *UpdateAwsConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateAwsConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateAwsConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateAwsConnectionResponse) GetIntegrationsAwsConnectionEntity() *shared.IntegrationsAwsConnectionEntity {
	if o == nil {
		return nil
	}
	return o.IntegrationsAwsConnectionEntity
}

func (o *UpdateAwsConnectionResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *UpdateAwsConnectionResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *UpdateAwsConnectionResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *UpdateAwsConnectionResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *UpdateAwsConnectionResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *UpdateAwsConnectionResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}
