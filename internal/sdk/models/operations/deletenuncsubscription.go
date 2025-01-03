// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type DeleteNuncSubscriptionRequest struct {
	UnsubscribeToken string `pathParam:"style=simple,explode=false,name=unsubscribe_token"`
}

func (o *DeleteNuncSubscriptionRequest) GetUnsubscribeToken() string {
	if o == nil {
		return ""
	}
	return o.UnsubscribeToken
}

type DeleteNuncSubscriptionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unsubscribe from status page updates
	NuncNuncSubscription *shared.NuncNuncSubscription
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

func (o *DeleteNuncSubscriptionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteNuncSubscriptionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteNuncSubscriptionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteNuncSubscriptionResponse) GetNuncNuncSubscription() *shared.NuncNuncSubscription {
	if o == nil {
		return nil
	}
	return o.NuncNuncSubscription
}

func (o *DeleteNuncSubscriptionResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *DeleteNuncSubscriptionResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *DeleteNuncSubscriptionResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *DeleteNuncSubscriptionResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *DeleteNuncSubscriptionResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *DeleteNuncSubscriptionResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}
