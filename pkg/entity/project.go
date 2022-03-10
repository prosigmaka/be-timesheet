package entity

type Project struct {
	ID               int
	UserID           int
	Project_name     string
	PlacementAddress string
	StartDate        string
	EndDate          string
}

type ProjectRequest struct {
	Project_name     string `json:"project_name" binding:"required"`
	UserID           int    `json:"user_id"`
	PlacementAddress string `json:"placement_address"`
	StartDate        string `json:"start_date"`
	EndDate          string `json:"end_date"`
}

type ProjectResponse struct {
	ID               int    `json:"id" gorm:"primaryKey"`
	UserID           int    `json:"user_id"`
	Project_name     string `json:"project_name"`
	PlacementAddress string `json:"placement_address"`
	StartDate        string `json:"start_date"`
	EndDate          string `json:"end_date"`
}
