// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type UpdateIncidentEventVotesRequest struct {
	IncidentID                                   string                                              `pathParam:"style=simple,explode=false,name=incident_id"`
	EventID                                      string                                              `pathParam:"style=simple,explode=false,name=event_id"`
	PatchV1IncidentsIncidentIDEventsEventIDVotes shared.PatchV1IncidentsIncidentIDEventsEventIDVotes `request:"mediaType=application/json"`
}

func (o *UpdateIncidentEventVotesRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *UpdateIncidentEventVotesRequest) GetEventID() string {
	if o == nil {
		return ""
	}
	return o.EventID
}

func (o *UpdateIncidentEventVotesRequest) GetPatchV1IncidentsIncidentIDEventsEventIDVotes() shared.PatchV1IncidentsIncidentIDEventsEventIDVotes {
	if o == nil {
		return shared.PatchV1IncidentsIncidentIDEventsEventIDVotes{}
	}
	return o.PatchV1IncidentsIncidentIDEventsEventIDVotes
}

type UpdateIncidentEventVotesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Allows for upvoting or downvoting an event
	VotesEntity *shared.VotesEntity
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

func (o *UpdateIncidentEventVotesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateIncidentEventVotesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateIncidentEventVotesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateIncidentEventVotesResponse) GetVotesEntity() *shared.VotesEntity {
	if o == nil {
		return nil
	}
	return o.VotesEntity
}

func (o *UpdateIncidentEventVotesResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *UpdateIncidentEventVotesResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *UpdateIncidentEventVotesResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *UpdateIncidentEventVotesResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *UpdateIncidentEventVotesResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *UpdateIncidentEventVotesResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}