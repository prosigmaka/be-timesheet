package main

import (
	"be-timesheet/pkg/api/middleware"
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

func main() {
	DB := db.InitDB()
	timesheetRepository := repository.NewTimesheetRepository(DB)
	timesheetService := service.NewServiceTimesheet(timesheetRepository)
	timesheetHandler := handler.NewTimesheetHandler(timesheetService)

	statusRepository := repository.NewStatusRepository(DB)
	statusService := service.NewServiceStatus(statusRepository)
	statusHandler := handler.NewStatusHandler(statusService)

	userRepository := repository.NewUserRepository(DB)
	userService := service.NewServiceUser(userRepository)
	userHandler := handler.NewUserHandler(userService)

	projectRepository := repository.NewProjectRepository(DB)
	projectService := service.NewServiceProject(projectRepository)
	projectHandler := handler.NewProjectHandler(projectService)

	dataRepository := repository.NewDataRepository(DB)
	dataService := service.NewServiceData(dataRepository)
	dataHandler := handler.NewDataHandler(dataService)

	// router := mux.NewRouter()
	router := gin.Default()
	timesheet := router.Group("/v1/timesheets")

	timesheet.GET("/", middleware.AuthMiddleware(), timesheetHandler.GetAllTimesheets)
	timesheet.GET("/:id_timesheet", middleware.AuthMiddleware(), timesheetHandler.GetTimesheetByID)
	timesheet.POST("/", middleware.AuthMiddleware(), timesheetHandler.AddTimesheet)
	timesheet.PUT("/:id_timesheet", middleware.AuthMiddleware(), timesheetHandler.UpdateTimesheet)
	timesheet.DELETE("/:id", middleware.AuthMiddleware(), timesheetHandler.DeleteTimesheet)

	status := router.Group("/v1/status")

	status.GET("/", statusHandler.GetAllStatus)
	status.GET("/:id_status", statusHandler.GetStatusByID)

	user := router.Group("v1/user")

	user.GET("/", userHandler.GetAllUsers)
	user.GET("/:user_id", middleware.AuthMiddleware(), userHandler.GetUserByID)
	user.POST("/", userHandler.RegisterUser)
	user.PUT("/:user_id", userHandler.UpdateUser)
	user.DELETE("/:user_id", userHandler.DeleteUser)
	user.POST("/login", userHandler.Login)

	project := router.Group("v1/project")

	project.GET("/", projectHandler.AddProject)
	project.GET("/:project_id", middleware.AuthMiddleware(), projectHandler.GetAllProject)
	project.PUT("/:project_id", projectHandler.UpdatedProject)
	project.DELETE("/:project_id", projectHandler.DeleteProject)
	

	data := router.Group("v1/data")
	
	data.POST("/", dataHandler.AddData)
	data.GET("/:data_id", middleware.AuthMiddleware(), dataHandler.GetAllData)
	data.PUT("/:data_id", dataHandler.UpdateData)
	data.DELETE("/:data_id", dataHandler.DeleteData)


	log.Println("API is running")
	port := fmt.Sprintf(":%d", viper.GetInt("App.Port"))
	router.Use(middleware.CORSMiddleware())
	router.Run(port)

}
