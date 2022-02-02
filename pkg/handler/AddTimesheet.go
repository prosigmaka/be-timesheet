package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	// "math/rand"
	"net/http"

	// "timesheet-be/pkg/mocks"
	"be-timesheet/pkg/models"
)

func (h handler) AddTimesheet(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var timesheet models.Timesheet
	json.Unmarshal(body, &timesheet) //decode json, &book berarti data hasil ditampung di book.go (folder models)

	// Append to the book mocks
	// timesheet.ID = rand.Intn(100)
	// mocks.Timesheets = append(mocks.Timesheets, timesheet)
	if result := h.DB.Create(&timesheet); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Send a 201 created response
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")
}
