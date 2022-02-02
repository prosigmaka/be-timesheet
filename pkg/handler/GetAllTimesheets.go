package handler

import (
	"be-timesheet/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h handler) GetAllTimesheets(w http.ResponseWriter, r *http.Request) {
	var timesheet []models.Timesheet

	if result := h.DB.Find(&timesheet); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(timesheet)
}
