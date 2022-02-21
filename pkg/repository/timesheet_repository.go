package repository

import (
	"be-timesheet/pkg/entity"

	"gorm.io/gorm"
)

type Repository interface {
	GetAllTimesheet() ([]entity.Timesheet, error)
	// GetTimesheetByID(ID int) (*entity.Timesheet, error)
	GetTimesheetByID(ID int) (entity.Timesheet, error)
	// AddTimesheet(timesheet *entity.Timesheet) (*entity.Timesheet, error)
	AddTimesheet(timesheet entity.Timesheet) (entity.Timesheet, error)
	// UpdateTimesheet(timesheet *entity.Timesheet) (*entity.Timesheet, error)
	UpdateTimesheet(timesheet entity.Timesheet) (entity.Timesheet, error)
	DeleteTimesheet(ID int) error
}

type repository struct {
	db *gorm.DB
}

func NewTimesheetRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllTimesheet() ([]entity.Timesheet, error) {
	var timesheets []entity.Timesheet

	err := r.db.Order("id asc").Find(&timesheets).Error

	if err != nil {
		return nil, err
	}
	return timesheets, nil
}

// func (r *repository) GetTimesheetByID(ID int) (*entity.Timesheet, error) {
// 	var timesheet entity.Timesheet

// 	err := r.db.Where("id = ?", ID).Take(&timesheet).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &timesheet, nil
// }

func (r *repository) GetTimesheetByID(ID int) (entity.Timesheet, error) {
	var timesheet entity.Timesheet

	err := r.db.Where("id = ?", ID).Take(&timesheet).Error

	return timesheet, err
}

// func (r *repository) AddTimesheet(timesheet *entity.Timesheet) (*entity.Timesheet, error) {
// 	err := r.db.Create(&timesheet).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return timesheet, nil
// }

func (r *repository) AddTimesheet(timesheet entity.Timesheet) (entity.Timesheet, error) {
	err := r.db.Create(&timesheet).Error
	return timesheet, err
}

// func (r *repository) UpdateTimesheet(timesheet *entity.Timesheet) (*entity.Timesheet, error) {
// 	err := r.db.Save(&timesheet).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return timesheet, nil
// }

func (r *repository) UpdateTimesheet(timesheet entity.Timesheet) (entity.Timesheet, error) {
	err := r.db.Save(&timesheet).Error
	return timesheet, err
}

// func (r *repository) UpdateTimesheet(timesheet entity.Timesheet) (entity.Timesheet, error) {
// 	err := r.db.Save(&timesheet).Error
// 	return timesheet, err
// }

func (r *repository) DeleteTimesheet(ID int) error {
	var timesheet entity.Timesheet
	err := r.db.Where("id = ?", ID).Delete(&timesheet).Error
	if err != nil {
		return err
	}

	return nil
}
