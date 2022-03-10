package repository

import (
	"be-timesheet/pkg/entity"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	GetAllProject() ([]entity.Project, error)
	GetProjectByID(ID int) (entity.Project, error)
	AddProject(project entity.Project) (entity.Project, error)
	UpdateProject(project entity.Project) (entity.Project, error)
	DeletProject(ID int) error
}

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *projectRepository {
	return &projectRepository{db}
}

func (r *projectRepository) GetAllProject() ([]entity.Project, error) {
	var projects []entity.Project

	err := r.db.Order("id asc").Find(&projects).Error

	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *projectRepository) GetProjectByID(ID int) (entity.Project, error) {
	var project entity.Project

	err := r.db.Where("id=?",ID).Take(&project).Error

	return project, err
}

func (r *projectRepository) AddProject(project entity.Project) (entity.Project, error) {
	err := r.db.Create(&project).Error
	return project, err
}

func (r *projectRepository) UpdateProject(project entity.Project) (entity.Project, error) {
	err := r.db.Save(&project).Error
	return project, err
}

func (r *projectRepository) DeletProject(ID int) error {
	var project entity.Project
	err := r.db.Where("id=?",ID).Delete(&project).Error
	if err != nil {
		return err
	}
	return nil
}
