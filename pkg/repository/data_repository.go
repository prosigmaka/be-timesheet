package repository

import (
	"be-timesheet/pkg/entity"

	"gorm.io/gorm"
)

type DataRespository interface {
	GetAllData() ([]entity.Data, error)
	GetDataByID(ID int) (entity.Data, error)
	AddData(data entity.Data) (entity.Data, error)
	UpdateData(data entity.Data) (entity.Data, error)
	DeleteData(ID int) error
}

type dataRepository struct {
	db *gorm.DB
}

func NewDataRepository(db *gorm.DB) *dataRepository {
	return &dataRepository{db}
}

func (r *dataRepository) GetAllData() ([]entity.Data, error) {
	var datas []entity.Data

	err := r.db.Order("id asc").Find(&datas).Error

	if err != nil {
		return nil, err
	}
	return datas, nil
}

func (r *dataRepository) GetDataByID(ID int) (entity.Data, error) {
	var data entity.Data

	err := r.db.Where("id = ?", ID).Take(&data).Error

	return data, err
}

func (r *dataRepository) AddData(data entity.Data) (entity.Data, error) {
	err := r.db.Create(&data).Error
	return data, err
}

func (r *dataRepository) UpdateData(data entity.Data) (entity.Data, error) {
	err := r.db.Save(&data).Error
	return data, err
}

func (r *dataRepository) DeleteData(ID int) error {
	var data entity.Data
	err := r.db.Where("id=?",ID).Delete(&data).Error
	if err != nil {
		return err
	}
	return nil
	
}