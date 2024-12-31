// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type CreateIncidentRelatedChangeRequest struct {
	IncidentID                                   string                                              `pathParam:"style=simple,explode=false,name=incident_id"`
	PostV1IncidentsIncidentIDRelatedChangeEvents shared.PostV1IncidentsIncidentIDRelatedChangeEvents `request:"mediaType=application/json"`
}

func (o *CreateIncidentRelatedChangeRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *CreateIncidentRelatedChangeRequest) GetPostV1IncidentsIncidentIDRelatedChangeEvents() shared.PostV1IncidentsIncidentIDRelatedChangeEvents {
	if o == nil {
		return shared.PostV1IncidentsIncidentIDRelatedChangeEvents{}
	}
	return o.PostV1IncidentsIncidentIDRelatedChangeEvents
}

type CreateIncidentRelatedChangeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Add a related change to an incident. Changes added to an incident can be causes, fixes, or suspects. To remove a change from an incident, the type field should be set to dismissed.
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

func (o *CreateIncidentRelatedChangeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateIncidentRelatedChangeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateIncidentRelatedChangeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateIncidentRelatedChangeResponse) GetIncidentsRelatedChangeEventEntity() *shared.IncidentsRelatedChangeEventEntity {
	if o == nil {
		return nil
	}
	return o.IncidentsRelatedChangeEventEntity
}

func (o *CreateIncidentRelatedChangeResponse) GetErrorEntity() *shared.ErrorEntity {
	if o == nil {
		return nil
	}
	return o.ErrorEntity
}

func (o *CreateIncidentRelatedChangeResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *CreateIncidentRelatedChangeResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *CreateIncidentRelatedChangeResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *CreateIncidentRelatedChangeResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *CreateIncidentRelatedChangeResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *CreateIncidentRelatedChangeResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}