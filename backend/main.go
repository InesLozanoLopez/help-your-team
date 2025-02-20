package main

import (
	"backend/api/tasks"
	"backend/internal/config"
	"backend/internal/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config.LoadEnvVariables()

	dbConfig := config.NewDBConfig()
	db, err := dbConfig.SetupDB()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	err = r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatal(err)
	}
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	taskService := services.NewTasksService(db)
	tasks.RegisterRoutes(r.Group("api/v1"), taskService)

	r.Run(":8080")
}
