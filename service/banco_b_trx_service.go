package service

import (
	"my-app-tx/data"
	"my-app-tx/utils/constants"
	. "my-app-tx/utils/http"
	. "my-app-tx/utils/models"
	"net/http"

	"github.com/sirupsen/logrus"
)

type BancoB struct{}

func (b *BancoB) AddTrx(t *Transaction, log *logrus.Entry) (bool, ApiResponse, error) {
	log.Info("addTrx: ", constants.BANCO_B)
	newTrx, err := data.AddTrx(*t, constants.BANCO_B)

	if err != nil {
		log.Error("Error guardando la transacción")
		return false, ErrorGeneric("Error registrando la transacción", err), nil
	}

	return true, Success(newTrx), nil
}

func (b *BancoB) GetAll(log *logrus.Entry) (bool, ApiResponse, error) {
	log.Info("getAll: ", constants.BANCO_B)
	lstTrx, err := data.GetTrx(constants.BANCO_B)

	if err != nil {
		log.Error("Error obteniendo las transacciones")
		return false, ErrorGeneric("Error listando las transacciones", err), nil
	}

	return true, Success(lstTrx), nil
}

func (b *BancoB) GetById(id *int8, log *logrus.Entry) (bool, ApiResponse, error) {
	log.Info("getById: ", constants.BANCO_B)
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

func (b *BancoB) Reverse(t *Transaction, log *logrus.Entry) (bool, ApiResponse, error) {
	log.Info("reverse Banco A")
	return true, Success("reverse Banco A"), nil
}

func init() {
	RegisterTrxService(constants.BANCO_B, &BancoB{})
}
