package controller_mux

import (
	"encoding/json"
	"fmt"
	"my-app-tx/service"
	rc "my-app-tx/utils/http"
	. "my-app-tx/utils/middleware"
	"my-app-tx/utils/models"
	"net/http"
)

// Listar todas las transacciones
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	log := GetLogger(r)
	tenant := r.Header.Get("x-tenant-id")
	var response rc.ApiResponse

	txService, ok := service.GetTrxService(tenant)
	if !ok {
		json.NewEncoder(w).Encode(rc.ErrorBad("TenantId invalido", http.StatusBadRequest))
		log.Error("TenantId invalido:", r.Header)
		return
	}
	success, response, error := txService.GetAll(log)
	if !success {
		log.Error("Error: no se encontraron transacciones")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}
	if error != nil {
		log.Error("Error:", error)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	log.Info("Obteniendo transacciones")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Agregar una nueva transacción
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	log := GetLogger(r)
	tenant := r.Header.Get("x-tenant-id")
	var response rc.ApiResponse

	txService, ok := service.GetTrxService(tenant)
	if !ok {
		json.NewEncoder(w).Encode(response)
		log.Error("TenantId invalido:", r.Header)
		return
	}
	var t models.Transaction
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		log.WithField("error", err.Error()).Error("Error decodificando la request")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(rc.ErrorBad(fmt.Sprintf("Error en el formato JSON: %v", err), http.StatusBadRequest))
		return
	}

	// Insertar la transacción usando el servicio
	insert, response, err := txService.AddTrx(&t, log)
	if err != nil {
		log.WithField("error", err.Error()).Error("Error al insertar la transacción")
		w.WriteHeader(response.Code)
		json.NewEncoder(w).Encode(response)
		return
	}

	if !insert {
		log.Error("No se pudo insertar la transacción")
		w.WriteHeader(response.Code)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(response.Code)
	json.NewEncoder(w).Encode(response)
}
