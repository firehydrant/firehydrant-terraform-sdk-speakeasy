// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/internal/utils"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/types"
	"time"
)

type MetricsMilestonesFunnelEntityDataBucketFilterParamsEntity struct {
	// The start datetime for the period
	StartDate *time.Time `json:"start_date,omitempty"`
	// The end datetime for the period not inclusive
	EndDate *types.Date `json:"end_date,omitempty"`
}

func (m MetricsMilestonesFunnelEntityDataBucketFilterParamsEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MetricsMilestonesFunnelEntityDataBucketFilterParamsEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MetricsMilestonesFunnelEntityDataBucketFilterParamsEntity) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *MetricsMilestonesFunnelEntityDataBucketFilterParamsEntity) GetEndDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.EndDate
}