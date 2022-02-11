package handler

import (
	"be-timesheet/pkg/entity"
	"be-timesheet/pkg/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type timesheetHandler struct {
	timesheetService service.Service
}

func NewTimesheetHandler(timesheetService service.Service) *timesheetHandler {
	return &timesheetHandler{timesheetService}
}

func (h *timesheetHandler) GetAllTimesheets(c *gin.Context) {
	timesheets, err := h.timesheetService.GetAllTimesheet()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var timesheetsResponse []entity.TimesheetResponse
	for _, timesheet := range timesheets {
		timesheetResponse := entity.TimesheetResponse{
			ID:            timesheet.ID,
			Date:          timesheet.Date,
			WorkingStart:  timesheet.WorkingStart,
			WorkingEnd:    timesheet.WorkingEnd,
			OvertimeStart: timesheet.OvertimeStart,
			OvertimeEnd:   timesheet.OvertimeEnd,
			Activity:      timesheet.Activity,
			ProjectID:     timesheet.ProjectID,
			StatusID:      timesheet.StatusID,
		}
		// timesheetResponse := convertToTimesheetResponse(timesheet)
		timesheetsResponse = append(timesheetsResponse, timesheetResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": timesheetsResponse,
	})
}

func (h *timesheetHandler) GetTimesheetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	timesheet, err := h.timesheetService.GetTimesheetByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	timesheetResponse := entity.TimesheetResponse{
		ID:            timesheet.ID,
		Date:          timesheet.Date,
		WorkingStart:  timesheet.WorkingStart,
		WorkingEnd:    timesheet.WorkingEnd,
		OvertimeStart: timesheet.OvertimeStart,
		OvertimeEnd:   timesheet.OvertimeEnd,
		Activity:      timesheet.Activity,
		ProjectID:     timesheet.ProjectID,
		StatusID:      timesheet.StatusID,
	}

	// timesheetResponse := convertToTimesheetResponse(timesheet)

	c.JSON(http.StatusOK, gin.H{
		"data": timesheetResponse,
	})

}

func (h *timesheetHandler) AddTimesheet(c *gin.Context) {
	var timesheetInput entity.TimesheetRequest

	err := c.ShouldBindJSON(&timesheetInput)
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

	timesheet, err := h.timesheetService.AddTimesheet(timesheetInput)

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

func (h *timesheetHandler) UpdateTimesheet(c *gin.Context) {
	var timesheetInput entity.TimesheetRequest

	err := c.ShouldBindJSON(&timesheetInput)
	if err != nil {
		// errorMessages := []string{}
		// for _, e := range err.(validator.ValidationErrors) {
		// 	errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
		// 	errorMessages = append(errorMessages, errorMessage)
		// }

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	timesheet, err := h.timesheetService.UpdateTimesheet(id, timesheetInput)

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

func (h *timesheetHandler) DeleteTimesheet(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	timesheet, err := h.timesheetService.DeleteTimesheet(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	timesheetResponse := entity.TimesheetResponse{
		ID:            timesheet.ID,
		Date:          timesheet.Date,
		WorkingStart:  timesheet.WorkingStart,
		WorkingEnd:    timesheet.WorkingEnd,
		OvertimeStart: timesheet.OvertimeStart,
		OvertimeEnd:   timesheet.OvertimeEnd,
		Activity:      timesheet.Activity,
		ProjectID:     timesheet.ProjectID,
		StatusID:      timesheet.StatusID,
	}

	// timesheetResponse := convertToTimesheetResponse(timesheet)

	c.JSON(http.StatusOK, gin.H{
		"data": timesheetResponse,
	})

}

// func convertToTimesheetResponse(timesheet entity.Timesheet) entity.TimesheetResponse {
// 	return entity.TimesheetResponse{
// 		ID:            timesheet.ID,
// 		WorkingStart:  timesheet.WorkingStart,
// 		WorkingEnd:    timesheet.WorkingEnd,
// 		OvertimeStart: timesheet.OvertimeStart,
// 		OvertimeEnd:   timesheet.OvertimeEnd,
// 		Activity:      timesheet.Activity,
// 		ProjectID:     timesheet.ProjectID,
// 		StatusID:      timesheet.StatusID,
// 	}
// }
