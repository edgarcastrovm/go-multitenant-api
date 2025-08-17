package service

import (
	"my-app-tx/data"
	"my-app-tx/utils/constants"
	. "my-app-tx/utils/http"
	. "my-app-tx/utils/middleware"
	. "my-app-tx/utils/models"
	"net/http"

	"github.com/sirupsen/logrus"
)

type BancoA struct{}

const (
	Pkg  = "service"
	File = "banco_a_trx_service.go"
)

func (b *BancoA) AddTrx(t *Transaction, _log *logrus.Entry) (bool, ApiResponse, error) {
	log := GetLog(Pkg, File, _log)
	log.Info("addTrx: ", constants.BANCO_A)

	err := t.ValidateCreate()
	if err != nil {
		log.Error("Error guardando la transacción")
		return false, ErrorBad("Error registrando la transacción: ", err), nil

	}

	newTrx, err := data.AddTrx(*t, constants.BANCO_A)

	if err != nil {
		log.Error("Error guardando la transacción")
		return false, ErrorGeneric("Error registrando la transacción: ", err), nil
	}

	return true, Success(newTrx), nil
}

func (b *BancoA) GetAll(_log *logrus.Entry) (bool, ApiResponse, error) {
	log := GetLog(Pkg, File, _log)
	log.Info("getAll: ", constants.BANCO_A)
	lstTrx, err := data.GetTrx(constants.BANCO_A)

	if err != nil {
		log.Errorf("Error obteniendo las transacciones:%v", err)
		return false, ErrorGeneric("Error listando las transacciones: ", err), nil
	}

	return true, Success(lstTrx), nil
}

func (b *BancoA) GetById(id *int8, log *logrus.Entry) (bool, ApiResponse, error) {
	log.Info("getById: ", constants.BANCO_A)
	trx, exist, err := data.GetTrxByID(*id, constants.BANCO_A)

	if err != nil {
		log.Error("Error obteniendo la transaccion")
		return false, ErrorGeneric("Error listando las transacciones: ", err), nil
	}

	if !exist {
		log.Error("La transaccion no existe")
		return false, ErrorCode(http.StatusNotFound, "Error listando las transacciones", nil), nil
	}

	return true, Success(trx), nil
}

func (b *BancoA) Reverse(t *Transaction, log *logrus.Entry) (bool, ApiResponse, error) {
	log.Info("reverse Banco A")
	return true, Success("reverse Banco A"), nil
}

func init() {
	RegisterTrxService(constants.BANCO_A, &BancoA{})
}
