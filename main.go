package main

import (
	"fmt"
	"log"
	"timesheet-be/pkg/config"
	"timesheet-be/pkg/db"
	"timesheet-be/pkg/handler"
	"timesheet-be/pkg/repository"
	"timesheet-be/pkg/service"

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
	projectRepository := repository.NewProjectRepository(DB)
	projectService := service.NewServiceProject(projectRepository)
	projectHandler := handler.NewProjectHandler(projectService)

	router := gin.Default()
	project := router.Group("/v1/projects")

	project.GET("/", projectHandler.GetAllProject)
	project.GET("/:id_project", projectHandler.GetAllProjectByID)
	project.POST("/", projectHandler.AddProject)
	project.PUT("/:id_project", projectHandler.UpdatedProject)
	project.DELETE("/:id", projectHandler.DeleteProject)

	log.Println("API is running")
	port := fmt.Sprintf(":%d", viper.GetInt("App.Port"))
	router.Use(CORSMiddleware())
	router.Run(port)


}
