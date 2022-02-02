package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"be-timesheet/pkg/models"

	"github.com/gorilla/mux"
)

func (h handler) GetTimesheet(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find book by id
	var timesheet models.Timesheet

	if result := h.DB.First(&timesheet, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(timesheet)

}
