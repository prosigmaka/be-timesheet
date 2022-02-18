package handler

import (
	"be-timesheet/pkg/entity"
	"be-timesheet/pkg/jwttoken"
	"be-timesheet/pkg/response"
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
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	if result == nil {
		result = []entity.TimesheetResponse{}

	}

	response.ResponseOKWithData(c, result)
}

func (h *timesheetHandler) GetTimesheetByID(c *gin.Context) {
	idStr := c.Param("id_timesheet")
	id, _ := strconv.Atoi(idStr)

	result, err := h.timesheetService.GetTimesheetByID(id)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	response.ResponseOKWithData(c, result)

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
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
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

	addTimesheetError := timesheetInput.Validate("")
	if len(addTimesheetError) > 0 {
		response.ResponseCustomError(c, addTimesheetError, http.StatusBadRequest)
		return
	}

	result, err := h.timesheetService.AddTimesheet(&timesheetInput)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseCreated(c, result)
}

func (h *timesheetHandler) UpdateTimesheet(c *gin.Context) {
	idStr := c.Param("id_timesheet")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	tokenMetadata, err := jwttoken.ExtractTokenMetadata(c.Request)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	userId := tokenMetadata.UserID
	// userId, _ := strconv.Atoi(c.DefaultPostForm("user_id", "user_id"))
	date := c.DefaultPostForm("date", "date")
	workingStart := c.DefaultPostForm("working_start", "working_start")
	workingEnd := c.DefaultPostForm("working_end", "working_end")
	overtimeStart := c.DefaultPostForm("overtime_start", "overtime_start")
	overtimeEnd := c.DefaultPostForm("overtime_end", "overtime_end")
	activity := c.DefaultPostForm("activity", "activity")
	projectId, _ := strconv.Atoi(c.DefaultPostForm("project_id", "project_id"))
	statusId, _ := strconv.Atoi(c.DefaultPostForm("status_id", "status_id"))

	var timesheet = entity.TimesheetResponse{
		ID:            id,
		Date:          date,
		UserID:        int(userId),
		WorkingStart:  workingStart,
		WorkingEnd:    workingEnd,
		OvertimeStart: overtimeStart,
		OvertimeEnd:   overtimeEnd,
		Activity:      activity,
		ProjectID:     projectId,
		StatusID:      statusId,
	}

	updateTimesheetError := timesheet.Validate("")
	if len(updateTimesheetError) > 0 {
		response.ResponseCustomError(c, updateTimesheetError, http.StatusBadRequest)
		return
	}

	result, err := h.timesheetService.UpdateTimesheet(&timesheet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"data": result,
	// })
	c.JSON(http.StatusOK, result)

}

func (h *timesheetHandler) DeleteTimesheet(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.timesheetService.DeleteTimesheet(id)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseOK(c, "Succesfully Deleted Timesheet")

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
