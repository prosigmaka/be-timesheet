package service

import (
	"be-timesheet/pkg/entity"
	"be-timesheet/pkg/repository"
	"be-timesheet/pkg/security"

	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAllUsers() (*[]entity.UserResponse, error)
	GetUserByID(ID int) (*entity.UserResponse, error)
	AddUser(registerRequest *entity.RegisterRequest) (*entity.UserResponse, error)
	UpdateUser(user *entity.User) (*entity.UserResponse, error)
	DeleteUser(ID int) error
	GetUserByEmailPassword(loginRequest entity.LoginRequest) (*entity.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewServiceUser(userRepo repository.UserRepository) *userService {
	return &userService{userRepo}
}

func (s *userService) GetAllUsers() (*[]entity.UserResponse, error) {
	result, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var users []entity.UserResponse
	for _, item := range result {
		var user entity.UserResponse
		user.ID = item.ID
		user.Email = item.Email
		user.FullName = fmt.Sprintf("%s %s", item.FirstName, item.LastName)
		user.RoleID = item.RoleID
		users = append(users, user)
	}

	return &users, nil
}

func (s *userService) GetUserByID(ID int) (*entity.UserResponse, error) {
	// user, err := s.userRepo.GetUserByID(ID)
	// return user, err
	var userResponse entity.UserResponse

	result, err := s.userRepo.GetUserByID(ID)
	if err != nil {
		return nil, err
	}

	if result != nil {
		userResponse = entity.UserResponse{
			ID:       result.ID,
			FullName: fmt.Sprintf("%s %s", result.FirstName, result.LastName),
			Email:    result.Email,
			RoleID:   result.RoleID,
		}
	}

	return &userResponse, nil
}

func (s *userService) AddUser(registerRequest *entity.RegisterRequest) (*entity.UserResponse, error) {
	// user := entity.User{
	// 	Email: userRequest.Email,
	// 	// Password: userRequest.Password,
	// }

	// password, err := user.EncryptPassword(userRequest.Password)
	// user.Password = password

	// newUser, err := s.userRepo.AddUser(&user)
	// return newUser, err
	var user = entity.User{
		FirstName: registerRequest.FirstName,
		LastName:  registerRequest.LastName,
		RoleID:    registerRequest.RoleID,
		Email:     registerRequest.Email,
	}

	password, err := user.EncryptPassword(registerRequest.Password)
	if err != nil {
		return nil, err
	}

	user.Password = password

	result, err := s.userRepo.AddUser(&user)
	if err != nil {
		return nil, err
	}

	var afterRegister entity.UserResponse

	if result != nil {
		afterRegister = entity.UserResponse{
			ID:       result.ID,
			FullName: fmt.Sprintf("%s %s", result.FirstName, result.LastName),
			Email:    result.Email,
			RoleID:   result.RoleID,
		}
	}

	return &afterRegister, nil
}

func (s *userService) UpdateUser(user *entity.User) (*entity.UserResponse, error) {
	// user, err := s.userRepo.GetUserByID(ID)

	// user.Email = userRequest.Email

	// password, err := user.EncryptPassword(userRequest.Password)

	// user.Password = password

	// updatedUser, err := s.userRepo.UpdateUser(user)
	// return updatedUser, err
	password, err := user.EncryptPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = password

	result, err := s.userRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	var userAfterUpdate entity.UserResponse
	userAfterUpdate = entity.UserResponse{
		ID:       result.ID,
		FullName: fmt.Sprintf("%s %s", result.FirstName, result.LastName),
		Email:    result.Email,
		RoleID:   result.RoleID,
	}
	return &userAfterUpdate, err
}

func (s *userService) DeleteUser(ID int) error {
	// user, err := s.userRepo.GetUserByID(ID)
	// deleteUser, err := s.userRepo.DeleteUser(&user)
	// return deleteUser, err
	err := s.userRepo.DeleteUser(ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) GetUserByEmailPassword(loginRequest entity.LoginRequest) (*entity.User, error) {
	result, err := s.userRepo.GetUserByEmailPassword(loginRequest)
	if err != nil {
		return nil, err
	}

	err = security.VerifyPassword(result.Password, loginRequest.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, fmt.Errorf("incorrect password. error %s", err.Error())
	}

	return result, nil
}
