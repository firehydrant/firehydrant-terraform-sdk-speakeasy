// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

// TagMatchStrategy - The strategy to match tags. `any` will return alerts that have at least one of the supplied tags, `match_all` will return only alerts that have all of the supplied tags, and `exclude` will only return alerts that have none of the supplied tags. This currently only works for Signals alerts.
type TagMatchStrategy string

const (
	TagMatchStrategyAny      TagMatchStrategy = "any"
	TagMatchStrategyMatchAll TagMatchStrategy = "match_all"
	TagMatchStrategyExclude  TagMatchStrategy = "exclude"
)

func (e TagMatchStrategy) ToPointer() *TagMatchStrategy {
	return &e
}
func (e *TagMatchStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "any":
		fallthrough
	case "match_all":
		fallthrough
	case "exclude":
		*e = TagMatchStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TagMatchStrategy: %v", v)
	}
}

type ListAlertsRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
	// A text query for alerts
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// A comma separated list of user IDs. This currently only works for Signals alerts.
	Users *string `queryParam:"style=form,explode=true,name=users"`
	// A comma separated list of team IDs. This currently only works for Signals alerts.
	Teams *string `queryParam:"style=form,explode=true,name=teams"`
	// A comma separated list of signals rule IDs. This currently only works for Signals alerts.
	SignalRules *string `queryParam:"style=form,explode=true,name=signal_rules"`
	// A comma separated list of environment IDs. This currently only works for Signals alerts.
	Environments *string `queryParam:"style=form,explode=true,name=environments"`
	// A comma separated list of functionality IDs. This currently only works for Signals alerts.
	Functionalities *string `queryParam:"style=form,explode=true,name=functionalities"`
	// A comma separated list of service IDs. This currently only works for Signals alerts.
	Services *string `queryParam:"style=form,explode=true,name=services"`
	// A comma separated list of tags. This currently only works for Signals alerts.
	Tags *string `queryParam:"style=form,explode=true,name=tags"`
	// The strategy to match tags. `any` will return alerts that have at least one of the supplied tags, `match_all` will return only alerts that have all of the supplied tags, and `exclude` will only return alerts that have none of the supplied tags. This currently only works for Signals alerts.
	TagMatchStrategy *TagMatchStrategy `queryParam:"style=form,explode=true,name=tag_match_strategy"`
	// A comma separated list of statuses to filter by. Valid statuses are: opened, acknowledged, resolved, ignored, or expired
	Statuses *string `queryParam:"style=form,explode=true,name=statuses"`
}

func (o *ListAlertsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListAlertsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *ListAlertsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *ListAlertsRequest) GetUsers() *string {
	if o == nil {
		return nil
	}
	return o.Users
}

func (o *ListAlertsRequest) GetTeams() *string {
	if o == nil {
		return nil
	}
	return o.Teams
}

func (o *ListAlertsRequest) GetSignalRules() *string {
	if o == nil {
		return nil
	}
	return o.SignalRules
}

func (o *ListAlertsRequest) GetEnvironments() *string {
	if o == nil {
		return nil
	}
	return o.Environments
}

func (o *ListAlertsRequest) GetFunctionalities() *string {
	if o == nil {
		return nil
	}
	return o.Functionalities
}

func (o *ListAlertsRequest) GetServices() *string {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *ListAlertsRequest) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *ListAlertsRequest) GetTagMatchStrategy() *TagMatchStrategy {
	if o == nil {
		return nil
	}
	return o.TagMatchStrategy
}

func (o *ListAlertsRequest) GetStatuses() *string {
	if o == nil {
		return nil
	}
	return o.Statuses
}

type ListAlertsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Retrieve all alerts from third parties
	AlertsAlertEntityPaginated *shared.AlertsAlertEntityPaginated
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

func (o *ListAlertsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAlertsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAlertsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListAlertsResponse) GetAlertsAlertEntityPaginated() *shared.AlertsAlertEntityPaginated {
	if o == nil {
		return nil
	}
	return o.AlertsAlertEntityPaginated
}

func (o *ListAlertsResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *ListAlertsResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *ListAlertsResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *ListAlertsResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *ListAlertsResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *ListAlertsResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}