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
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
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

	statusRepository := repository.NewStatusRepository(DB)
	statusService := service.NewServiceStatus(statusRepository)
	statusHandler := handler.NewStatusHandler(statusService)

	// router := mux.NewRouter()
	router := gin.Default()
	timesheet := router.Group("/v1/timesheets")

	timesheet.GET("/", timesheetHandler.GetAllTimesheets)
	timesheet.GET("/:id_timesheet", timesheetHandler.GetTimesheetByID)
	timesheet.POST("/", timesheetHandler.AddTimesheet)
	timesheet.PUT("/:id_timesheet", timesheetHandler.UpdateTimesheet)
	timesheet.DELETE("/:id", timesheetHandler.DeleteTimesheet)

	status := router.Group("/v1/status")

	status.GET("/", statusHandler.GetAllStatus)
	status.GET("/:id_status", statusHandler.GetStatusByID)

	log.Println("API is running")
	port := fmt.Sprintf(":%d", viper.GetInt("App.Port"))
	router.Use(CORSMiddleware())
	router.Run(port)

}
