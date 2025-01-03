// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/internal/utils"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/types"
	"net/http"
	"time"
)

// GetTicketFunnelMetricsQueryParamTagMatchStrategy - A matching strategy for the tags provided
type GetTicketFunnelMetricsQueryParamTagMatchStrategy string

const (
	GetTicketFunnelMetricsQueryParamTagMatchStrategyAny      GetTicketFunnelMetricsQueryParamTagMatchStrategy = "any"
	GetTicketFunnelMetricsQueryParamTagMatchStrategyMatchAll GetTicketFunnelMetricsQueryParamTagMatchStrategy = "match_all"
	GetTicketFunnelMetricsQueryParamTagMatchStrategyExclude  GetTicketFunnelMetricsQueryParamTagMatchStrategy = "exclude"
)

func (e GetTicketFunnelMetricsQueryParamTagMatchStrategy) ToPointer() *GetTicketFunnelMetricsQueryParamTagMatchStrategy {
	return &e
}
func (e *GetTicketFunnelMetricsQueryParamTagMatchStrategy) UnmarshalJSON(data []byte) error {
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
		*e = GetTicketFunnelMetricsQueryParamTagMatchStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetTicketFunnelMetricsQueryParamTagMatchStrategy: %v", v)
	}
}

type GetTicketFunnelMetricsGroupBy string

const (
	GetTicketFunnelMetricsGroupByStartedDay   GetTicketFunnelMetricsGroupBy = "started_day"
	GetTicketFunnelMetricsGroupByStartedWeek  GetTicketFunnelMetricsGroupBy = "started_week"
	GetTicketFunnelMetricsGroupByStartedMonth GetTicketFunnelMetricsGroupBy = "started_month"
	GetTicketFunnelMetricsGroupByAllTime      GetTicketFunnelMetricsGroupBy = "all_time"
)

func (e GetTicketFunnelMetricsGroupBy) ToPointer() *GetTicketFunnelMetricsGroupBy {
	return &e
}
func (e *GetTicketFunnelMetricsGroupBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "started_day":
		fallthrough
	case "started_week":
		fallthrough
	case "started_month":
		fallthrough
	case "all_time":
		*e = GetTicketFunnelMetricsGroupBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetTicketFunnelMetricsGroupBy: %v", v)
	}
}

type GetTicketFunnelMetricsRequestBody struct {
	GroupBy []GetTicketFunnelMetricsGroupBy `multipartForm:"name=group_by"`
}

func (o *GetTicketFunnelMetricsRequestBody) GetGroupBy() []GetTicketFunnelMetricsGroupBy {
	if o == nil {
		return nil
	}
	return o.GroupBy
}

type GetTicketFunnelMetricsRequest struct {
	// A JSON string that defines 'logic' and 'user_data'
	Conditions *string `queryParam:"style=form,explode=true,name=conditions"`
	// A comma separated list of environment IDs or 'is_empty' to filter for incidents with no impacted environments
	Environments *string `queryParam:"style=form,explode=true,name=environments"`
	// A comma separated list of service IDs or 'is_empty' to filter for incidents with no impacted services
	Services *string `queryParam:"style=form,explode=true,name=services"`
	// A comma separated list of functionality IDs or 'is_empty' to filter for incidents with no impacted functionalities
	Functionalities *string `queryParam:"style=form,explode=true,name=functionalities"`
	// A comma separated list of infrastructure IDs. Returns incidents that do not have the following infrastructure ids associated with them.
	ExcludedInfrastructureIds *string `queryParam:"style=form,explode=true,name=excluded_infrastructure_ids"`
	// A comma separated list of team IDs
	Teams *string `queryParam:"style=form,explode=true,name=teams"`
	// A comma separated list of IDs for assigned teams or 'is_empty' to filter for incidents with no active team assignments
	AssignedTeams *string `queryParam:"style=form,explode=true,name=assigned_teams"`
	// Incident status
	Status *string `queryParam:"style=form,explode=true,name=status"`
	// Filters for incidents that started on or after this date
	StartDate *types.Date `queryParam:"style=form,explode=true,name=start_date"`
	// Filters for incidents that started on or before this date
	EndDate *types.Date `queryParam:"style=form,explode=true,name=end_date"`
	// Filters for incidents that were resolved at or after this time. Combine this with the `current_milestones` parameter if you wish to omit incidents that were re-opened and are still active.
	ResolvedAtOrAfter *time.Time `queryParam:"style=form,explode=true,name=resolved_at_or_after"`
	// Filters for incidents that were resolved at or before this time. Combine this with the `current_milestones` parameter if you wish to omit incidents that were re-opened and are still active.
	ResolvedAtOrBefore *time.Time `queryParam:"style=form,explode=true,name=resolved_at_or_before"`
	// Filters for incidents that were created at or after this time
	CreatedAtOrAfter *time.Time `queryParam:"style=form,explode=true,name=created_at_or_after"`
	// Filters for incidents that were created at or before this time
	CreatedAtOrBefore *time.Time `queryParam:"style=form,explode=true,name=created_at_or_before"`
	// A text query for an incident that searches on name, summary, and desciption
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// A query to search incidents by their name
	Name *string `queryParam:"style=form,explode=true,name=name"`
	// The id of a previously saved search.
	SavedSearchID *string `queryParam:"style=form,explode=true,name=saved_search_id"`
	// A text value of priority
	Priorities *string `queryParam:"style=form,explode=true,name=priorities"`
	// Flag for including incidents where priority has not been set
	PriorityNotSet *bool `queryParam:"style=form,explode=true,name=priority_not_set"`
	// A text value of severity
	Severities *string `queryParam:"style=form,explode=true,name=severities"`
	// Flag for including incidents where severity has not been set
	SeverityNotSet *bool `queryParam:"style=form,explode=true,name=severity_not_set"`
	// A comma separated list of current milestones
	CurrentMilestones *string `queryParam:"style=form,explode=true,name=current_milestones"`
	// A comma separated list of tags
	Tags *string `queryParam:"style=form,explode=true,name=tags"`
	// A matching strategy for the tags provided
	TagMatchStrategy *GetTicketFunnelMetricsQueryParamTagMatchStrategy `queryParam:"style=form,explode=true,name=tag_match_strategy"`
	// Return archived incidents
	Archived *bool `queryParam:"style=form,explode=true,name=archived"`
	// Filters for incidents that were updated after this date
	UpdatedAfter *time.Time `queryParam:"style=form,explode=true,name=updated_after"`
	// Filters for incidents that were updated before this date
	UpdatedBefore *time.Time `queryParam:"style=form,explode=true,name=updated_before"`
	// A comma separated list of incident type IDs
	IncidentTypeID *string                            `queryParam:"style=form,explode=true,name=incident_type_id"`
	RequestBody    *GetTicketFunnelMetricsRequestBody `request:"mediaType=multipart/form-data"`
}

func (g GetTicketFunnelMetricsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetTicketFunnelMetricsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetTicketFunnelMetricsRequest) GetConditions() *string {
	if o == nil {
		return nil
	}
	return o.Conditions
}

func (o *GetTicketFunnelMetricsRequest) GetEnvironments() *string {
	if o == nil {
		return nil
	}
	return o.Environments
}

func (o *GetTicketFunnelMetricsRequest) GetServices() *string {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *GetTicketFunnelMetricsRequest) GetFunctionalities() *string {
	if o == nil {
		return nil
	}
	return o.Functionalities
}

func (o *GetTicketFunnelMetricsRequest) GetExcludedInfrastructureIds() *string {
	if o == nil {
		return nil
	}
	return o.ExcludedInfrastructureIds
}

func (o *GetTicketFunnelMetricsRequest) GetTeams() *string {
	if o == nil {
		return nil
	}
	return o.Teams
}

func (o *GetTicketFunnelMetricsRequest) GetAssignedTeams() *string {
	if o == nil {
		return nil
	}
	return o.AssignedTeams
}

func (o *GetTicketFunnelMetricsRequest) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *GetTicketFunnelMetricsRequest) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *GetTicketFunnelMetricsRequest) GetEndDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *GetTicketFunnelMetricsRequest) GetResolvedAtOrAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.ResolvedAtOrAfter
}

func (o *GetTicketFunnelMetricsRequest) GetResolvedAtOrBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.ResolvedAtOrBefore
}

func (o *GetTicketFunnelMetricsRequest) GetCreatedAtOrAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAtOrAfter
}

func (o *GetTicketFunnelMetricsRequest) GetCreatedAtOrBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAtOrBefore
}

func (o *GetTicketFunnelMetricsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *GetTicketFunnelMetricsRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetTicketFunnelMetricsRequest) GetSavedSearchID() *string {
	if o == nil {
		return nil
	}
	return o.SavedSearchID
}

func (o *GetTicketFunnelMetricsRequest) GetPriorities() *string {
	if o == nil {
		return nil
	}
	return o.Priorities
}

func (o *GetTicketFunnelMetricsRequest) GetPriorityNotSet() *bool {
	if o == nil {
		return nil
	}
	return o.PriorityNotSet
}

func (o *GetTicketFunnelMetricsRequest) GetSeverities() *string {
	if o == nil {
		return nil
	}
	return o.Severities
}

func (o *GetTicketFunnelMetricsRequest) GetSeverityNotSet() *bool {
	if o == nil {
		return nil
	}
	return o.SeverityNotSet
}

func (o *GetTicketFunnelMetricsRequest) GetCurrentMilestones() *string {
	if o == nil {
		return nil
	}
	return o.CurrentMilestones
}

func (o *GetTicketFunnelMetricsRequest) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *GetTicketFunnelMetricsRequest) GetTagMatchStrategy() *GetTicketFunnelMetricsQueryParamTagMatchStrategy {
	if o == nil {
		return nil
	}
	return o.TagMatchStrategy
}

func (o *GetTicketFunnelMetricsRequest) GetArchived() *bool {
	if o == nil {
		return nil
	}
	return o.Archived
}

func (o *GetTicketFunnelMetricsRequest) GetUpdatedAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAfter
}

func (o *GetTicketFunnelMetricsRequest) GetUpdatedBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedBefore
}

func (o *GetTicketFunnelMetricsRequest) GetIncidentTypeID() *string {
	if o == nil {
		return nil
	}
	return o.IncidentTypeID
}

func (o *GetTicketFunnelMetricsRequest) GetRequestBody() *GetTicketFunnelMetricsRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type GetTicketFunnelMetricsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Returns a report with task and follow up creation and completion data
	MetricsTicketFunnelMetricsEntity *shared.MetricsTicketFunnelMetricsEntity
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

func (o *GetTicketFunnelMetricsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTicketFunnelMetricsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTicketFunnelMetricsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTicketFunnelMetricsResponse) GetMetricsTicketFunnelMetricsEntity() *shared.MetricsTicketFunnelMetricsEntity {
	if o == nil {
		return nil
	}
	return o.MetricsTicketFunnelMetricsEntity
}

func (o *GetTicketFunnelMetricsResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *GetTicketFunnelMetricsResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *GetTicketFunnelMetricsResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *GetTicketFunnelMetricsResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *GetTicketFunnelMetricsResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *GetTicketFunnelMetricsResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}
