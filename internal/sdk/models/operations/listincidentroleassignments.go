// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

// Status - Filter on status of the role assignment
type Status string

const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
)

func (e Status) ToPointer() *Status {
	return &e
}
func (e *Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "active":
		fallthrough
	case "inactive":
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

type ListIncidentRoleAssignmentsRequest struct {
	IncidentID string `pathParam:"style=simple,explode=false,name=incident_id"`
	// Filter on status of the role assignment
	Status *Status `queryParam:"style=form,explode=true,name=status"`
}

func (o *ListIncidentRoleAssignmentsRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *ListIncidentRoleAssignmentsRequest) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}

type ListIncidentRoleAssignmentsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Retrieve a list of all of the current role assignments for the incident
	IncidentsRoleAssignmentEntityPaginated *shared.IncidentsRoleAssignmentEntityPaginated
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

func (o *ListIncidentRoleAssignmentsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListIncidentRoleAssignmentsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListIncidentRoleAssignmentsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListIncidentRoleAssignmentsResponse) GetIncidentsRoleAssignmentEntityPaginated() *shared.IncidentsRoleAssignmentEntityPaginated {
	if o == nil {
		return nil
	}
	return o.IncidentsRoleAssignmentEntityPaginated
}

func (o *ListIncidentRoleAssignmentsResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *ListIncidentRoleAssignmentsResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *ListIncidentRoleAssignmentsResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *ListIncidentRoleAssignmentsResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *ListIncidentRoleAssignmentsResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *ListIncidentRoleAssignmentsResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}