package entity

import "strings"

type Timesheet struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	UserID        int    `json:"user_id"`
	Date          string `json:"date"`
	WorkingStart  string `json:"working_start"`
	WorkingEnd    string `json:"working_end"`
	OvertimeStart string `json:"overtime_start"`
	OvertimeEnd   string `json:"overtime_end"`
	Activity      string `json:"activity"`
	ProjectID     int    `json:"project_id"`
	StatusID      int    `json:"status_id"`
}

type TimesheetRequest struct {
	Date          string `json:"date"`
	UserID        int    `json:"user_id"`
	WorkingStart  string `json:"working_start"`
	WorkingEnd    string `json:"working_end"`
	OvertimeStart string `json:"overtime_start"`
	OvertimeEnd   string `json:"overtime_end"`
	Activity      string `json:"activity"`
	ProjectID     int    `json:"project_id"`
	StatusID      int    `json:"status_id"`
}

type TimesheetResponse struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	UserID        int    `json:"user_id"`
	Date          string `json:"date"`
	WorkingStart  string `json:"working_start"`
	WorkingEnd    string `json:"working_end"`
	OvertimeStart string `json:"overtime_start"`
	OvertimeEnd   string `json:"overtime_end"`
	Activity      string `json:"activity"`
	ProjectID     int    `json:"project_id"`
	StatusID      int    `json:"status_id"`
}

func (t *TimesheetRequest) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)

	switch strings.ToLower(action) {
	case "update":
		if t.Date == "" || t.Date == "null" {
			errorMessages["date_required"] = "title is required"
		}
		if t.WorkingStart == "" || t.WorkingStart == "null" {
			errorMessages["workingStart_required"] = "working start is required"
		}
		if t.WorkingEnd == "" || t.WorkingEnd == "null" {
			errorMessages["workingEnd_required"] = "working end is required"
		}
		if t.Activity == "" || t.Activity == "null" {
			errorMessages["activity_required"] = "activity is required"
		}
	default:
		if t.Date == "" || t.Date == "null" {
			errorMessages["date_required"] = "title is required"
		}
		if t.WorkingStart == "" || t.WorkingStart == "null" {
			errorMessages["workingStart_required"] = "working start is required"
		}
		if t.WorkingEnd == "" || t.WorkingEnd == "null" {
			errorMessages["workingEnd_required"] = "working end is required"
		}
		if t.Activity == "" || t.Activity == "null" {
			errorMessages["activity_required"] = "activity is required"
		}
	}
	return errorMessages
}
