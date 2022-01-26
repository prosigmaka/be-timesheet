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
	

	vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    // Read request body
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
    }

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

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Updated")


}