package controller_gin

import (
	"fmt"
	"my-app-tx/service"
	rc "my-app-tx/utils/http"
	"my-app-tx/utils/logger"
	"my-app-tx/utils/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Listar todas las transacciones
func GetTransactions(c *gin.Context) {
	ctx := c.Request.Context()
	log := logger.WithContext(ctx).Sugar()
	tenant := c.GetHeader("x-tenant-id")
	var response rc.ApiResponse

	txService, ok := service.GetTrxService(tenant)
	if !ok {
		c.JSON(http.StatusBadRequest, rc.ErrorBad("TenantId invalido", http.StatusBadRequest))
		log.Errorf("TenantId invalido: %v", c.Request.Header)
		return
	}
	// Obtener todas las transacciones usando el servicio
	success, response, error := txService.GetAll(ctx)
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
	ctx := c.Request.Context()
	log := logger.WithContext(ctx).Sugar()
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
		log.Error("Error decodificando la request")
		c.JSON(http.StatusNotFound, rc.ErrorBad(fmt.Sprintf("Error en el formato JSON: %v", err), http.StatusBadRequest))
		return
	}

	// Insertar la transacción usando el servicio
	insert, response, err := txService.AddTrx(&t, ctx)
	if err != nil {
		log.Error("Error al insertar la transacción")
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
