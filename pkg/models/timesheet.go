package models

type Timesheet struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	Date          string `json:"date"`
	WorkingStart  string `json:"working_start"`
	WorkingEnd    string `json:"working_end"`
	OvertimeStart string `json:"overtime_start"`
	OvertimeEnd   string `json:"overtime_end"`
	Activity      string `json:"activity"`
	ProjectID     int    `json:"project_id"`
	StatusID      int    `json:"status_id"`
}
