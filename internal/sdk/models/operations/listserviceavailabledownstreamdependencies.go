// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type ListServiceAvailableDownstreamDependenciesRequest struct {
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
}

func (o *ListServiceAvailableDownstreamDependenciesRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

type ListServiceAvailableDownstreamDependenciesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Retrieves all services that are available to be downstream dependencies
	ServiceEntityLite *shared.ServiceEntityLite
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

func (o *ListServiceAvailableDownstreamDependenciesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListServiceAvailableDownstreamDependenciesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListServiceAvailableDownstreamDependenciesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListServiceAvailableDownstreamDependenciesResponse) GetServiceEntityLite() *shared.ServiceEntityLite {
	if o == nil {
		return nil
	}
	return o.ServiceEntityLite
}

func (o *ListServiceAvailableDownstreamDependenciesResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *ListServiceAvailableDownstreamDependenciesResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *ListServiceAvailableDownstreamDependenciesResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *ListServiceAvailableDownstreamDependenciesResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *ListServiceAvailableDownstreamDependenciesResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *ListServiceAvailableDownstreamDependenciesResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}
