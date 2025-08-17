package http

import (
	"net/http"
)

type ApiResponse struct {
	Code  int         `json:"code"`  // Código HTTP (ej. 200, 500)
	Msg   string      `json:"msg"`   // Mensaje descriptivo
	Data  interface{} `json:"data"`  // Datos de la respuesta
	Error interface{} `json:"error"` // Errores, si los hay
}

func Success(data interface{}) ApiResponse {
	return ApiResponse{
		Code:  http.StatusOK,
		Msg:   "Ok",
		Data:  data,
		Error: nil,
	}
}

func ErrorBad(message string, err interface{}) ApiResponse {
	return ApiResponse{
		Code:  http.StatusBadRequest,
		Msg:   message,
		Data:  nil,
		Error: err,
	}
}
func ErrorGeneric(message string, err interface{}) ApiResponse {
	return ApiResponse{
		Code:  http.StatusInternalServerError,
		Msg:   message,
		Data:  nil,
		Error: err,
	}
}

func ErrorCode(code int, message string, err interface{}) ApiResponse {
	return ApiResponse{
		Code:  code,
		Msg:   message,
		Data:  nil,
		Error: err,
	}
}
