package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/CircuitBreaker - Federico Salas/src/api/services"
	"net/http"
) //si hay dos dependencias que se llamen igual agrego otro nombre adelante

func GetSites(context *gin.Context) {
	sites, apiErr := services.GetSitesS()
	if apiErr != nil {
		context.JSON(apiErr.Status, apiErr)
		return
	}
	context.JSON(http.StatusOK, sites)
}
