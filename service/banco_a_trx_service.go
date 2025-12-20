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

	"go.uber.org/zap"
)

type BancoA struct{}

const (
	Pkg  = "service"
	File = "banco_a_trx_service.go"
)

func (b *BancoA) AddTrx(t *Transaction, ctx context.Context) (bool, ApiResponse, error) {
	log := logger.WithContext(ctx).Sugar()
	log.Infof("addTrx: ", constants.BANCO_A)
	t.Tipo = strings.ToUpper(t.Tipo)
	t.Empresa = constants.BANCO_A
	t.Fecha = time.Now()
	t.Estado = constants.STATE_TRX_ACTIVE

	// Validar la transacción antes de insertarla
	err := t.ValidateCreate()
	if err != nil {
		log.Error("Error guardando la transacción")
		return false, ErrorBad("Error registrando la transacción: ", err), nil

	}

	newTrx, err := data.AddTrx(*t, constants.BANCO_A)

	if err != nil {
		log.Errorf("Error guardando la transacción: %v", zap.Error(err))
		return false, ErrorGeneric("Error registrando la transacción: ", err), nil
	}

	return true, Success(newTrx), nil
}

func (b *BancoA) GetAll(ctx context.Context) (bool, ApiResponse, error) {
	log := logger.WithContext(ctx).Sugar()
	log.Infof("getAll: %s ", constants.BANCO_A)
	lstTrx, err := data.GetTrx(constants.BANCO_A)

	if err != nil {
		log.Errorf("Error obteniendo las transacciones:%v", zap.Error(err))
		return false, ErrorGeneric("Error listando las transacciones: ", err), nil
	}

	return true, Success(lstTrx), nil
}

func (b *BancoA) GetById(id *int8, ctx context.Context) (bool, ApiResponse, error) {
	log := logger.WithContext(ctx).Sugar()
	trx, exist, err := data.GetTrxByID(*id, constants.BANCO_A)

	if err != nil {
		log.Errorf("Error obteniendo la transaccion", zap.Error(err))
		return false, ErrorGeneric("Error listando las transacciones: ", err), nil
	}

	if !exist {
		log.Error("La transaccion no existe")
		return false, ErrorCode(http.StatusNotFound, "Error listando las transacciones", nil), nil
	}

	return true, Success(trx), nil
}

func (b *BancoA) Reverse(t *Transaction, ctx context.Context) (bool, ApiResponse, error) {
	log := logger.WithContext(ctx)
	log.Info("reverse Banco A")
	return true, Success("reverse Banco A"), nil
}

func init() {
	RegisterTrxService(constants.BANCO_A, &BancoA{})
}
