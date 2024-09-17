package main

import (
	"CRUD/config"
	"CRUD/controller"
	"CRUD/helper"
	"CRUD/model"
	"CRUD/repository"
	"CRUD/router"
	"CRUD/service"
	"github.com/go-playground/validator/v10"
	"net/http"

	"time"
)

func main() {

	//Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("computador").AutoMigrate(&model.Computador{})

	//Init Repository
	computadorRepository := repository.NewComputadorRepositoryImpl(db)

	//Init Service
	computadorService := service.NewComputadorServiceImpl(computadorRepository, validate)

	//Init controller
	computadorController := controller.NewComputadorController(computadorService)

	//Router
	routes := router.NewRouter(computadorController)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}
