// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/internal/utils"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/types"
	"net/http"
)

type InfraType string

const (
	InfraTypeEnvironments    InfraType = "environments"
	InfraTypeFunctionalities InfraType = "functionalities"
	InfraTypeServices        InfraType = "services"
	InfraTypeCustomers       InfraType = "customers"
)

func (e InfraType) ToPointer() *InfraType {
	return &e
}
func (e *InfraType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "environments":
		fallthrough
	case "functionalities":
		fallthrough
	case "services":
		fallthrough
	case "customers":
		*e = InfraType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InfraType: %v", v)
	}
}

type ListInfrastructureMetricsRequest struct {
	InfraType InfraType `pathParam:"style=simple,explode=false,name=infra_type"`
	// The start date to return metrics from; defaults to 30 days ago
	StartDate *types.Date `queryParam:"style=form,explode=true,name=start_date"`
	// The end date to return metrics from, defaults to today
	EndDate *types.Date `queryParam:"style=form,explode=true,name=end_date"`
}

func (l ListInfrastructureMetricsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListInfrastructureMetricsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListInfrastructureMetricsRequest) GetInfraType() InfraType {
	if o == nil {
		return InfraType("")
	}
	return o.InfraType
}

func (o *ListInfrastructureMetricsRequest) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *ListInfrastructureMetricsRequest) GetEndDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.EndDate
}

type ListInfrastructureMetricsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Returns metrics for all components of a given type
	MetricsInfrastructureListEntity *shared.MetricsInfrastructureListEntity
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

func (o *ListInfrastructureMetricsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListInfrastructureMetricsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListInfrastructureMetricsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListInfrastructureMetricsResponse) GetMetricsInfrastructureListEntity() *shared.MetricsInfrastructureListEntity {
	if o == nil {
		return nil
	}
	return o.MetricsInfrastructureListEntity
}

func (o *ListInfrastructureMetricsResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *ListInfrastructureMetricsResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *ListInfrastructureMetricsResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *ListInfrastructureMetricsResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *ListInfrastructureMetricsResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *ListInfrastructureMetricsResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}
