// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/internal/utils"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/types"
	"net/http"
)

type GetMeanTimeReportRequest struct {
	// A comma separated list of environment IDs
	Environments *string `queryParam:"style=form,explode=true,name=environments"`
	// A comma separated list of team IDs
	Teams *string `queryParam:"style=form,explode=true,name=teams"`
	// A comma separated list of service IDs
	Services *string `queryParam:"style=form,explode=true,name=services"`
	// Incident status
	Status *string `queryParam:"style=form,explode=true,name=status"`
	// The start date to return incidents from
	StartDate *types.Date `queryParam:"style=form,explode=true,name=start_date"`
	// The end date to return incidents from
	EndDate *types.Date `queryParam:"style=form,explode=true,name=end_date"`
	// A text query for an incident that searches on name, summary, and desciption
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// The id of a previously saved search.
	SavedSearchID *string `queryParam:"style=form,explode=true,name=saved_search_id"`
	// A comma separated list of priorities
	Priorities *string `queryParam:"style=form,explode=true,name=priorities"`
	// Flag for including incidents where priority has not been set
	PriorityNotSet *bool `queryParam:"style=form,explode=true,name=priority_not_set"`
	// A comma separated list of severities
	Severities *string `queryParam:"style=form,explode=true,name=severities"`
	// Flag for including incidents where severity has not been set
	SeverityNotSet *bool `queryParam:"style=form,explode=true,name=severity_not_set"`
	// A comma separated list of current milestones
	CurrentMilestones *string `queryParam:"style=form,explode=true,name=current_milestones"`
}

func (g GetMeanTimeReportRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetMeanTimeReportRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetMeanTimeReportRequest) GetEnvironments() *string {
	if o == nil {
		return nil
	}
	return o.Environments
}

func (o *GetMeanTimeReportRequest) GetTeams() *string {
	if o == nil {
		return nil
	}
	return o.Teams
}

func (o *GetMeanTimeReportRequest) GetServices() *string {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *GetMeanTimeReportRequest) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *GetMeanTimeReportRequest) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *GetMeanTimeReportRequest) GetEndDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *GetMeanTimeReportRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *GetMeanTimeReportRequest) GetSavedSearchID() *string {
	if o == nil {
		return nil
	}
	return o.SavedSearchID
}

func (o *GetMeanTimeReportRequest) GetPriorities() *string {
	if o == nil {
		return nil
	}
	return o.Priorities
}

func (o *GetMeanTimeReportRequest) GetPriorityNotSet() *bool {
	if o == nil {
		return nil
	}
	return o.PriorityNotSet
}

func (o *GetMeanTimeReportRequest) GetSeverities() *string {
	if o == nil {
		return nil
	}
	return o.Severities
}

func (o *GetMeanTimeReportRequest) GetSeverityNotSet() *bool {
	if o == nil {
		return nil
	}
	return o.SeverityNotSet
}

func (o *GetMeanTimeReportRequest) GetCurrentMilestones() *string {
	if o == nil {
		return nil
	}
	return o.CurrentMilestones
}

type GetMeanTimeReportResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Returns a report with time bucketed analytics data
	ReportEntity *shared.ReportEntity
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

func (o *GetMeanTimeReportResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMeanTimeReportResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMeanTimeReportResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMeanTimeReportResponse) GetReportEntity() *shared.ReportEntity {
	if o == nil {
		return nil
	}
	return o.ReportEntity
}

func (o *GetMeanTimeReportResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *GetMeanTimeReportResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *GetMeanTimeReportResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *GetMeanTimeReportResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *GetMeanTimeReportResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *GetMeanTimeReportResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}