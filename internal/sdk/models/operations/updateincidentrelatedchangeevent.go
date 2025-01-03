// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type UpdateIncidentRelatedChangeEventRequest struct {
	RelatedChangeEventID                                              string                                                                   `pathParam:"style=simple,explode=false,name=related_change_event_id"`
	IncidentID                                                        string                                                                   `pathParam:"style=simple,explode=false,name=incident_id"`
	PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID shared.PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID `request:"mediaType=application/json"`
}

func (o *UpdateIncidentRelatedChangeEventRequest) GetRelatedChangeEventID() string {
	if o == nil {
		return ""
	}
	return o.RelatedChangeEventID
}

func (o *UpdateIncidentRelatedChangeEventRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *UpdateIncidentRelatedChangeEventRequest) GetPatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID() shared.PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID {
	if o == nil {
		return shared.PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID{}
	}
	return o.PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID
}

type UpdateIncidentRelatedChangeEventResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Update a change attached to an incident
	IncidentsRelatedChangeEventEntity *shared.IncidentsRelatedChangeEventEntity
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

func (o *UpdateIncidentRelatedChangeEventResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateIncidentRelatedChangeEventResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateIncidentRelatedChangeEventResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateIncidentRelatedChangeEventResponse) GetIncidentsRelatedChangeEventEntity() *shared.IncidentsRelatedChangeEventEntity {
	if o == nil {
		return nil
	}
	return o.IncidentsRelatedChangeEventEntity
}

func (o *UpdateIncidentRelatedChangeEventResponse) GetErrorEntity() *shared.ErrorEntity {
	if o == nil {
		return nil
	}
	return o.ErrorEntity
}

func (o *UpdateIncidentRelatedChangeEventResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *UpdateIncidentRelatedChangeEventResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *UpdateIncidentRelatedChangeEventResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *UpdateIncidentRelatedChangeEventResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *UpdateIncidentRelatedChangeEventResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *UpdateIncidentRelatedChangeEventResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}
