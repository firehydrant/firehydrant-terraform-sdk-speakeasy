// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/internal/utils"
	"time"
)

type TicketingTicketEntityState string

const (
	TicketingTicketEntityStateOpen       TicketingTicketEntityState = "open"
	TicketingTicketEntityStateInProgress TicketingTicketEntityState = "in_progress"
	TicketingTicketEntityStateCancelled  TicketingTicketEntityState = "cancelled"
	TicketingTicketEntityStateDone       TicketingTicketEntityState = "done"
)

func (e TicketingTicketEntityState) ToPointer() *TicketingTicketEntityState {
	return &e
}
func (e *TicketingTicketEntityState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "open":
		fallthrough
	case "in_progress":
		fallthrough
	case "cancelled":
		fallthrough
	case "done":
		*e = TicketingTicketEntityState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TicketingTicketEntityState: %v", v)
	}
}

type TicketingTicketEntityType string

const (
	TicketingTicketEntityTypeIncident TicketingTicketEntityType = "incident"
	TicketingTicketEntityTypeTask     TicketingTicketEntityType = "task"
	TicketingTicketEntityTypeFollowUp TicketingTicketEntityType = "follow_up"
)

func (e TicketingTicketEntityType) ToPointer() *TicketingTicketEntityType {
	return &e
}
func (e *TicketingTicketEntityType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "incident":
		fallthrough
	case "task":
		fallthrough
	case "follow_up":
		*e = TicketingTicketEntityType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TicketingTicketEntityType: %v", v)
	}
}

type TicketingTicketEntityAttachments struct {
}

// TicketingTicketEntity - Ticketing_TicketEntity model
type TicketingTicketEntity struct {
	ID          *string                     `json:"id,omitempty"`
	Summary     *string                     `json:"summary,omitempty"`
	Description *string                     `json:"description,omitempty"`
	State       *TicketingTicketEntityState `json:"state,omitempty"`
	Type        *TicketingTicketEntityType  `json:"type,omitempty"`
	Assignees   []AuthorEntity              `json:"assignees,omitempty"`
	// Ticketing_PriorityEntity model
	Priority  *TicketingPriorityEntity `json:"priority,omitempty"`
	CreatedBy *AuthorEntity            `json:"created_by,omitempty"`
	// A list of objects attached to this item. Can be one of: LinkEntity, CustomerSupportIssueEntity, or GenericAttachmentEntity
	Attachments []TicketingTicketEntityAttachments `json:"attachments,omitempty"`
	CreatedAt   *time.Time                         `json:"created_at,omitempty"`
	UpdatedAt   *time.Time                         `json:"updated_at,omitempty"`
	TagList     []string                           `json:"tag_list,omitempty"`
	// ID of incident that this ticket is related to
	IncidentID *string `json:"incident_id,omitempty"`
	// Name of incident that this ticket is related to
	IncidentName *string `json:"incident_name,omitempty"`
	// Milestone of incident that this ticket is related to
	IncidentCurrentMilestone *string `json:"incident_current_milestone,omitempty"`
	// ID of task that this ticket is related to
	TaskID *string    `json:"task_id,omitempty"`
	DueAt  *time.Time `json:"due_at,omitempty"`
	// Attachments_LinkEntity model
	Link *AttachmentsLinkEntity `json:"link,omitempty"`
}

func (t TicketingTicketEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TicketingTicketEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TicketingTicketEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TicketingTicketEntity) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *TicketingTicketEntity) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *TicketingTicketEntity) GetState() *TicketingTicketEntityState {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *TicketingTicketEntity) GetType() *TicketingTicketEntityType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *TicketingTicketEntity) GetAssignees() []AuthorEntity {
	if o == nil {
		return nil
	}
	return o.Assignees
}

func (o *TicketingTicketEntity) GetPriority() *TicketingPriorityEntity {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *TicketingTicketEntity) GetCreatedBy() *AuthorEntity {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *TicketingTicketEntity) GetAttachments() []TicketingTicketEntityAttachments {
	if o == nil {
		return nil
	}
	return o.Attachments
}

func (o *TicketingTicketEntity) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *TicketingTicketEntity) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *TicketingTicketEntity) GetTagList() []string {
	if o == nil {
		return nil
	}
	return o.TagList
}

func (o *TicketingTicketEntity) GetIncidentID() *string {
	if o == nil {
		return nil
	}
	return o.IncidentID
}

func (o *TicketingTicketEntity) GetIncidentName() *string {
	if o == nil {
		return nil
	}
	return o.IncidentName
}

func (o *TicketingTicketEntity) GetIncidentCurrentMilestone() *string {
	if o == nil {
		return nil
	}
	return o.IncidentCurrentMilestone
}

func (o *TicketingTicketEntity) GetTaskID() *string {
	if o == nil {
		return nil
	}
	return o.TaskID
}

func (o *TicketingTicketEntity) GetDueAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DueAt
}

func (o *TicketingTicketEntity) GetLink() *AttachmentsLinkEntity {
	if o == nil {
		return nil
	}
	return o.Link
}