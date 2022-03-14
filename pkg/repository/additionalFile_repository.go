package repository

import (
	"be-timesheet/pkg/entity"

	"gorm.io/gorm"
)

type FileRepository interface {
	GetAllFiles() ([]entity.AdditionalFile, error)
	GetFileByID(ID int) (entity.AdditionalFile, error)
	AddNewFile(file *entity.AdditionalFile) (*entity.AdditionalFile, error)
	UpdateFile(file *entity.AdditionalFile) (*entity.AdditionalFile, error)
	DeleteFile(ID int) error
}

type fileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) *fileRepository {
	return &fileRepository{db}
}

func (r *fileRepository) GetAllFiles() ([]entity.AdditionalFile, error) {
	var files []entity.AdditionalFile

	err := r.db.Order("id asc").Find(&files).Error

	if err != nil {
		return nil, err
	}
	return files, nil
}

func (r *fileRepository) GetFileByID(ID int) (entity.AdditionalFile, error) {
	var file entity.AdditionalFile

	err := r.db.Where("id = ?", ID).Take(&file).Error

	return file, err
}

func (r *fileRepository) AddNewFile(file *entity.AdditionalFile) (*entity.AdditionalFile, error) {
	err := r.db.Create(&file).Error
	return file, err
}

func (r *fileRepository) UpdateFile(file *entity.AdditionalFile) (*entity.AdditionalFile, error) {
	err := r.db.Save(&file).Error
	return file, err
}

func (r *fileRepository) DeleteFile(ID int) error {
	var file entity.AdditionalFile
	err := r.db.Where("id = ?", ID).Delete(&file).Error
	if err != nil {
		return err
	}
	return nil
}
