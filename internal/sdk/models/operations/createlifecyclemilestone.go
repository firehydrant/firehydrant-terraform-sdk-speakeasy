// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type CreateLifecycleMilestoneRequestBody struct {
	// The name of the milestone
	Name string `json:"name"`
	// A long-form description of the milestone's purpose
	Description string `json:"description"`
	// A unique identifier for the milestone. If not provided, one will be generated from the name.
	Slug *string `json:"slug,omitempty"`
	// The ID of the phase to which the milestone should belong
	PhaseID string `json:"phase_id"`
	// The position of the milestone within the phase. If not provided, the milestone will be added as the last milestone in the phase.
	Position *int `json:"position,omitempty"`
	// The ID of a later milestone that cannot be started until this milestone has a timestamp populated
	RequiredAtMilestoneID *string `json:"required_at_milestone_id,omitempty"`
}

func (o *CreateLifecycleMilestoneRequestBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateLifecycleMilestoneRequestBody) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *CreateLifecycleMilestoneRequestBody) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *CreateLifecycleMilestoneRequestBody) GetPhaseID() string {
	if o == nil {
		return ""
	}
	return o.PhaseID
}

func (o *CreateLifecycleMilestoneRequestBody) GetPosition() *int {
	if o == nil {
		return nil
	}
	return o.Position
}

func (o *CreateLifecycleMilestoneRequestBody) GetRequiredAtMilestoneID() *string {
	if o == nil {
		return nil
	}
	return o.RequiredAtMilestoneID
}

type CreateLifecycleMilestoneResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Create a new milestone
	LifecyclesPhaseEntityList *shared.LifecyclesPhaseEntityList
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

func (o *CreateLifecycleMilestoneResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateLifecycleMilestoneResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateLifecycleMilestoneResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateLifecycleMilestoneResponse) GetLifecyclesPhaseEntityList() *shared.LifecyclesPhaseEntityList {
	if o == nil {
		return nil
	}
	return o.LifecyclesPhaseEntityList
}

func (o *CreateLifecycleMilestoneResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *CreateLifecycleMilestoneResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *CreateLifecycleMilestoneResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *CreateLifecycleMilestoneResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *CreateLifecycleMilestoneResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *CreateLifecycleMilestoneResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}