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

type timesheetHandler struct {
	timesheetService service.Service
}

func NewTimesheetHandler(timesheetService service.Service) *timesheetHandler {
	return &timesheetHandler{timesheetService}
}

func (h *timesheetHandler) GetAllTimesheets(c *gin.Context) {
	result, err := h.timesheetService.GetAllTimesheet()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err,
		})
		return
	}

	if result == nil {
		result = []entity.TimesheetResponse{}

	}

	// var timesheetsResponse []entity.TimesheetResponse
	// for _, timesheet := range timesheets {
	// 	timesheetResponse := entity.TimesheetResponse{
	// 		ID:            timesheet.ID,
	// 		Date:          timesheet.Date,
	// 		UserID:        timesheet.UserID,
	// 		WorkingStart:  timesheet.WorkingStart,
	// 		WorkingEnd:    timesheet.WorkingEnd,
	// 		OvertimeStart: timesheet.OvertimeStart,
	// 		OvertimeEnd:   timesheet.OvertimeEnd,
	// 		Activity:      timesheet.Activity,
	// 		ProjectID:     timesheet.ProjectID,
	// 		StatusID:      timesheet.StatusID,
	// 	}
	// 	// timesheetResponse := convertToTimesheetResponse(timesheet)
	// 	timesheetsResponse = append(timesheetsResponse, timesheetResponse)
	// }

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (h *timesheetHandler) GetTimesheetByID(c *gin.Context) {
	idStr := c.Param("id_timesheet")
	id, _ := strconv.Atoi(idStr)

	result, err := h.timesheetService.GetTimesheetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err,
		})
		return
	}

	// timesheetResponse := entity.TimesheetResponse{
	// 	ID:            timesheet.ID,
	// 	Date:          timesheet.Date,
	// 	UserID:        timesheet.UserID,
	// 	WorkingStart:  timesheet.WorkingStart,
	// 	WorkingEnd:    timesheet.WorkingEnd,
	// 	OvertimeStart: timesheet.OvertimeStart,
	// 	OvertimeEnd:   timesheet.OvertimeEnd,
	// 	Activity:      timesheet.Activity,
	// 	ProjectID:     timesheet.ProjectID,
	// 	StatusID:      timesheet.StatusID,
	// }

	// timesheetResponse := convertToTimesheetResponse(timesheet)

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})

}

func (h *timesheetHandler) AddTimesheet(c *gin.Context) {
	// var timesheetInput entity.TimesheetRequest
	date := c.DefaultPostForm("date", "date")
	workingstart := c.DefaultPostForm("working_start", "working_start")
	workingend := c.DefaultPostForm("working_end", "working_end")
	overtimestart := c.DefaultPostForm("overtime_start", "overtime_start")
	overtimeend := c.DefaultPostForm("overtime_end", "overtime_end")
	activity := c.DefaultPostForm("activity", "activity")
	projectid, _ := strconv.Atoi(c.DefaultPostForm("project_id", "project_id"))
	statusid, _ := strconv.Atoi(c.DefaultPostForm("status_id", "status_id"))

	tokenMetadata, err := jwttoken.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	userId := tokenMetadata.UserID

	var timesheetInput = entity.TimesheetResponse{
		Date:          date,
		UserID:        int(userId),
		WorkingStart:  workingstart,
		WorkingEnd:    workingend,
		OvertimeStart: overtimestart,
		OvertimeEnd:   overtimeend,
		Activity:      activity,
		ProjectID:     projectid,
		StatusID:      statusid,
	}

	// err := c.ShouldBindJSON(&timesheetInput)
	// if err != nil {
	// 	errorMessages := []string{}
	// 	for _, e := range err.(validator.ValidationErrors) {
	// 		errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
	// 		errorMessages = append(errorMessages, errorMessage)
	// 	}

	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": errorMessages,
	// 	})
	// 	return
	// }

	result, err := h.timesheetService.AddTimesheet(&timesheetInput)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})

}

func (h *timesheetHandler) UpdateTimesheet(c *gin.Context) {
	// var timesheetInput entity.TimesheetRequest

	// err := c.ShouldBindJSON(&timesheetInput)
	// if err != nil {
	// 	// errorMessages := []string{}
	// 	// for _, e := range err.(validator.ValidationErrors) {
	// 	// 	errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
	// 	// 	errorMessages = append(errorMessages, errorMessage)
	// 	// }

	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": err,
	// 	})
	// 	return
	// }

	// userId, _ := strconv.Atoi(c.DefaultPostForm("user_id", "user_id"))
	date := c.DefaultPostForm("date", "date")
	workingstart := c.DefaultPostForm("working_start", "working_start")
	workingend := c.DefaultPostForm("working_end", "working_end")
	overtimestart := c.DefaultPostForm("overtime_start", "overtime_start")
	overtimeend := c.DefaultPostForm("overtime_end", "overtime_end")
	activity := c.DefaultPostForm("activity", "activity")
	projectid, _ := strconv.Atoi(c.DefaultPostForm("project_id", "project_id"))
	statusid, _ := strconv.Atoi(c.DefaultPostForm("status_id", "status_id"))

	idStr := c.Param("id_timesheet")
	id, _ := strconv.Atoi(idStr)

	var timesheet = entity.TimesheetResponse{
		ID:            id,
		Date:          date,
		// UserID:        int(userId),
		WorkingStart:  workingstart,
		WorkingEnd:    workingend,
		OvertimeStart: overtimestart,
		OvertimeEnd:   overtimeend,
		Activity:      activity,
		ProjectID:     projectid,
		StatusID:      statusid,
	}

	result, err := h.timesheetService.UpdateTimesheet(&timesheet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})

}

func (h *timesheetHandler) DeleteTimesheet(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	err := h.timesheetService.DeleteTimesheet(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"errors": err,
	// 	})
	// 	return
	// }

	// timesheetResponse := entity.TimesheetResponse{
	// 	ID:            timesheet.ID,
	// 	Date:          timesheet.Date,
	// 	UserID:        timesheet.UserID,
	// 	WorkingStart:  timesheet.WorkingStart,
	// 	WorkingEnd:    timesheet.WorkingEnd,
	// 	OvertimeStart: timesheet.OvertimeStart,
	// 	OvertimeEnd:   timesheet.OvertimeEnd,
	// 	Activity:      timesheet.Activity,
	// 	ProjectID:     timesheet.ProjectID,
	// 	StatusID:      timesheet.StatusID,
	// }

	// timesheetResponse := convertToTimesheetResponse(timesheet)

	c.JSON(http.StatusOK, gin.H{
		"data": "Succesfully deleted timesheet",
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
