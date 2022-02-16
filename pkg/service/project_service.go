package service

import (
	"timesheet-be/pkg/entity"
	"timesheet-be/pkg/repository"
)

type Service interface {
	GetAllProject() ([]entity.Project, error)
	GetProjectByID(ID int) (entity.Project, error)
	AddProject(projectRequest entity.ProjectRequest) (entity.Project, error)
	UpdateProject(ID int, projectRequest entity.ProjectRequest) (entity.Project, error)
	DeletProject(ID int) (entity.Project, error)
}

type service struct {
	repository repository.Repository
}

func NewServiceProject(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) GetAllProject() ([]entity.Project, error) {
	projects, err := s.repository.GetAllProject()
	return projects, err
}

func (s *service) GetProjectByID(ID int) (entity.Project, error) {
	project, err := s.repository.GetProjectByID(ID)
	return project, err
}


func (s *service) AddProject(projectRequest entity.ProjectRequest) (entity.Project, error) {
	project := entity.Project{
		ProjectName: projectRequest.ProjectName,
		PlacementAddress: projectRequest.PlacementAddress,
		StartDate: projectRequest.StartDate,
		EndDate: projectRequest.EndDate,
	}
	newProject, err := s.repository.AddProject(project)
	return newProject, err
}

func (s *service) UpdateProject(ID int, projectRequest entity.ProjectRequest) (entity.Project, error) {
	project, err := s.repository.GetProjectByID(ID)

	project.ProjectName = projectRequest.ProjectName
	project.PlacementAddress = projectRequest.PlacementAddress
	project.StartDate = projectRequest.StartDate
	project.EndDate = projectRequest.EndDate

	updatedProject, err := s.repository.UpdateProject(project)
	return updatedProject, err
}

func (s *service) DeletProject(ID int) (entity.Project, error) {
	project, err := s.repository.GetProjectByID(ID)

	deleteProject, err := s.repository.DeletProject(project)
	return deleteProject, err
}