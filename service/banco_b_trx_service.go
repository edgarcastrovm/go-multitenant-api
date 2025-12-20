package service

import (
	"context"
	"my-app-tx/data"
	"my-app-tx/utils/constants"
	. "my-app-tx/utils/http"
	"my-app-tx/utils/logger"
	. "my-app-tx/utils/models"
	"net/http"
	"strings"
	"time"
)

type BancoB struct{}

func (b *BancoB) AddTrx(t *Transaction, ctx context.Context) (bool, ApiResponse, error) {
	log := logger.WithContext(ctx).Sugar()
	log.Infof("addTrx: ", constants.BANCO_B)
	t.Tipo = strings.ToUpper(t.Tipo)
	t.Empresa = constants.BANCO_B
	t.Fecha = time.Now()
	t.Estado = constants.STATE_TRX_ACTIVE

	// Validar la transacción antes de insertarla
	err := t.ValidateCreate()
	if err != nil {
		log.Errorf("Error validando la transacción [ %v ]", err)
		return false, ErrorBad("Error validando la transacción:", err), nil

	}

	newTrx, err := data.AddTrx(*t, constants.BANCO_B)

	if err != nil {
		log.Error("Error guardando la transacción")
		return false, ErrorGeneric("Error registrando la transacción", err), nil
	}

	return true, Success(newTrx), nil
}

func (b *BancoB) GetAll(ctx context.Context) (bool, ApiResponse, error) {
	log := logger.WithContext(ctx).Sugar()
	log.Infof("getAll: ", constants.BANCO_B)
	lstTrx, err := data.GetTrx(constants.BANCO_B)

	if err != nil {
		log.Error("Error obteniendo las transacciones")
		return false, ErrorGeneric("Error listando las transacciones", err), nil
	}

	return true, Success(lstTrx), nil
}

func (b *BancoB) GetById(id *int8, ctx context.Context) (bool, ApiResponse, error) {
	log := logger.WithContext(ctx).Sugar()
	log.Infof("getById: ", constants.BANCO_B)
	trx, exist, err := data.GetTrxByID(*id, constants.BANCO_B)

	if err != nil {
		log.Error("Error obteniendo la transaccion")
		return false, ErrorGeneric("Error listando las transacciones", err), nil
	}

	if !exist {
		log.Error("La transaccion no existe")
		return false, ErrorCode(http.StatusNotFound, "Error listando las transacciones", nil), nil
	}

	return true, Success(trx), nil
}

func (b *BancoB) Reverse(t *Transaction, ctx context.Context) (bool, ApiResponse, error) {
	log := logger.WithContext(ctx).Sugar()
	log.Infof("reverse Banco A")
	return true, Success("reverse Banco A"), nil
}

func init() {
	RegisterTrxService(constants.BANCO_B, &BancoB{})
}
