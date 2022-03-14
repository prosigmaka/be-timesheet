package main

import (
	"be-timesheet/pkg/api/middleware"
	"be-timesheet/pkg/config"
	"be-timesheet/pkg/db"
	"be-timesheet/pkg/handler"
	"be-timesheet/pkg/repository"
	"be-timesheet/pkg/service"
	"fmt"

	// "log"
	"io"
	"os"
	"time"

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

	addFileRepository := repository.NewFileRepository(DB)
	addFileService := service.NewServiceAdditionalFile(addFileRepository)
	addFileHandler := handler.NewAddFileHandler(addFileService)

	statusRepository := repository.NewStatusRepository(DB)
	statusService := service.NewServiceStatus(statusRepository)
	statusHandler := handler.NewStatusHandler(statusService)

	userRepository := repository.NewUserRepository(DB)
	userService := service.NewServiceUser(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()
	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// router := mux.NewRouter()
	router := gin.Default()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	timesheet := router.Group("/v1/timesheets")

	timesheet.GET("/", middleware.AuthMiddleware(), timesheetHandler.GetAllTimesheets)
	timesheet.GET("/:id_timesheet", middleware.AuthMiddleware(), timesheetHandler.GetTimesheetByID)
	timesheet.POST("/", middleware.AuthMiddleware(), timesheetHandler.AddTimesheet)
	timesheet.PUT("/:id_timesheet", middleware.AuthMiddleware(), timesheetHandler.UpdateTimesheet)
	timesheet.DELETE("/:id", middleware.AuthMiddleware(), timesheetHandler.DeleteTimesheet)

	status := router.Group("/v1/status")

	status.GET("/", statusHandler.GetAllStatus)
	status.GET("/:id_status", statusHandler.GetStatusByID)

	addFile := router.Group("/v1/addfile")

	addFile.GET("/", middleware.AuthMiddleware(), timesheetHandler.GetAllTimesheets)
	addFile.GET("/:id_file", middleware.AuthMiddleware(), timesheetHandler.GetTimesheetByID)
	addFile.POST("/", middleware.AuthMiddleware(), timesheetHandler.AddTimesheet)
	addFile.PUT("/:id_file", middleware.AuthMiddleware(), timesheetHandler.UpdateTimesheet)
	addFile.DELETE("/:id", middleware.AuthMiddleware(), timesheetHandler.DeleteTimesheet)

	user := router.Group("v1/user")

	user.GET("/", userHandler.GetAllUsers)
	user.GET("/:user_id", middleware.AuthMiddleware(), userHandler.GetUserByID)
	user.POST("/", userHandler.RegisterUser)
	user.PUT("/:user_id", userHandler.UpdateUser)
	user.DELETE("/:user_id", userHandler.DeleteUser)
	user.POST("/login", userHandler.Login)

	file := router.Group("v1/file")

	file.GET("/", addFileHandler.GetAllFiles)
	file.GET("/:file_id", middleware.AuthMiddleware(), addFileHandler.GetFileByID)
	file.POST("/", middleware.AuthMiddleware(), addFileHandler.AddNewFile)
	file.PUT("/:file_id", middleware.AuthMiddleware(), addFileHandler.UpdateFile)
	file.DELETE("/:file_id", middleware.AuthMiddleware(), addFileHandler.DeleteFile)
	// file.GET("/:name", addFileHandler.DownloadFile)

	// log.Println("API is running")

	port := fmt.Sprintf(":%d", viper.GetInt("App.Port"))
	router.Use(middleware.CORSMiddleware())
	router.Run(port)

}

// router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

// 	// your custom format
// 	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
// 		param.ClientIP,
// 		param.TimeStamp.Format(time.RFC1123),
// 		param.Method,
// 		param.Path,
// 		param.Request.Proto,
// 		param.StatusCode,
// 		param.Latency,
// 		param.Request.UserAgent(),
// 		param.ErrorMessage,
// 	)
// }))

// Disable Console Color, you don't need console color when writing the logs to file.
// gin.DisableConsoleColor()

// // Logging to a file.
// f, _ := os.Create("gin.log")
// gin.DefaultWriter = io.MultiWriter(f)
