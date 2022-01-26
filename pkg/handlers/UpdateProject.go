package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"timesheet-app/pkg/models"

	"github.com/gorilla/mux"
)

func (h handler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	// read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// read requset body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	// iterate over all the moc projects
	// for index, project := range mocks.Projects {
	// 	if project.Id == id {
	// 		project.ProjectName = updatedProject.ProjectName
	// 		project.PlacemetAddress = updatedProject.PlacemetAddress
	// 		project.StartPeriode = updatedProject.StartPeriode
	// 		project.EndPeriode = updatedProject.EndPeriode

	// 		mocks.Projects[index] = project

			
	// 		w.Header().Add("content-type", "application/json")
	// 		w.WriteHeader(http.StatusOK)
	// 		json.NewEncoder(w).Encode("updated")
	// 		break

	// 	}
	// }
	// update and send response when book Id matches dynamic Id

	var updatedProject models.Project
	json.Unmarshal(body, &updatedProject)

	var project models.Project

	if result := h.DB.First(&project, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	project.ProjectName = updatedProject.ProjectName
	project.PlacemetAddress = updatedProject.PlacemetAddress
	project.StartPeriode = updatedProject.StartPeriode
	project.EndPeriode = updatedProject.EndPeriode

	h.DB.Save(&project)

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("updated")


}