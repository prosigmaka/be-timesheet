package service

import (
	"be-timesheet/pkg/entity"
	"be-timesheet/pkg/repository"
)

type StatusService interface {
	GetAllStatus() ([]entity.Status, error)
	GetStatusByID(ID int) (entity.Status, error)
}

type statusService struct {
	statusRepo repository.StatusRepository
}

func NewServiceStatus(statusRepo repository.StatusRepository) *statusService {
	return &statusService{statusRepo}
}

func (s *statusService) GetAllStatus() ([]entity.Status, error) {
	status, err := s.statusRepo.GetAllStatus()
	return status, err
}

func (s *statusService) GetStatusByID(ID int) (entity.Status, error) {
	status, err := s.statusRepo.GetStatusByID(ID)
	return status, err
}
