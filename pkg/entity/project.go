package entity

type Project struct {
	ID               int
	ProjectName      string
	PlacementAddress string
	StartDate        string
	EndDate          string
}

type ProjectRequest struct {
	ProjectName      string `json:"project_name" binding:"required"`
	PlacementAddress string `json:"placement_address"`
	StartDate        string `json:"start_date"`
	EndDate          string `json:"end_date"`
}

type ProjectResponse struct {
	ID               int    `json:"id" gorm:"primaryKey"`
	ProjectName      string `json:"project_name"`
	PlacementAddress string `json:"placement_address"`
	StartDate        string `json:"start_date"`
	EndDate          string `json:"end_date"`
}
