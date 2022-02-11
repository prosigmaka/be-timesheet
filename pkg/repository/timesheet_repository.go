package repository

import (
	"be-timesheet/pkg/entity"

	"gorm.io/gorm"
)

type Repository interface {
	GetAllTimesheet() ([]entity.Timesheet, error)
	GetTimesheetByID(ID int) (entity.Timesheet, error)
	AddTimesheet(timesheet entity.Timesheet) (entity.Timesheet, error)
	UpdateTimesheet(timesheet entity.Timesheet) (entity.Timesheet, error)
	DeleteTimesheet(timesheet entity.Timesheet) (entity.Timesheet, error)
}

type repository struct {
	db *gorm.DB
}

func NewTimesheetRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllTimesheet() ([]entity.Timesheet, error) {
	var timesheets []entity.Timesheet

	err := r.db.Find(&timesheets).Error

	return timesheets, err
}

func (r *repository) GetTimesheetByID(ID int) (entity.Timesheet, error) {
	var timesheet entity.Timesheet

	err := r.db.First(&timesheet, ID).Error

	return timesheet, err
}

func (r *repository) AddTimesheet(timesheet entity.Timesheet) (entity.Timesheet, error) {
	err := r.db.Create(&timesheet).Error
	return timesheet, err
}

func (r *repository) UpdateTimesheet(timesheet entity.Timesheet) (entity.Timesheet, error) {
	err := r.db.Save(&timesheet).Error
	return timesheet, err
}

func (r *repository) DeleteTimesheet(timesheet entity.Timesheet) (entity.Timesheet, error) {
	err := r.db.Delete(&timesheet).Error
	return timesheet, err
}
