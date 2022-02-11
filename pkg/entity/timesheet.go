package entity

type Timesheet struct {
	ID            int
	Date          string
	WorkingStart  string
	WorkingEnd    string
	OvertimeStart string
	OvertimeEnd   string
	Activity      string
	ProjectID     int
	StatusID      int
}

type TimesheetRequest struct {
	Date          string `json:"date" binding:"required"`
	WorkingStart  string `json:"working_start"`
	WorkingEnd    string `json:"working_end"`
	OvertimeStart string `json:"overtime_start"`
	OvertimeEnd   string `json:"overtime_end"`
	Activity      string `json:"activity" binding:"required"`
	ProjectID     int    `json:"project_id"`
	StatusID      int    `json:"status_id"`
}

type TimesheetResponse struct {
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
