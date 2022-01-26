package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"timesheet-app/pkg/models"

	"github.com/gorilla/mux"
)

func (h handler) DeletProject(w http.ResponseWriter, r *http.Request) {
	// // read the dynamic id parameter
	 vars := mux.Vars(r)
	 id, _ := strconv.Atoi(vars["id"])

	// // iterate over all the mock project
	// for index, project := range mocks.Projects {
	// 	if project.Id == id {
	// 		// delet book and send response if the book id matches dynamic Id
	// 		mocks.Projects = append(mocks.Projects[:index], mocks.Projects[index+1:]...)
	// 		w.Header().Add("content-type", "application/json")
	// 		w.WriteHeader(http.StatusOK)
	// 		json.NewEncoder(w).Encode("deleted")
	// 		break
	// 	}
	// }
	
	var project models.Project 

	if result := h.DB.First(&project, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	h.DB.Delete(&project)

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("deleted")
	

	
}
