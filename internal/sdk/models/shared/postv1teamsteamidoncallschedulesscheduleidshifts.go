// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// PostV1TeamsTeamIDOnCallSchedulesScheduleIDShifts - Create a Signals on-call shift in a schedule.
type PostV1TeamsTeamIDOnCallSchedulesScheduleIDShifts struct {
	// The start time of the shift in ISO8601 format.
	StartTime string `json:"start_time"`
	// The end time of the shift in ISO8601 format.
	EndTime string `json:"end_time"`
	// The ID of the user who is on-call for the shift. If not provided, the shift will be unassigned.
	UserID *string `json:"user_id,omitempty"`
}

func (o *PostV1TeamsTeamIDOnCallSchedulesScheduleIDShifts) GetStartTime() string {
	if o == nil {
		return ""
	}
	return o.StartTime
}

func (o *PostV1TeamsTeamIDOnCallSchedulesScheduleIDShifts) GetEndTime() string {
	if o == nil {
		return ""
	}
	return o.EndTime
}

func (o *PostV1TeamsTeamIDOnCallSchedulesScheduleIDShifts) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}
