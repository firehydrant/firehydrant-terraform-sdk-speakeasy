// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type UpdateRetrospectiveReasonRequest struct {
	ReportID                                         string                                                  `pathParam:"style=simple,explode=false,name=report_id"`
	ReasonID                                         string                                                  `pathParam:"style=simple,explode=false,name=reason_id"`
	PatchV1PostMortemsReportsReportIDReasonsReasonID shared.PatchV1PostMortemsReportsReportIDReasonsReasonID `request:"mediaType=application/json"`
}

func (o *UpdateRetrospectiveReasonRequest) GetReportID() string {
	if o == nil {
		return ""
	}
	return o.ReportID
}

func (o *UpdateRetrospectiveReasonRequest) GetReasonID() string {
	if o == nil {
		return ""
	}
	return o.ReasonID
}

func (o *UpdateRetrospectiveReasonRequest) GetPatchV1PostMortemsReportsReportIDReasonsReasonID() shared.PatchV1PostMortemsReportsReportIDReasonsReasonID {
	if o == nil {
		return shared.PatchV1PostMortemsReportsReportIDReasonsReasonID{}
	}
	return o.PatchV1PostMortemsReportsReportIDReasonsReasonID
}

type UpdateRetrospectiveReasonResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Update a contributing factor
	PostMortemsReasonEntity *shared.PostMortemsReasonEntity
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

func (o *UpdateRetrospectiveReasonResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateRetrospectiveReasonResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateRetrospectiveReasonResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateRetrospectiveReasonResponse) GetPostMortemsReasonEntity() *shared.PostMortemsReasonEntity {
	if o == nil {
		return nil
	}
	return o.PostMortemsReasonEntity
}

func (o *UpdateRetrospectiveReasonResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *UpdateRetrospectiveReasonResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *UpdateRetrospectiveReasonResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *UpdateRetrospectiveReasonResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *UpdateRetrospectiveReasonResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *UpdateRetrospectiveReasonResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}
