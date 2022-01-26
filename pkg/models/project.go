package models

type Project struct {
	Id              int    `json:"id" gorm:"primaryKey"`
	ProjectName     string `json:"projectname"`
	PlacemetAddress string `json:"placement"`
	StartPeriode    string `json:"startperiode"`
	EndPeriode      string `json:"endperiode"`
}