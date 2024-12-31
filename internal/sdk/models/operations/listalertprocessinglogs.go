// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

// OfLevel - Returns logs of all levels equal to or above the provided level
type OfLevel string

const (
	OfLevelUnknown OfLevel = "unknown"
	OfLevelDebug   OfLevel = "debug"
	OfLevelInfo    OfLevel = "info"
	OfLevelWarn    OfLevel = "warn"
	OfLevelError   OfLevel = "error"
	OfLevelFatal   OfLevel = "fatal"
)

func (e OfLevel) ToPointer() *OfLevel {
	return &e
}
func (e *OfLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "unknown":
		fallthrough
	case "debug":
		fallthrough
	case "info":
		fallthrough
	case "warn":
		fallthrough
	case "error":
		fallthrough
	case "fatal":
		*e = OfLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OfLevel: %v", v)
	}
}

// ExactLevel - Returns log entries of all levels equal to the provided level
type ExactLevel string

const (
	ExactLevelUnknown ExactLevel = "unknown"
	ExactLevelDebug   ExactLevel = "debug"
	ExactLevelInfo    ExactLevel = "info"
	ExactLevelWarn    ExactLevel = "warn"
	ExactLevelError   ExactLevel = "error"
	ExactLevelFatal   ExactLevel = "fatal"
)

func (e ExactLevel) ToPointer() *ExactLevel {
	return &e
}
func (e *ExactLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "unknown":
		fallthrough
	case "debug":
		fallthrough
	case "info":
		fallthrough
	case "warn":
		fallthrough
	case "error":
		fallthrough
	case "fatal":
		*e = ExactLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ExactLevel: %v", v)
	}
}

type ListAlertProcessingLogsRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
	// Scopes returned log entries to a specific integration ID
	IntegrationSlug *string `queryParam:"style=form,explode=true,name=integration_slug"`
	// Scopes returned log entries to a specific connection ID
	ConnectionID *string `queryParam:"style=form,explode=true,name=connection_id"`
	// Returns logs of all levels equal to or above the provided level
	OfLevel *OfLevel `queryParam:"style=form,explode=true,name=of_level"`
	// Returns log entries of all levels equal to the provided level
	ExactLevel *ExactLevel `queryParam:"style=form,explode=true,name=exact_level"`
}

func (o *ListAlertProcessingLogsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListAlertProcessingLogsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *ListAlertProcessingLogsRequest) GetIntegrationSlug() *string {
	if o == nil {
		return nil
	}
	return o.IntegrationSlug
}

func (o *ListAlertProcessingLogsRequest) GetConnectionID() *string {
	if o == nil {
		return nil
	}
	return o.ConnectionID
}

func (o *ListAlertProcessingLogsRequest) GetOfLevel() *OfLevel {
	if o == nil {
		return nil
	}
	return o.OfLevel
}

func (o *ListAlertProcessingLogsRequest) GetExactLevel() *ExactLevel {
	if o == nil {
		return nil
	}
	return o.ExactLevel
}

type ListAlertProcessingLogsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Processing Log Entries for a specific alert
	AlertsProcessingLogEntryEntityPaginated *shared.AlertsProcessingLogEntryEntityPaginated
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

func (o *ListAlertProcessingLogsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAlertProcessingLogsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAlertProcessingLogsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListAlertProcessingLogsResponse) GetAlertsProcessingLogEntryEntityPaginated() *shared.AlertsProcessingLogEntryEntityPaginated {
	if o == nil {
		return nil
	}
	return o.AlertsProcessingLogEntryEntityPaginated
}

func (o *ListAlertProcessingLogsResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *ListAlertProcessingLogsResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *ListAlertProcessingLogsResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *ListAlertProcessingLogsResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *ListAlertProcessingLogsResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *ListAlertProcessingLogsResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}