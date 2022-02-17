package handler

import (
	"be-timesheet/pkg/entity"
	"be-timesheet/pkg/jwttoken"
	"be-timesheet/pkg/service"

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
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err,
		})
		return
	}

	registerUserError := registerUser.Validate()
	if len(registerUserError) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": registerUserError,
		})
		return
	}

	result, err := h.userService.AddUser(&registerUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (h *userHandler) GetAllUsers(c *gin.Context) {
	result, err := h.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err,
		})
		return
	}

	if result == nil {
		result = &[]entity.UserResponse{}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (h *userHandler) GetUserByID(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	result, err := h.userService.GetUserByID(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err,
		})
		return
	}

	if result == nil {
		result = &entity.UserResponse{}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var updateUser entity.User
	err = c.ShouldBindJSON(&updateUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err,
		})
		return
	}

	updateUser.ID = userId

	updateUserError := updateUser.Validate()
	if len(updateUserError) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": updateUserError,
		})
		return
	}

	result, err := h.userService.UpdateUser(&updateUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err,
		})
		return
	}

	if result == nil {
		result = &entity.UserResponse{}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	err = h.userService.DeleteUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "Succesfully Deleted User",
	})
}

func (h *userHandler) Login(c *gin.Context) {
	var loginRequest entity.LoginRequest

	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": err,
		})
		return
	}

	validateUser, err := h.userService.GetUserByEmailPassword(loginRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err,
		})
		return
	}

	if validateUser == nil {
		validateUser = &entity.User{}
	}

	// Generate JWT
	token, err := jwttoken.CreateToken(int64(validateUser.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err,
		})
		return
	}

	userData := map[string]interface{}{
		"access_token": token.AccessToken,
		"expired":      token.ExpiredToken,
		"user_id":      validateUser.ID,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userData,
	})
}
