// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type ListIncidentAttachmentsRequest struct {
	IncidentID     string  `pathParam:"style=simple,explode=false,name=incident_id"`
	AttachableType *string `queryParam:"style=form,explode=true,name=attachable_type"`
	Page           *int    `queryParam:"style=form,explode=true,name=page"`
	PerPage        *int    `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *ListIncidentAttachmentsRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *ListIncidentAttachmentsRequest) GetAttachableType() *string {
	if o == nil {
		return nil
	}
	return o.AttachableType
}

func (o *ListIncidentAttachmentsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListIncidentAttachmentsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type ListIncidentAttachmentsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
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

func (o *ListIncidentAttachmentsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListIncidentAttachmentsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListIncidentAttachmentsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListIncidentAttachmentsResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *ListIncidentAttachmentsResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *ListIncidentAttachmentsResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *ListIncidentAttachmentsResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *ListIncidentAttachmentsResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *ListIncidentAttachmentsResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}