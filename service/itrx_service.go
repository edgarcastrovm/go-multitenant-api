package service

import (
	. "my-app-tx/utils/http"
	. "my-app-tx/utils/models"

	"strings"

	"github.com/sirupsen/logrus"
)

type ITrxService interface {
	AddTrx(t *Transaction, log *logrus.Entry) (bool, ApiResponse, error)
	GetAll(log *logrus.Entry) (bool, ApiResponse, error)
	GetById(id *int8, log *logrus.Entry) (bool, ApiResponse, error)
	Reverse(t *Transaction, log *logrus.Entry) (bool, ApiResponse, error)
}

var trxServices = map[string]ITrxService{}

func RegisterTrxService(name string, service ITrxService) {
	key := strings.ToUpper(strings.TrimSpace(name))
	trxServices[key] = service
}

func GetTrxService(name string) (ITrxService, bool) {
	key := strings.ToUpper(strings.TrimSpace(name))
	srv, ok := trxServices[key]
	return srv, ok
}

//
