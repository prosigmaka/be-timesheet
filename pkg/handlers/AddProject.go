package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"timesheet-app/pkg/models"
)

func (h handler) AddProject(w http.ResponseWriter, r *http.Request) {
	//read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var project models.Project
	json.Unmarshal(body, &project)


	//append  to the project mocks
	if result := h.DB.Create(&project); result.Error != nil {
		fmt.Println(result.Error)
	}

	// project.Id = rand.Intn(100)
	// mocks.Projects = append(mocks.Projects, project)

	//send a 201 created reposnse 
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("created")
}