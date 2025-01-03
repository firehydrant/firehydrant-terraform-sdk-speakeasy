// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type ListIncidentEventsRequest struct {
	IncidentID string `pathParam:"style=simple,explode=false,name=incident_id"`
	// A comma separated list of types of events to filter by
	Types   *string `queryParam:"style=form,explode=true,name=types"`
	Page    *int    `queryParam:"style=form,explode=true,name=page"`
	PerPage *int    `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *ListIncidentEventsRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *ListIncidentEventsRequest) GetTypes() *string {
	if o == nil {
		return nil
	}
	return o.Types
}

func (o *ListIncidentEventsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListIncidentEventsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type ListIncidentEventsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List all events for an incident. An event is a timeline entry. This can be filtered with params to retrieve events of a certain type.
	IncidentEventEntityPaginated *shared.IncidentEventEntityPaginated
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

func (o *ListIncidentEventsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListIncidentEventsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListIncidentEventsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListIncidentEventsResponse) GetIncidentEventEntityPaginated() *shared.IncidentEventEntityPaginated {
	if o == nil {
		return nil
	}
	return o.IncidentEventEntityPaginated
}

func (o *ListIncidentEventsResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *ListIncidentEventsResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *ListIncidentEventsResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *ListIncidentEventsResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *ListIncidentEventsResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *ListIncidentEventsResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}
