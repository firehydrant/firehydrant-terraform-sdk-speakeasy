// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type CreateTeamOnCallScheduleRequest struct {
	TeamID                           string                                  `pathParam:"style=simple,explode=false,name=team_id"`
	PostV1TeamsTeamIDOnCallSchedules shared.PostV1TeamsTeamIDOnCallSchedules `request:"mediaType=application/json"`
}

func (o *CreateTeamOnCallScheduleRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *CreateTeamOnCallScheduleRequest) GetPostV1TeamsTeamIDOnCallSchedules() shared.PostV1TeamsTeamIDOnCallSchedules {
	if o == nil {
		return shared.PostV1TeamsTeamIDOnCallSchedules{}
	}
	return o.PostV1TeamsTeamIDOnCallSchedules
}

type CreateTeamOnCallScheduleResponse struct {
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

func (o *CreateTeamOnCallScheduleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateTeamOnCallScheduleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTeamOnCallScheduleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateTeamOnCallScheduleResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *CreateTeamOnCallScheduleResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *CreateTeamOnCallScheduleResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *CreateTeamOnCallScheduleResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *CreateTeamOnCallScheduleResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *CreateTeamOnCallScheduleResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}