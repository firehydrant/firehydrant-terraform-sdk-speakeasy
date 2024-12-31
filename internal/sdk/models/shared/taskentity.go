// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/internal/utils"
	"time"
)

// TaskEntity model
type TaskEntity struct {
	ID          *string       `json:"id,omitempty"`
	Title       *string       `json:"title,omitempty"`
	Description *string       `json:"description,omitempty"`
	State       *string       `json:"state,omitempty"`
	Assignee    *AuthorEntity `json:"assignee,omitempty"`
	CreatedBy   *AuthorEntity `json:"created_by,omitempty"`
	CreatedAt   *time.Time    `json:"created_at,omitempty"`
	UpdatedAt   *time.Time    `json:"updated_at,omitempty"`
	DueAt       *time.Time    `json:"due_at,omitempty"`
}

func (t TaskEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaskEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TaskEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TaskEntity) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *TaskEntity) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *TaskEntity) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *TaskEntity) GetAssignee() *AuthorEntity {
	if o == nil {
		return nil
	}
	return o.Assignee
}

func (o *TaskEntity) GetCreatedBy() *AuthorEntity {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *TaskEntity) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *TaskEntity) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *TaskEntity) GetDueAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DueAt
}