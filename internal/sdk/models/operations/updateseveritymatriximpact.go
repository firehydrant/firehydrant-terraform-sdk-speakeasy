// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type UpdateSeverityMatrixImpactRequest struct {
	ImpactID                             string                                      `pathParam:"style=simple,explode=false,name=impact_id"`
	PatchV1SeverityMatrixImpactsImpactID shared.PatchV1SeverityMatrixImpactsImpactID `request:"mediaType=application/json"`
}

func (o *UpdateSeverityMatrixImpactRequest) GetImpactID() string {
	if o == nil {
		return ""
	}
	return o.ImpactID
}

func (o *UpdateSeverityMatrixImpactRequest) GetPatchV1SeverityMatrixImpactsImpactID() shared.PatchV1SeverityMatrixImpactsImpactID {
	if o == nil {
		return shared.PatchV1SeverityMatrixImpactsImpactID{}
	}
	return o.PatchV1SeverityMatrixImpactsImpactID
}

type UpdateSeverityMatrixImpactResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Update a specific impact
	SeverityMatrixImpactEntity *shared.SeverityMatrixImpactEntity
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

func (o *UpdateSeverityMatrixImpactResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateSeverityMatrixImpactResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateSeverityMatrixImpactResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateSeverityMatrixImpactResponse) GetSeverityMatrixImpactEntity() *shared.SeverityMatrixImpactEntity {
	if o == nil {
		return nil
	}
	return o.SeverityMatrixImpactEntity
}

func (o *UpdateSeverityMatrixImpactResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *UpdateSeverityMatrixImpactResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *UpdateSeverityMatrixImpactResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *UpdateSeverityMatrixImpactResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *UpdateSeverityMatrixImpactResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *UpdateSeverityMatrixImpactResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}