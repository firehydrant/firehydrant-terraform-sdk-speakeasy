// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/internal/utils"
	"time"
)

// ServiceEntityChecklistLabels - An object of label key and values
type ServiceEntityChecklistLabels struct {
}

type ServiceEntityChecklist struct {
	ID            *string    `json:"id,omitempty"`
	Name          *string    `json:"name,omitempty"`
	Description   *string    `json:"description,omitempty"`
	Slug          *string    `json:"slug,omitempty"`
	ServiceTier   *int       `json:"service_tier,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
	AllowedParams []string   `json:"allowed_params,omitempty"`
	// An object of label key and values
	Labels                    *ServiceEntityChecklistLabels `json:"labels,omitempty"`
	AlertOnAdd                *bool                         `json:"alert_on_add,omitempty"`
	AutoAddRespondingTeam     *bool                         `json:"auto_add_responding_team,omitempty"`
	CompletedChecks           *int                          `json:"completed_checks,omitempty"`
	Owner                     *TeamEntityLite               `json:"owner,omitempty"`
	ServiceChecklistUpdatedAt *time.Time                    `json:"service_checklist_updated_at,omitempty"`
}

func (s ServiceEntityChecklist) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *ServiceEntityChecklist) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ServiceEntityChecklist) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ServiceEntityChecklist) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ServiceEntityChecklist) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ServiceEntityChecklist) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *ServiceEntityChecklist) GetServiceTier() *int {
	if o == nil {
		return nil
	}
	return o.ServiceTier
}

func (o *ServiceEntityChecklist) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ServiceEntityChecklist) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *ServiceEntityChecklist) GetAllowedParams() []string {
	if o == nil {
		return nil
	}
	return o.AllowedParams
}

func (o *ServiceEntityChecklist) GetLabels() *ServiceEntityChecklistLabels {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *ServiceEntityChecklist) GetAlertOnAdd() *bool {
	if o == nil {
		return nil
	}
	return o.AlertOnAdd
}

func (o *ServiceEntityChecklist) GetAutoAddRespondingTeam() *bool {
	if o == nil {
		return nil
	}
	return o.AutoAddRespondingTeam
}

func (o *ServiceEntityChecklist) GetCompletedChecks() *int {
	if o == nil {
		return nil
	}
	return o.CompletedChecks
}

func (o *ServiceEntityChecklist) GetOwner() *TeamEntityLite {
	if o == nil {
		return nil
	}
	return o.Owner
}

func (o *ServiceEntityChecklist) GetServiceChecklistUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ServiceChecklistUpdatedAt
}