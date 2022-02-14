package handler

import (
	"be-timesheet/pkg/entity"
	"be-timesheet/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type statusHandler struct {
	statusService service.StatusService
}

func NewStatusHandler(statusService service.StatusService) *statusHandler {
	return &statusHandler{statusService}
}

func (h *statusHandler) GetAllStatus(c *gin.Context) {
	status, err := h.statusService.GetAllStatus()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var statusResponses []entity.StatusResponse
	for _, option := range status {
		statusResponse := entity.StatusResponse{
			ID:         option.ID,
			StatusName: option.StatusName,
		}

		statusResponses = append(statusResponses, statusResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": statusResponses,
	})
}

func (h *statusHandler) GetStatusByID(c *gin.Context) {
	idStr := c.Param("id_status")
	id, _ := strconv.Atoi(idStr)
	status, err := h.statusService.GetStatusByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	statusResponse := entity.StatusResponse{
		ID:         status.ID,
		StatusName: status.StatusName,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": statusResponse,
	})

}
