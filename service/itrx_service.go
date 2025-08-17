package service

import (
	"context"
	. "my-app-tx/utils/http"
	. "my-app-tx/utils/models"

	"strings"
)

type ITrxService interface {
	AddTrx(t *Transaction, ctx context.Context) (bool, ApiResponse, error)
	GetAll(ctx context.Context) (bool, ApiResponse, error)
	GetById(id *int8, ctx context.Context) (bool, ApiResponse, error)
	Reverse(t *Transaction, ctx context.Context) (bool, ApiResponse, error)
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
