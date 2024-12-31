// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type PostV1TeamsTeamIDOnCallSchedulesMembers struct {
	// The ID of a user who should be added to the schedule's rotation. You can add a user to the schedule
	// multiple times to construct more complex rotations, and you can specify a `null` user ID to create
	// unassigned slots in the rotation.
	//
	UserID *string `json:"user_id,omitempty"`
}

func (o *PostV1TeamsTeamIDOnCallSchedulesMembers) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

// PostV1TeamsTeamIDOnCallSchedulesType - The type of strategy. Must be one of "daily", "weekly", or "custom".
type PostV1TeamsTeamIDOnCallSchedulesType string

const (
	PostV1TeamsTeamIDOnCallSchedulesTypeDaily  PostV1TeamsTeamIDOnCallSchedulesType = "daily"
	PostV1TeamsTeamIDOnCallSchedulesTypeWeekly PostV1TeamsTeamIDOnCallSchedulesType = "weekly"
	PostV1TeamsTeamIDOnCallSchedulesTypeCustom PostV1TeamsTeamIDOnCallSchedulesType = "custom"
)

func (e PostV1TeamsTeamIDOnCallSchedulesType) ToPointer() *PostV1TeamsTeamIDOnCallSchedulesType {
	return &e
}
func (e *PostV1TeamsTeamIDOnCallSchedulesType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "daily":
		fallthrough
	case "weekly":
		fallthrough
	case "custom":
		*e = PostV1TeamsTeamIDOnCallSchedulesType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostV1TeamsTeamIDOnCallSchedulesType: %v", v)
	}
}

// HandoffDay - The day of the week on which on-call shifts should hand off, as its long-form name (e.g. "monday", "tuesday", etc). This value is only used if the strategy type is "weekly".
type HandoffDay string

const (
	HandoffDayMonday    HandoffDay = "monday"
	HandoffDayTuesday   HandoffDay = "tuesday"
	HandoffDayWednesday HandoffDay = "wednesday"
	HandoffDayThursday  HandoffDay = "thursday"
	HandoffDayFriday    HandoffDay = "friday"
	HandoffDaySaturday  HandoffDay = "saturday"
	HandoffDaySunday    HandoffDay = "sunday"
)

func (e HandoffDay) ToPointer() *HandoffDay {
	return &e
}
func (e *HandoffDay) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "monday":
		fallthrough
	case "tuesday":
		fallthrough
	case "wednesday":
		fallthrough
	case "thursday":
		fallthrough
	case "friday":
		fallthrough
	case "saturday":
		fallthrough
	case "sunday":
		*e = HandoffDay(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HandoffDay: %v", v)
	}
}

// Strategy - An object that specifies how the schedule's on-call shifts should be generated.
type Strategy struct {
	// The type of strategy. Must be one of "daily", "weekly", or "custom".
	Type PostV1TeamsTeamIDOnCallSchedulesType `json:"type"`
	// An ISO8601 time string specifying when on-call shifts should hand off. This value is only used if the strategy type is "daily" or "weekly".
	HandoffTime *string `json:"handoff_time,omitempty"`
	// The day of the week on which on-call shifts should hand off, as its long-form name (e.g. "monday", "tuesday", etc). This value is only used if the strategy type is "weekly".
	HandoffDay *HandoffDay `json:"handoff_day,omitempty"`
	// An ISO8601 duration string specifying how long each shift should last. This value is only used if the strategy type is "custom".
	ShiftDuration *string `json:"shift_duration,omitempty"`
}

func (o *Strategy) GetType() PostV1TeamsTeamIDOnCallSchedulesType {
	if o == nil {
		return PostV1TeamsTeamIDOnCallSchedulesType("")
	}
	return o.Type
}

func (o *Strategy) GetHandoffTime() *string {
	if o == nil {
		return nil
	}
	return o.HandoffTime
}

func (o *Strategy) GetHandoffDay() *HandoffDay {
	if o == nil {
		return nil
	}
	return o.HandoffDay
}

func (o *Strategy) GetShiftDuration() *string {
	if o == nil {
		return nil
	}
	return o.ShiftDuration
}

// StartDay - The day of the week on which the restriction should start, as its long-form name (e.g. "monday", "tuesday", etc).
type StartDay string

const (
	StartDayMonday    StartDay = "monday"
	StartDayTuesday   StartDay = "tuesday"
	StartDayWednesday StartDay = "wednesday"
	StartDayThursday  StartDay = "thursday"
	StartDayFriday    StartDay = "friday"
	StartDaySaturday  StartDay = "saturday"
	StartDaySunday    StartDay = "sunday"
)

func (e StartDay) ToPointer() *StartDay {
	return &e
}
func (e *StartDay) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "monday":
		fallthrough
	case "tuesday":
		fallthrough
	case "wednesday":
		fallthrough
	case "thursday":
		fallthrough
	case "friday":
		fallthrough
	case "saturday":
		fallthrough
	case "sunday":
		*e = StartDay(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StartDay: %v", v)
	}
}

// EndDay - The day of the week on which the restriction should end, as its long-form name (e.g. "monday", "tuesday", etc).
type EndDay string

const (
	EndDayMonday    EndDay = "monday"
	EndDayTuesday   EndDay = "tuesday"
	EndDayWednesday EndDay = "wednesday"
	EndDayThursday  EndDay = "thursday"
	EndDayFriday    EndDay = "friday"
	EndDaySaturday  EndDay = "saturday"
	EndDaySunday    EndDay = "sunday"
)

func (e EndDay) ToPointer() *EndDay {
	return &e
}
func (e *EndDay) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "monday":
		fallthrough
	case "tuesday":
		fallthrough
	case "wednesday":
		fallthrough
	case "thursday":
		fallthrough
	case "friday":
		fallthrough
	case "saturday":
		fallthrough
	case "sunday":
		*e = EndDay(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EndDay: %v", v)
	}
}

type Restrictions struct {
	// The day of the week on which the restriction should start, as its long-form name (e.g. "monday", "tuesday", etc).
	StartDay StartDay `json:"start_day"`
	// An ISO8601 time string specifying when the restriction should start.
	StartTime string `json:"start_time"`
	// The day of the week on which the restriction should end, as its long-form name (e.g. "monday", "tuesday", etc).
	EndDay EndDay `json:"end_day"`
	// An ISO8601 time string specifying when the restriction should end.
	EndTime string `json:"end_time"`
}

func (o *Restrictions) GetStartDay() StartDay {
	if o == nil {
		return StartDay("")
	}
	return o.StartDay
}

func (o *Restrictions) GetStartTime() string {
	if o == nil {
		return ""
	}
	return o.StartTime
}

func (o *Restrictions) GetEndDay() EndDay {
	if o == nil {
		return EndDay("")
	}
	return o.EndDay
}

func (o *Restrictions) GetEndTime() string {
	if o == nil {
		return ""
	}
	return o.EndTime
}

// PostV1TeamsTeamIDOnCallSchedules - Create a Signals on-call schedule for a team.
type PostV1TeamsTeamIDOnCallSchedules struct {
	// The on-call schedule's name.
	Name string `json:"name"`
	// A detailed description of the on-call schedule.
	Description *string `json:"description,omitempty"`
	// The time zone in which the on-call schedule operates. This value must be a valid IANA time zone name.
	TimeZone string `json:"time_zone"`
	// The ID of a Slack user group for syncing purposes. If provided, we will automatically sync whoever is on call to the user group in Slack.
	SlackUserGroupID *string `json:"slack_user_group_id,omitempty"`
	// An ordered list of objects that specify members of the on-call schedule's rotation.
	Members []PostV1TeamsTeamIDOnCallSchedulesMembers `json:"members,omitempty"`
	// An object that specifies how the schedule's on-call shifts should be generated.
	Strategy Strategy `json:"strategy"`
	// A list of objects that restrict the schedule to speccific on-call periods.
	Restrictions []Restrictions `json:"restrictions,omitempty"`
	// An ISO8601 time string specifying when the schedule's first shift should start. This value is only used if the schedule's strategy is "custom".
	StartTime *string `json:"start_time,omitempty"`
	// A hex color code that will be used to represent the schedule in the UI and iCal subscriptions.
	Color *string `json:"color,omitempty"`
	// This parameter is deprecated; use `members` instead.
	MemberIds []string `json:"member_ids,omitempty"`
}

func (o *PostV1TeamsTeamIDOnCallSchedules) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PostV1TeamsTeamIDOnCallSchedules) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PostV1TeamsTeamIDOnCallSchedules) GetTimeZone() string {
	if o == nil {
		return ""
	}
	return o.TimeZone
}

func (o *PostV1TeamsTeamIDOnCallSchedules) GetSlackUserGroupID() *string {
	if o == nil {
		return nil
	}
	return o.SlackUserGroupID
}

func (o *PostV1TeamsTeamIDOnCallSchedules) GetMembers() []PostV1TeamsTeamIDOnCallSchedulesMembers {
	if o == nil {
		return nil
	}
	return o.Members
}

func (o *PostV1TeamsTeamIDOnCallSchedules) GetStrategy() Strategy {
	if o == nil {
		return Strategy{}
	}
	return o.Strategy
}

func (o *PostV1TeamsTeamIDOnCallSchedules) GetRestrictions() []Restrictions {
	if o == nil {
		return nil
	}
	return o.Restrictions
}

func (o *PostV1TeamsTeamIDOnCallSchedules) GetStartTime() *string {
	if o == nil {
		return nil
	}
	return o.StartTime
}

func (o *PostV1TeamsTeamIDOnCallSchedules) GetColor() *string {
	if o == nil {
		return nil
	}
	return o.Color
}

func (o *PostV1TeamsTeamIDOnCallSchedules) GetMemberIds() []string {
	if o == nil {
		return nil
	}
	return o.MemberIds
}