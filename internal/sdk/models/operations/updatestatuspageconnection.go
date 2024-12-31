// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type UpdateStatuspageConnectionRequest struct {
	// Connection UUID
	ConnectionID                                         string                                                      `pathParam:"style=simple,explode=false,name=connection_id"`
	PatchV1IntegrationsStatuspageConnectionsConnectionID shared.PatchV1IntegrationsStatuspageConnectionsConnectionID `request:"mediaType=application/json"`
}

func (o *UpdateStatuspageConnectionRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateStatuspageConnectionRequest) GetPatchV1IntegrationsStatuspageConnectionsConnectionID() shared.PatchV1IntegrationsStatuspageConnectionsConnectionID {
	if o == nil {
		return shared.PatchV1IntegrationsStatuspageConnectionsConnectionID{}
	}
	return o.PatchV1IntegrationsStatuspageConnectionsConnectionID
}

type UpdateStatuspageConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Update the given Statuspage integration connection.
	IntegrationsStatuspageConnectionEntity *shared.IntegrationsStatuspageConnectionEntity
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

func (o *UpdateStatuspageConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateStatuspageConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateStatuspageConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateStatuspageConnectionResponse) GetIntegrationsStatuspageConnectionEntity() *shared.IntegrationsStatuspageConnectionEntity {
	if o == nil {
		return nil
	}
	return o.IntegrationsStatuspageConnectionEntity
}

func (o *UpdateStatuspageConnectionResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *UpdateStatuspageConnectionResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *UpdateStatuspageConnectionResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *UpdateStatuspageConnectionResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *UpdateStatuspageConnectionResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *UpdateStatuspageConnectionResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}