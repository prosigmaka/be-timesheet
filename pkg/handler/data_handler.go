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

type dataHandler struct {
	dataService service.DataService
}

func NewDataHandler(dataService service.DataService) *dataHandler {
	return &dataHandler{dataService}
}

func (h *dataHandler) GetAllData(c *gin.Context) {
	datas, err := h.dataService.GetAllData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var datasResponse []entity.DataResponse
	for _, data := range datas {
		dataResponse := entity.DataResponse{
			ID: data.ID,
			UserID: data.UserID,
			ProjectID: data.ProjectID,
			TimesheetID: data.TimesheetID,
			Employee_name: data.Employee_name,
			Project_name: data.Project_name,
			StartDate: data.StartDate,
			EndDate: data.EndDate,
		}
		datasResponse = append(datasResponse, dataResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": datasResponse,
	})
	
}

func (h *dataHandler) GetDataByID(c *gin.Context) {
	idStr := c.Param("id_data")
	id, _ := strconv.Atoi(idStr)
	data, err := h.dataService.GetDataByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	dataResponse := entity.DataResponse{
		ID: data.ID,
		UserID: data.UserID,
		ProjectID: data.ProjectID,
		TimesheetID: data.TimesheetID,
		Employee_name: data.Employee_name,
		Project_name: data.Project_name,
		StartDate: data.StartDate,
		EndDate: data.EndDate,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": dataResponse,
	})
}

func (h *dataHandler) AddData(c *gin.Context) {
	var dataInput entity.DataRequest

	tokenMetadata, _ := jwttoken.ExtractTokenMetadata(c.Request)

	err := c.ShouldBindJSON(&dataInput)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	} else {
		userId := tokenMetadata.UserID
		dataInput.UserID = int(userId)
	}

	data, err := h.dataService.AddData(dataInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (h *dataHandler) UpdateData(c *gin.Context) {
	var dataInput entity.DataRequest

	tokenMetadata, _ := jwttoken.ExtractTokenMetadata(c.Request)

	err := c.ShouldBindJSON(&dataInput)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	} else {
		userId := tokenMetadata.UserID
		dataInput.UserID = int(userId)
	}

	idStr := c.Param("id_data")
	id, _ := strconv.Atoi(idStr)

	data, err := h.dataService.UpdateData(id, dataInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}

func (h *dataHandler) DeleteData(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.dataService.DeleteData(id)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseOK(c, "Succesfully Deleted Data")
}