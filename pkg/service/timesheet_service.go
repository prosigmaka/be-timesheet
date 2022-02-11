package service

import (
	"be-timesheet/pkg/entity"
	"be-timesheet/pkg/repository"
)

type Service interface {
	GetAllTimesheet() ([]entity.Timesheet, error)
	GetTimesheetByID(ID int) (entity.Timesheet, error)
	AddTimesheet(timesheetRequest entity.TimesheetRequest) (entity.Timesheet, error)
	UpdateTimesheet(ID int, timesheetRequest entity.TimesheetRequest) (entity.Timesheet, error)
	DeleteTimesheet(ID int) (entity.Timesheet, error)
}

type service struct {
	repository repository.Repository
}

func NewServiceTimesheet(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) GetAllTimesheet() ([]entity.Timesheet, error) {
	timesheets, err := s.repository.GetAllTimesheet()
	return timesheets, err

}

func (s *service) GetTimesheetByID(ID int) (entity.Timesheet, error) {
	timesheet, err := s.repository.GetTimesheetByID(ID)
	return timesheet, err
}

func (s *service) AddTimesheet(timesheetRequest entity.TimesheetRequest) (entity.Timesheet, error) {
	timesheet := entity.Timesheet{
		Date:          timesheetRequest.Date,
		WorkingStart:  timesheetRequest.WorkingStart,
		WorkingEnd:    timesheetRequest.WorkingEnd,
		OvertimeStart: timesheetRequest.OvertimeStart,
		OvertimeEnd:   timesheetRequest.OvertimeEnd,
		Activity:      timesheetRequest.Activity,
		ProjectID:     timesheetRequest.ProjectID,
		StatusID:      timesheetRequest.StatusID,
	}
	newTimesheet, err := s.repository.AddTimesheet(timesheet)
	return newTimesheet, err
}

func (s *service) UpdateTimesheet(ID int, timesheetRequest entity.TimesheetRequest) (entity.Timesheet, error) {
	timesheet, err := s.repository.GetTimesheetByID(ID)

	timesheet.Date = timesheetRequest.Date
	timesheet.WorkingStart = timesheetRequest.WorkingStart
	timesheet.WorkingEnd = timesheetRequest.WorkingEnd
	timesheet.OvertimeStart = timesheetRequest.OvertimeStart
	timesheet.OvertimeEnd = timesheetRequest.OvertimeEnd
	timesheet.Activity = timesheetRequest.Activity
	timesheet.ProjectID = timesheetRequest.ProjectID
	timesheet.StatusID = timesheetRequest.StatusID

	updatedTimesheet, err := s.repository.UpdateTimesheet(timesheet)
	return updatedTimesheet, err

}

func (s *service) DeleteTimesheet(ID int) (entity.Timesheet, error) {
	timesheet, err := s.repository.GetTimesheetByID(ID)

	deleteBook, err := s.repository.DeleteTimesheet(timesheet)
	return deleteBook, err
}
