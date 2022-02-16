package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"timesheet-be/pkg/entity"
	"timesheet-be/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type projectHandler struct {
	projectService service.Service
}

func NewProjectHandler(projectService service.Service) *projectHandler {
	return &projectHandler{projectService}
}

func (h *projectHandler) GetAllProject(c *gin.Context) {
	projects, err := h.projectService.GetAllProject()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var projectsResponse []entity.ProjectResponse
	for _, project := range projects {
		projectResponse := entity.ProjectResponse{
			ID: project.ID,
			ProjectName: project.ProjectName,
			PlacementAddress: project.PlacementAddress,
			StartDate: project.StartDate,
			EndDate: project.EndDate,
		}
		projectsResponse = append(projectsResponse, projectResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": projectsResponse,
	})
}

func (h *projectHandler) GetAllProjectByID(c *gin.Context) {
	idStr := c.Param("id_project")
	id, _ := strconv.Atoi(idStr)
	project, err := h.projectService.GetProjectByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	projectResponse := entity.ProjectResponse{
		ID: project.ID,
		ProjectName: project.ProjectName,
		PlacementAddress: project.PlacementAddress,
		StartDate: project.StartDate,
		EndDate: project.EndDate,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": projectResponse,
	})
}


func (h *projectHandler) AddProject(c *gin.Context) {
	var projectInput entity.ProjectRequest

	err := c.ShouldBindJSON(&projectInput)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	project, err := h.projectService.AddProject(projectInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": project,
	})

}


func (h *projectHandler) UpdatedProject(c *gin.Context) {
	var projectInput entity.ProjectRequest

	err := c.ShouldBindJSON(&projectInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	idStr := c.Param("id_project")
	id, _ := strconv.Atoi(idStr)

	project, err := h.projectService.UpdateProject(id, projectInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": project,
	})

}

func (h *projectHandler) DeleteProject(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	project, err := h.projectService.DeletProject(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	projectResponse := entity.ProjectResponse{
		ID: project.ID,
		ProjectName: project.ProjectName,
		PlacementAddress: project.PlacementAddress,
		StartDate: project.StartDate,
		EndDate: project.EndDate,
	}

	

	c.JSON(http.StatusOK, gin.H{
		"data": projectResponse,
	})

}