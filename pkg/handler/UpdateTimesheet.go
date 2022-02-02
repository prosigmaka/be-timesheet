package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"be-timesheet/pkg/models"

	"github.com/gorilla/mux"
)

func (h handler) UpdateTimesheet(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedTimesheet models.Timesheet
	json.Unmarshal(body, &updatedTimesheet)

	var timesheet models.Timesheet

	if result := h.DB.First(&timesheet, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Update and send response when book id matches id
	timesheet.Date = updatedTimesheet.Date
	timesheet.WorkingStart = updatedTimesheet.WorkingStart
	timesheet.WorkingEnd = updatedTimesheet.WorkingEnd
	timesheet.OvertimeStart = updatedTimesheet.OvertimeStart
	timesheet.OvertimeEnd = updatedTimesheet.OvertimeEnd
	timesheet.Activity = updatedTimesheet.Activity
	timesheet.ProjectID = updatedTimesheet.ProjectID
	timesheet.StatusID = updatedTimesheet.StatusID

	h.DB.Save(&timesheet)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")

}
