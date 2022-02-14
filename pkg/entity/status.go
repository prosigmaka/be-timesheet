package entity

type Status struct {
	ID         int
	StatusName string
}

type StatusResponse struct {
	ID         int    `json:"status_id" gorm:"primaryKey"`
	StatusName string `json:"status_name"`
}
