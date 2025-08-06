package data

import (
	"my-app-tx/models"
)

var transacciones []models.Transaccion
var currentID int

// Agrega una transacción al slice y asigna un ID incremental
func AddTransaccion(t models.Transaccion) models.Transaccion {
	currentID++
	t.ID = currentID
	transacciones = append(transacciones, t)
	return t
}

// Obtiene todas las transacciones
func GetTransacciones() []models.Transaccion {
	return transacciones
}

// Obtiene una transacción por ID
func GetTransaccionByID(id int) (models.Transaccion, bool) {
	for _, t := range transacciones {
		if t.ID == id {
			return t, true
		}
	}
	return models.Transaccion{}, false
}

// Elimina una transacción por ID
func DeleteTransaccion(id int) bool {
	for i, t := range transacciones {
		if t.ID == id {
			transacciones = append(transacciones[:i], transacciones[i+1:]...)
			return true
		}
	}
	return false
}
