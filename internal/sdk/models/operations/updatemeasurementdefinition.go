// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type UpdateMeasurementDefinitionRequestBody struct {
	Name                *string `json:"name,omitempty"`
	Slug                *string `json:"slug,omitempty"`
	Description         *string `json:"description,omitempty"`
	StartsAtMilestoneID *string `json:"starts_at_milestone_id,omitempty"`
	EndsAtMilestoneID   *string `json:"ends_at_milestone_id,omitempty"`
}

func (o *UpdateMeasurementDefinitionRequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateMeasurementDefinitionRequestBody) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *UpdateMeasurementDefinitionRequestBody) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateMeasurementDefinitionRequestBody) GetStartsAtMilestoneID() *string {
	if o == nil {
		return nil
	}
	return o.StartsAtMilestoneID
}

func (o *UpdateMeasurementDefinitionRequestBody) GetEndsAtMilestoneID() *string {
	if o == nil {
		return nil
	}
	return o.EndsAtMilestoneID
}

type UpdateMeasurementDefinitionRequest struct {
	MeasurementDefinitionID string                                  `pathParam:"style=simple,explode=false,name=measurement_definition_id"`
	RequestBody             *UpdateMeasurementDefinitionRequestBody `request:"mediaType=application/json"`
}

func (o *UpdateMeasurementDefinitionRequest) GetMeasurementDefinitionID() string {
	if o == nil {
		return ""
	}
	return o.MeasurementDefinitionID
}

func (o *UpdateMeasurementDefinitionRequest) GetRequestBody() *UpdateMeasurementDefinitionRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type UpdateMeasurementDefinitionResponse struct {
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

func (o *UpdateMeasurementDefinitionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateMeasurementDefinitionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateMeasurementDefinitionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateMeasurementDefinitionResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *UpdateMeasurementDefinitionResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *UpdateMeasurementDefinitionResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *UpdateMeasurementDefinitionResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *UpdateMeasurementDefinitionResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *UpdateMeasurementDefinitionResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}