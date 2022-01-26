package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"timesheet-app/pkg/models"
)

func (h handler) GetAllProject(w http.ResponseWriter, r *http.Request) {
	// w.Header().Add("content-type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(mocks.Projects)

	var projects []models.Project

	if result := h.DB.Find(&projects); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(projects)
}