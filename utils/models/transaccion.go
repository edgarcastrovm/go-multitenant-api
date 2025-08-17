package models

import (
	. "my-app-tx/utils/constants"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Transaction struct {
	ID            int       `json:"id,omitempty"`
	Cuenta        string    `json:"cuenta"`
	CuentaDestino string    `json:"cuenta_destino"`
	Monto         float64   `json:"monto"`
	Tipo          string    `json:"tipo"`
	Fecha         time.Time `json:"fecha"`
	Descripcion   string    `json:"descripcion,omitempty"`
	Estado        int       `json:"estado,omitempty"`
	Empresa       string    `json:"empresa,omitempty"`
}

func (t Transaction) ValidateCreate() error {
	return validation.ValidateStruct(&t,
		// Cuenta: Requerida, longitud entre 5 y 20 caracteres
		validation.Field(&t.Cuenta, validation.Required, validation.Length(5, 20)),

		// CuentaDestino: Requerida, longitud entre 5 y 20 caracteres
		validation.Field(&t.CuentaDestino, validation.Required, validation.Length(5, 20)),

		// Monto: Requerido, mayor o igual a 0
		validation.Field(&t.Monto, validation.Required, validation.Min(0.0)),

		// Tipo: Requerido, debe ser "ingreso" o "egreso"
		validation.Field(&t.Tipo, validation.Required, validation.In(TYPE_TRX_INGRESO, TYPE_TRX_EGRESO)),

		// Fecha: Requerida, no en el futuro
		validation.Field(&t.Fecha, validation.Required, validation.Max(time.Now())),

		// Descripcion: Requerida, longitud entre 5 y 100 caracteres
		validation.Field(&t.Descripcion, validation.Required, validation.Length(5, 100)),

		// Estado: Requerido, entre 0 y 2
		validation.Field(&t.Estado, validation.Required, validation.In(STATE_TRX_ACTIVE, STATE_TRX_DELETED, STATE_TRX_REVERSED)),

		// Empresa: Requerida, longitud entre 3 y 50 caracteres, alfanumérica
		validation.Field(&t.Empresa, validation.Required, validation.Length(3, 25), is.Alphanumeric),
	)
}
