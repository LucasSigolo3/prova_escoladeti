package router

import (
	"CRUD/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(computadorController *controller.ComputadorController) *gin.Engine {
	service := gin.Default()

	// Configuração básica do CORS
	service.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // Permitir todas as origens
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},          // Permitir métodos específicos
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Permitir cabeçalhos específicos
		AllowCredentials: true,
	}))

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	computadorRouter := router.Group("/computador")
	computadorRouter.GET("", computadorController.FindAll)
	computadorRouter.GET("/:computadorId", computadorController.FindById)
	computadorRouter.POST("", computadorController.Create)
	computadorRouter.PATCH("/:computadorId", computadorController.Update)
	computadorRouter.DELETE("/:computadorId", computadorController.Delete)

	return service
}
