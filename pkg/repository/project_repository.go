package repository

import (
	"timesheet-be/pkg/entity"

	"gorm.io/gorm"
)

type Repository interface {
	GetAllProject() ([]entity.Project, error)
	GetProjectByID(ID int) (entity.Project, error)
	AddProject(project entity.Project) (entity.Project, error)
	UpdateProject(project entity.Project) (entity.Project, error)
	DeletProject(project entity.Project) (entity.Project, error)
}

type repository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllProject() ([]entity.Project, error) {
	var projects []entity.Project

	err := r.db.Find(&projects).Error

	return projects, err
}

func (r *repository) GetProjectByID(ID int) (entity.Project, error) {
	var project entity.Project

	err := r.db.First(&project, ID).Error

	return project, err
}

func (r *repository) AddProject(project entity.Project) (entity.Project, error) {
	err := r.db.Create(&project).Error
	return project, err
}

func (r *repository) UpdateProject(project entity.Project) (entity.Project, error) {
	err := r.db.Save(&project).Error
	return project, err
}

func (r *repository) DeletProject(project entity.Project) (entity.Project, error) {
	err := r.db.Delete(&project).Error
	return project, err
}
