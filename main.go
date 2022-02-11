package main

import (
	"be-timesheet/pkg/config"
	"be-timesheet/pkg/db"
	"be-timesheet/pkg/handler"
	"be-timesheet/pkg/repository"
	"be-timesheet/pkg/service"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	config.GetConfig()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Method", "POST, GET, DELETE, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	DB := db.InitDB()
	timesheetRepository := repository.NewTimesheetRepository(DB)
	timesheetService := service.NewServiceTimesheet(timesheetRepository)
	timesheetHandler := handler.NewTimesheetHandler(timesheetService)

	// router := mux.NewRouter()
	router := gin.Default()
	timesheet := router.Group("/v1")

	timesheet.GET("/timesheets", CORSMiddleware(), timesheetHandler.GetAllTimesheets)
	timesheet.GET("/timesheets/:id", CORSMiddleware(), timesheetHandler.GetTimesheetByID)
	timesheet.POST("/timesheets", CORSMiddleware(), timesheetHandler.AddTimesheet)
	timesheet.PUT("/timesheets/:id", CORSMiddleware(), timesheetHandler.UpdateTimesheet)
	timesheet.DELETE("/timesheets/:id", CORSMiddleware(), timesheetHandler.DeleteTimesheet)

	log.Println("API is running")
	port := fmt.Sprintf(":%d", viper.GetInt("App.Port"))
	router.Run(port)

}