// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type CreateIncidentTaskListRequest struct {
	IncidentID                         string                                    `pathParam:"style=simple,explode=false,name=incident_id"`
	PostV1IncidentsIncidentIDTaskLists shared.PostV1IncidentsIncidentIDTaskLists `request:"mediaType=application/json"`
}

func (o *CreateIncidentTaskListRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *CreateIncidentTaskListRequest) GetPostV1IncidentsIncidentIDTaskLists() shared.PostV1IncidentsIncidentIDTaskLists {
	if o == nil {
		return shared.PostV1IncidentsIncidentIDTaskLists{}
	}
	return o.PostV1IncidentsIncidentIDTaskLists
}

type CreateIncidentTaskListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Add all tasks from list to incident
	TaskEntity *shared.TaskEntity
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

func (o *CreateIncidentTaskListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateIncidentTaskListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateIncidentTaskListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateIncidentTaskListResponse) GetTaskEntity() *shared.TaskEntity {
	if o == nil {
		return nil
	}
	return o.TaskEntity
}

func (o *CreateIncidentTaskListResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *CreateIncidentTaskListResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *CreateIncidentTaskListResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *CreateIncidentTaskListResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *CreateIncidentTaskListResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *CreateIncidentTaskListResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}
