package handlers

import (
	"encoding/json"
	"my-app-tx/data"
	"my-app-tx/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Listar todas las transacciones
func GetTransacciones(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	transacciones, err := data.GetTransacciones()
	if err != nil {
		http.Error(w, "Error obteniendo las transacciones", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(transacciones)
}

// Agregar una nueva transacción
func CreateTransaccion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var t models.Transaccion
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Error en el formato JSON", http.StatusBadRequest)
		return
	}

	// Validación de datos
	if t.Monto <= 0 || (t.Tipo != "ingreso" && t.Tipo != "gasto") || t.Descripcion == "" {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	nuevaTransaccion, err := data.AddTransaccion(t)
	if err != nil {
		http.Error(w, "Error agregando una transaccion", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevaTransaccion)
}

// Obtener una transacción específica
func GetTransaccion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	t, found, err := data.GetTransaccionByID(id)
	if err != nil {
		http.Error(w, "Error obteniendo las transacciones", http.StatusInternalServerError)
	}
	if !found {
		http.Error(w, "Transacción no encontrada", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(t)
}

// Eliminar una transacción
func DeleteTransaccion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	deleted, err := data.DeleteTransaccion(id)

	if err != nil {
		http.Error(w, "Error eliminando la transacción", http.StatusNotFound)
		return
	}
	if deleted {
		http.Error(w, "Transacción no encontrada", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
