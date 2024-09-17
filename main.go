package main

import (
	"CRUD/config"
	"CRUD/controller"
	"CRUD/helper"
	"CRUD/model"
	"CRUD/repository"
	"CRUD/router"
	"CRUD/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	err := db.Table("computador").AutoMigrate(&model.Computador{})
	if err != nil {
		return
	}

	// Init Repository
	computadorRepository := repository.NewComputadorRepositoryImpl(db)

	// Init Service
	computadorService := service.NewComputadorServiceImpl(computadorRepository, validate)

	// Init controller
	computadorController := controller.NewComputadorController(computadorService)

	// Router
	routes := router.NewRouter(computadorController)

	// Adicionando middleware de CORS
	routes.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Configuração do servidor HTTP
	server := &http.Server{
		Addr:           ":2511",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Iniciar o servidor
	err = server.ListenAndServe()
	helper.ErrorPanic(err)
}
