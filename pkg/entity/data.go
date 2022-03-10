package entity

type Data struct {
	ID            int
	UserID        int
	TimesheetID   int
	ProjectID     int
	Employee_name string
	Project_name  string
	StartDate     string
	EndDate       string
}

type DataRequest struct {
	UserID        int    `json:"user_id"`
	TimesheetID   int    `json:"timesheet_id"`
	ProjectID     int    `json:"project_id"`
	Employee_name string `json:"employee_name"`
	Project_name  string `json:"project_name"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
}

type DataResponse struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	UserID        int    `json:"user_id" gorm:"foreignKey"`
	TimesheetID   int    `json:"timesheet_id" gorm:"foreignKey"`
	ProjectID     int    `json:"project_id" gorm:"foreignKey"`
	Employee_name string `json:"empoloyee_name"`
	Project_name  string `json:"project_name"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
}