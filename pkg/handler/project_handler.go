package handler

import (
	"be-timesheet/pkg/entity"
	"be-timesheet/pkg/jwttoken"
	"be-timesheet/pkg/response"
	"be-timesheet/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type projectHandler struct {
	projectService service.ProjectService
}

func NewProjectHandler(projectService service.ProjectService) *projectHandler {
	return &projectHandler{projectService}
}

func (h *projectHandler) GetAllProject(c *gin.Context) {
	result, err := h.projectService.GetAllProject()
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	if result == nil {
		result = []entity.ProjectResponse{}

	}

	response.ResponseOKWithData(c, result)


}

func (h *projectHandler) GetProjectByID(c *gin.Context) {
	idStr := c.Param("project_id")
	id, _ := strconv.Atoi(idStr)

	result, err := h.projectService.GetProjectByID(id)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	response.ResponseOKWithData(c, result)
	

}

func (h *projectHandler) AddProject(c *gin.Context) {
	var projectInput entity.ProjectRequest

	tokenMetadata, _ := jwttoken.ExtractTokenMetadata(c.Request)

	err := c.ShouldBindJSON(&projectInput)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	} else {
		userId := tokenMetadata.UserID
		projectInput.UserID = int(userId)
	}

	timesheet, err := h.projectService.AddProject(projectInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": timesheet,
	})


}

func (h *projectHandler) UpdatedProject(c *gin.Context) {
	var projectInput entity.ProjectRequest

	tokenMetadata, _ := jwttoken.ExtractTokenMetadata(c.Request)

	err := c.ShouldBindJSON(&projectInput)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	} else {
		userId := tokenMetadata.UserID
		projectInput.UserID = int(userId)
	}

	idStr := c.Param("id_timesheet")
	id, _ := strconv.Atoi(idStr)

	project, err := h.projectService.UpdateProject(id,projectInput)

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
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.projectService.DeletProject(id)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseOK(c, "Succesfully Deleted Project")
}