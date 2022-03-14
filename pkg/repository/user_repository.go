package repository

import (
	"be-timesheet/pkg/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]entity.User, error)
	GetUserByID(ID int) (*entity.User, error)
	AddUser(user *entity.User) (*entity.User, error)
	UpdateUser(user *entity.User) (*entity.User, error)
	DeleteUser(ID int) error
	GetUserByEmailPassword(login entity.LoginRequest) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAllUsers() ([]entity.User, error) {
	var users []entity.User

	err := r.db.Order("id desc").Find(&users).Error

	return users, err
}

func (r *userRepository) GetUserByID(ID int) (*entity.User, error) {
	var user entity.User

	err := r.db.Where("id = ?", ID).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) AddUser(user *entity.User) (*entity.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return user, err
}

func (r *userRepository) UpdateUser(user *entity.User) (*entity.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) DeleteUser(ID int) error {
	// err := r.db.Delete(&user).Error
	// return user, err
	var user entity.User
	err := r.db.Where("id = ?", ID).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetUserByEmailPassword(login entity.LoginRequest) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", login.Email).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
