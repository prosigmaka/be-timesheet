package handler

import (
	"be-timesheet/pkg/entity"
	"be-timesheet/pkg/jwttoken"
	"be-timesheet/pkg/response"
	"be-timesheet/pkg/service"
	"errors"

	// "fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var registerUser entity.RegisterRequest

	err := c.ShouldBindJSON(&registerUser)

	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	registerUserError := registerUser.Validate()
	if len(registerUserError) > 0 {
		response.ResponseCustomError(c, registerUserError, http.StatusBadRequest)
		return
	}

	result, err := h.userService.AddUser(&registerUser)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseCreated(c, result)
}

func (h *userHandler) GetAllUsers(c *gin.Context) {
	result, err := h.userService.GetAllUsers()
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	if result == nil {
		result = &[]entity.UserResponse{}
	}

	response.ResponseOKWithData(c, result)
}

func (h *userHandler) GetUserByID(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.userService.GetUserByID(userId)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	if result == nil {
		result = &entity.UserResponse{}
	}

	response.ResponseOKWithData(c, result)
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		response.ResponseError(c, errors.New("invalid User ID").Error(), http.StatusBadRequest)
		return
	}

	var updateUser entity.User
	err = c.ShouldBindJSON(&updateUser)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	updateUser.ID = userId

	updateUserError := updateUser.Validate()
	if len(updateUserError) > 0 {
		response.ResponseCustomError(c, updateUserError, http.StatusBadRequest)
		return
	}

	result, err := h.userService.UpdateUser(&updateUser)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	if result == nil {
		result = &entity.UserResponse{}
	}

	response.ResponseOKWithData(c, result)
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		response.ResponseError(c, errors.New("invalid User ID").Error(), http.StatusBadRequest)
		return
	}

	err = h.userService.DeleteUser(userId)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseOK(c, "Succesfully Deleted User")
}

func (h *userHandler) Login(c *gin.Context) {
	var loginRequest entity.LoginRequest

	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	validateUser, err := h.userService.GetUserByEmailPassword(loginRequest)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	if validateUser == nil {
		validateUser = &entity.User{}
	}

	// Generate JWT
	token, err := jwttoken.CreateToken(int64(validateUser.ID))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	userData := map[string]interface{}{
		"access_token": token.AccessToken,
		"expired":      token.ExpiredToken,
		"user_id":      validateUser.ID,
	}

	response.ResponseOKWithData(c, userData)
}
