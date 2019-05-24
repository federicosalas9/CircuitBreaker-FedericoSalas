package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/CircuitBreaker-FedericoSalas/src/api/controllers"
)

func StartApp() {
	router := gin.Default()
	router.GET("/ping", controllers.Ping)
	router.GET("/sites", controllers.GetSites)
	router.Run(":8080")
}
