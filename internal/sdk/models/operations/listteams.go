// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type ListTeamsRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
	// A query to search teams by their name or description
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// A query to search teams by their name
	Name *string `queryParam:"style=form,explode=true,name=name"`
	// A comma separated list of service IDs
	Services *string `queryParam:"style=form,explode=true,name=services"`
	// Filter by teams that have or do not have members with a default incident role asssigned. Value may be 'present', 'blank', or the ID of an incident role.
	DefaultIncidentRole *string `queryParam:"style=form,explode=true,name=default_incident_role"`
	// Boolean to determine whether to return a slimified version of the teams object
	Lite *bool `queryParam:"style=form,explode=true,name=lite"`
}

func (o *ListTeamsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListTeamsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *ListTeamsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListTeamsRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ListTeamsRequest) GetServices() *string {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *ListTeamsRequest) GetDefaultIncidentRole() *string {
	if o == nil {
		return nil
	}
	return o.DefaultIncidentRole
}

func (o *ListTeamsRequest) GetLite() *bool {
	if o == nil {
		return nil
	}
	return o.Lite
}

type ListTeamsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List all of the teams in the organization
	TeamEntityPaginated *shared.TeamEntityPaginated
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

func (o *ListTeamsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListTeamsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListTeamsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListTeamsResponse) GetTeamEntityPaginated() *shared.TeamEntityPaginated {
	if o == nil {
		return nil
	}
	return o.TeamEntityPaginated
}

func (o *ListTeamsResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *ListTeamsResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *ListTeamsResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *ListTeamsResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *ListTeamsResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *ListTeamsResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}