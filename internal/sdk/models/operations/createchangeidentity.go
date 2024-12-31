// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type CreateChangeIdentityRequest struct {
	ChangeID                        string                                 `pathParam:"style=simple,explode=false,name=change_id"`
	PostV1ChangesChangeIDIdentities shared.PostV1ChangesChangeIDIdentities `request:"mediaType=application/json"`
}

func (o *CreateChangeIdentityRequest) GetChangeID() string {
	if o == nil {
		return ""
	}
	return o.ChangeID
}

func (o *CreateChangeIdentityRequest) GetPostV1ChangesChangeIDIdentities() shared.PostV1ChangesChangeIDIdentities {
	if o == nil {
		return shared.PostV1ChangesChangeIDIdentities{}
	}
	return o.PostV1ChangesChangeIDIdentities
}

type CreateChangeIdentityResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Create an identity for this change
	ChangeIdentityEntity *shared.ChangeIdentityEntity
	// Bad Request
	ErrorEntity *shared.ErrorEntity
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

func (o *CreateChangeIdentityResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateChangeIdentityResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateChangeIdentityResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateChangeIdentityResponse) GetChangeIdentityEntity() *shared.ChangeIdentityEntity {
	if o == nil {
		return nil
	}
	return o.ChangeIdentityEntity
}

func (o *CreateChangeIdentityResponse) GetErrorEntity() *shared.ErrorEntity {
	if o == nil {
		return nil
	}
	return o.ErrorEntity
}

func (o *CreateChangeIdentityResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *CreateChangeIdentityResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *CreateChangeIdentityResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *CreateChangeIdentityResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *CreateChangeIdentityResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *CreateChangeIdentityResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}