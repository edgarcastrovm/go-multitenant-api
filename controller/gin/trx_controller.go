package controller_gin

import (
	"fmt"
	"my-app-tx/service"
	rc "my-app-tx/utils/http"
	. "my-app-tx/utils/middleware"
	"my-app-tx/utils/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Listar todas las transacciones
func GetTransactions(c *gin.Context) {
	log := GetLoggerGin(c)
	tenant := c.GetHeader("x-tenant-id")
	var response rc.ApiResponse

	txService, ok := service.GetTrxService(tenant)
	if !ok {
		c.JSON(http.StatusBadRequest, rc.ErrorBad("TenantId invalido", http.StatusBadRequest))
		log.Errorf("TenantId invalido: %v", c.Request.Header)
		return
	}
	success, response, error := txService.GetAll(log)
	if !success {
		log.Error("Error: no se encontraron transacciones")
		c.JSON(http.StatusNotFound, response)
		return
	}
	if error != nil {
		log.Error("Error:", error)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	log.Info("Obteniendo transacciones")
	c.JSON(http.StatusOK, response)
}

// Agregar una nueva transacción
func CreateTransaction(c *gin.Context) {
	log := GetLoggerGin(c)
	tenant := c.GetHeader("x-tenant-id")
	var response rc.ApiResponse

	txService, ok := service.GetTrxService(tenant)
	if !ok {
		c.JSON(http.StatusBadRequest, response)
		log.Errorf("TenantId invalido: %v", c.Request.Header)
		return
	}
	var t models.Transaction
	if err := c.ShouldBindJSON(&t); err != nil {
		log.WithField("error", err.Error()).Error("Error decodificando la request")
		c.JSON(http.StatusNotFound, rc.ErrorBad(fmt.Sprintf("Error en el formato JSON: %v", err), http.StatusBadRequest))
		return
	}

	// Insertar la transacción usando el servicio
	insert, response, err := txService.AddTrx(&t, log)
	if err != nil {
		log.WithField("error", err.Error()).Error("Error al insertar la transacción")
		c.JSON(http.StatusNotFound, response)
		return
	}

	if !insert {
		log.Error("No se pudo insertar la transacción")
		c.JSON(http.StatusNotFound, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
