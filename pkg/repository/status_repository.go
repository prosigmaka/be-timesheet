package repository

import (
	"be-timesheet/pkg/entity"

	"gorm.io/gorm"
)

type StatusRepository interface {
	GetAllStatus() ([]entity.Status, error)
	GetStatusByID(ID int) (entity.Status, error)
}

type repositoryStatus struct {
	db *gorm.DB
}

func NewStatusRepository(db *gorm.DB) *repositoryStatus {
	return &repositoryStatus{db}
}

func (r *repositoryStatus) GetAllStatus() ([]entity.Status, error) {
	var status []entity.Status

	err := r.db.Find(&status).Error

	return status, err
}

func (r *repositoryStatus) GetStatusByID(ID int) (entity.Status, error) {
	var status entity.Status

	err := r.db.First(&status, ID).Error

	return status, err
}
