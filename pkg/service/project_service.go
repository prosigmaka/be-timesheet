package service

import (
	"be-timesheet/pkg/entity"
	"be-timesheet/pkg/repository"
)

type ProjectService interface {
	GetAllProject() ([]entity.ProjectResponse, error)
	GetProjectByID(ID int) (entity.Project, error)
	AddProject(projectRequest entity.ProjectRequest) (entity.Project, error)
	UpdateProject(ID int, projectRequest entity.ProjectRequest) (entity.Project, error)
	DeletProject(ID int) error
}

type projectService struct {
	projectRepo repository.ProjectRepository
}

func NewServiceProject(projectRepo repository.ProjectRepository) *projectService {
	return &projectService{projectRepo}
}

func (s *projectService) GetAllProject() ([]entity.ProjectResponse, error) {
	result, err := s.projectRepo.GetAllProject()
	if err != nil {
		return nil, err
	}

	var projectList []entity.ProjectResponse

	if result != nil {
		for _, item := range result {
			project := entity.ProjectResponse{
				ID: item.ID,
				Project_name: item.Project_name,
				PlacementAddress: item.PlacementAddress,
				StartDate: item.StartDate,
				EndDate: item.EndDate,
			}
			projectList = append(projectList, project)
		}
	}
	return projectList, nil
}

func (s *projectService) GetProjectByID(ID int) (entity.Project, error) {
	project, err := s.projectRepo.GetProjectByID(ID)
	return project, err
}

func (s *projectService) AddProject(projectRequest entity.ProjectRequest) (entity.Project, error) {
	project := entity.Project{
		Project_name:      projectRequest.Project_name,
		PlacementAddress: projectRequest.PlacementAddress,
		StartDate:        projectRequest.StartDate,
		EndDate:          projectRequest.EndDate,
	}
	newProject, err := s.projectRepo.AddProject(project)
	return newProject, err
}

func (s *projectService) UpdateProject(ID int, projectRequest entity.ProjectRequest) (entity.Project, error) {
	project, err := s.projectRepo.GetProjectByID(ID)

	project.Project_name = projectRequest.Project_name
	project.PlacementAddress = projectRequest.PlacementAddress
	project.StartDate = projectRequest.StartDate
	project.EndDate = projectRequest.EndDate

	updatedProject, err := s.projectRepo.UpdateProject(project)
	return updatedProject, err
}

func (s *projectService) DeletProject(ID int) error {
	err := s.projectRepo.DeletProject(ID)
	if err != nil {
		return err
	}
	return nil
}