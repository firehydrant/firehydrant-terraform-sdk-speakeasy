// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type UpdateSeverityRequest struct {
	SeveritySlug                  string                               `pathParam:"style=simple,explode=false,name=severity_slug"`
	PatchV1SeveritiesSeveritySlug shared.PatchV1SeveritiesSeveritySlug `request:"mediaType=application/json"`
}

func (o *UpdateSeverityRequest) GetSeveritySlug() string {
	if o == nil {
		return ""
	}
	return o.SeveritySlug
}

func (o *UpdateSeverityRequest) GetPatchV1SeveritiesSeveritySlug() shared.PatchV1SeveritiesSeveritySlug {
	if o == nil {
		return shared.PatchV1SeveritiesSeveritySlug{}
	}
	return o.PatchV1SeveritiesSeveritySlug
}

type UpdateSeverityResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Update a specific severity
	SeverityEntity *shared.SeverityEntity
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

func (o *UpdateSeverityResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateSeverityResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateSeverityResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateSeverityResponse) GetSeverityEntity() *shared.SeverityEntity {
	if o == nil {
		return nil
	}
	return o.SeverityEntity
}

func (o *UpdateSeverityResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *UpdateSeverityResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *UpdateSeverityResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *UpdateSeverityResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *UpdateSeverityResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *UpdateSeverityResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}
