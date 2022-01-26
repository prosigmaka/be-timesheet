package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"timesheet-app/pkg/models"

	"github.com/gorilla/mux"
)

func (h handler) GetProject(w http.ResponseWriter, r *http.Request) {
	//read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//iterate over all the mock projects
	// for _, project := range mocks.Projects {
	// 	if project.Id == id{
	// 		//if ids are equal send book as response
	// 		w.WriteHeader(http.StatusOK)
	// 		w.Header().Add("content-type", "application-json")
	// 		json.NewEncoder(w).Encode(project)
	// 		break
	// 	}
	// }

	var project models.Project

	if result := h.DB.First(&project, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application-json")
	json.NewEncoder(w).Encode(project)
	
}