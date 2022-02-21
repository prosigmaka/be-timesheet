package service

import (
	"be-timesheet/pkg/entity"
	"be-timesheet/pkg/repository"
)

type Service interface {
	GetAllTimesheet() ([]entity.TimesheetResponse, error)
	// GetTimesheetByID(ID int) (*entity.TimesheetResponse, error)
	GetTimesheetByID(ID int) (entity.Timesheet, error)
	// AddTimesheet(timesheet *entity.TimesheetResponse) (*entity.TimesheetResponse, error)
	AddTimesheet(timesheetRequest entity.TimesheetRequest) (entity.Timesheet, error)
	// UpdateTimesheet(timesheetResponse *entity.TimesheetResponse) (*entity.TimesheetResponse, error)
	UpdateTimesheet(ID int, timesheetRequest entity.TimesheetRequest) (entity.Timesheet, error)
	DeleteTimesheet(ID int) error
}

type service struct {
	repository repository.Repository
}

func NewServiceTimesheet(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) GetAllTimesheet() ([]entity.TimesheetResponse, error) {
	// timesheets, err := s.repository.GetAllTimesheet()
	// return timesheets, err
	result, err := s.repository.GetAllTimesheet()
	if err != nil {
		return nil, err
	}

	var timesheetList []entity.TimesheetResponse

	if result != nil {
		for _, item := range result {
			timesheet := entity.TimesheetResponse{
				ID:            item.ID,
				UserID:        item.UserID,
				Date:          item.Date,
				WorkingStart:  item.WorkingStart,
				WorkingEnd:    item.WorkingEnd,
				OvertimeStart: item.OvertimeStart,
				OvertimeEnd:   item.OvertimeEnd,
				Activity:      item.Activity,
				ProjectID:     item.ProjectID,
				StatusID:      item.StatusID,
			}
			timesheetList = append(timesheetList, timesheet)
		}
	}
	return timesheetList, nil

}

// func (s *service) GetTimesheetByID(ID int) (*entity.TimesheetResponse, error) {
// 	// timesheet, err := s.repository.GetTimesheetByID(ID)
// 	// return timesheet, err
// 	result, err := s.repository.GetTimesheetByID(ID)

// 	if err != nil {
// 		return nil, err
// 	}

// 	var timesheet entity.TimesheetResponse

// 	if result != nil {
// 		timesheet = entity.TimesheetResponse{
// 			ID:            result.ID,
// 			UserID:        result.UserID,
// 			Date:          result.Date,
// 			WorkingStart:  result.WorkingStart,
// 			WorkingEnd:    result.WorkingEnd,
// 			OvertimeStart: result.OvertimeStart,
// 			OvertimeEnd:   result.OvertimeEnd,
// 			Activity:      result.Activity,
// 			ProjectID:     result.ProjectID,
// 			StatusID:      result.StatusID,
// 		}
// 	}

// 	return &timesheet, nil
// }

func (s *service) GetTimesheetByID(ID int) (entity.Timesheet, error) {
	timesheet, err := s.repository.GetTimesheetByID(ID)
	return timesheet, err
}

// func (s *service) AddTimesheet(timesheet *entity.TimesheetResponse) (*entity.TimesheetResponse, error) {
// 	var timesheetRes = entity.Timesheet{
// 		Date:          timesheet.Date,
// 		UserID:        timesheet.UserID,
// 		WorkingStart:  timesheet.WorkingStart,
// 		WorkingEnd:    timesheet.WorkingEnd,
// 		OvertimeStart: timesheet.OvertimeStart,
// 		OvertimeEnd:   timesheet.OvertimeEnd,
// 		Activity:      timesheet.Activity,
// 		ProjectID:     timesheet.ProjectID,
// 		StatusID:      timesheet.StatusID,
// 	}
// 	result, err := s.repository.AddTimesheet(&timesheetRes)
// 	// return newTimesheet, err

// 	if err != nil {
// 		return nil, err
// 	}

// 	// var timesheetResponse entity.TimesheetResponse

// 	if result != nil {
// 		timesheet = &entity.TimesheetResponse{
// 			ID:            result.ID,
// 			Date:          result.Date,
// 			UserID:        result.UserID,
// 			WorkingStart:  result.WorkingStart,
// 			WorkingEnd:    result.WorkingEnd,
// 			OvertimeStart: result.OvertimeStart,
// 			OvertimeEnd:   result.OvertimeEnd,
// 			Activity:      result.Activity,
// 			ProjectID:     result.ProjectID,
// 			StatusID:      result.StatusID,
// 		}
// 	}

// 	return timesheet, nil

// 	// result, err := s.repository.AddTimesheet(&timesheet)

// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// if result != nil {
// 	// 	timesheetRequest = &entity.TimesheetRequest{
// 	// 		ID: result.ID,
// 	// 		UserID: result.UserID,
// 	// 		Title: result.Title,
// 	// 		Description: result.Description,
// 	// 		FoodImage: result.FoodImage,
// 	// 	}
// 	// }

// 	// return foodVM, nil
// }

func (s *service) AddTimesheet(timesheetRequest entity.TimesheetRequest) (entity.Timesheet, error) {
	timesheet := entity.Timesheet{
		Date:          timesheetRequest.Date,
		UserID:        timesheetRequest.UserID,
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

// func (s *service) UpdateTimesheet(timesheetResponse *entity.TimesheetResponse) (*entity.TimesheetResponse, error) {
// 	// timesheet, _ := s.repository.GetTimesheetByID(ID)

// 	// timesheet.Date = timesheetRequest.Date
// 	// timesheet.UserID = timesheetRequest.UserID
// 	// timesheet.WorkingStart = timesheetRequest.WorkingStart
// 	// timesheet.WorkingEnd = timesheetRequest.WorkingEnd
// 	// timesheet.OvertimeStart = timesheetRequest.OvertimeStart
// 	// timesheet.OvertimeEnd = timesheetRequest.OvertimeEnd
// 	// timesheet.Activity = timesheetRequest.Activity
// 	// timesheet.ProjectID = timesheetRequest.ProjectID
// 	// timesheet.StatusID = timesheetRequest.StatusID

// 	// updatedTimesheet, err := s.repository.UpdateTimesheet(&timesheet)
// 	// return updatedTimesheet, err
// 	var timesheet = entity.Timesheet{
// 		ID:            timesheetResponse.ID,
// 		Date:          timesheetResponse.Date,
// 		UserID:        timesheetResponse.UserID,
// 		WorkingStart:  timesheetResponse.WorkingStart,
// 		WorkingEnd:    timesheetResponse.WorkingEnd,
// 		OvertimeStart: timesheetResponse.OvertimeStart,
// 		OvertimeEnd:   timesheetResponse.OvertimeEnd,
// 		Activity:      timesheetResponse.Activity,
// 		ProjectID:     timesheetResponse.ProjectID,
// 		StatusID:      timesheetResponse.StatusID,
// 	}

// 	_, err := s.repository.UpdateTimesheet(&timesheet)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return timesheetResponse, nil
// }

func (s *service) UpdateTimesheet(ID int, timesheetRequest entity.TimesheetRequest) (entity.Timesheet, error) {
	timesheet, err := s.repository.GetTimesheetByID(ID)

	timesheet.Date = timesheetRequest.Date
	timesheet.UserID = timesheetRequest.UserID
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

func (s *service) DeleteTimesheet(ID int) error {
	// timesheet, err := s.repository.GetTimesheetByID(ID)

	// deleteTimesheet, err := s.repository.DeleteTimesheet(timesheet)
	// return deleteTimesheet, err
	err := s.repository.DeleteTimesheet(ID)
	if err != nil {
		return err
	}

	return nil
}
