package models

import "time"

type Transaccion struct {
	ID          int       `json:"id"`
	Monto       float64   `json:"monto"`
	Tipo        string    `json:"tipo"` // "ingreso" o "gasto"
	Fecha       time.Time `json:"fecha"`
	Descripcion string    `json:"descripcion"`
}
