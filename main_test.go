package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"testing"

	"my-app-tx/config"
	"my-app-tx/data"
	my_http "my-app-tx/utils/http"
	"my-app-tx/utils/models"

	_ "github.com/lib/pq" // Driver PostgreSQL si no está en data
	"github.com/stretchr/testify/assert"
)

var tenantID = "banco_a"

func TestMain(m *testing.M) {
	config.SetupTestEnv()

	// Configurar y ejecutar pruebas para Gin
	config.SetupTestServer("gin")
	codeGin := m.Run()
	config.TearDownTestServer()

	// Configurar y ejecutar pruebas para Mux
	config.SetupTestServer("mux")
	codeMux := m.Run()
	config.TearDownTestServer()

	// Limpieza global (cerrar BD)
	data.CloseDB()

	// Salir con el código combinado (ej. si uno falla, falla todo)
	if codeGin != 0 || codeMux != 0 {
		os.Exit(1)
	}
	os.Exit(0)
}

func testCreateTransaction(t *testing.T) {
	if config.TestServer == nil {
		t.Skip("Servidor no configurado para Gin")
		return
	}

	payload := models.Transaction{
		Cuenta:        "12345",
		CuentaDestino: "67890",
		Monto:         250.75,
		Tipo:          "INGRESO",
		Descripcion:   "Pago mensual de prueba",
	}
	jsonPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", config.TestServer.URL+"/transacciones", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-tenant-id", tenantID)

	resp, err := config.TestClient.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var createdTx models.Transaction
	json.Unmarshal(body, &createdTx)
	assert.Equal(t, payload.Cuenta, createdTx.Cuenta)
	assert.Equal(t, payload.Monto, createdTx.Monto)
}

func TestListTransactions(t *testing.T) {
	if config.TestServer == nil {
		t.Skip("Servidor no configurado para Gin")
		return
	}

	// Asumiendo que se creó una transacción en TestGinCreateTransaction (pero como TRUNCATE es global, recreamos aquí para aislamiento)
	// Para mejor aislamiento, mueve la creación dentro de esta prueba o usa transacciones DB.

	// payload := models.Transaction{
	// 	Cuenta:        "54321",
	// 	CuentaDestino: "09876",
	// 	Monto:         100.50,
	// 	Tipo:          "EGRESO",
	// 	Descripcion:   "Transferencia de prueba",
	// }
	// jsonPayload, _ := json.Marshal(payload)

	// postReq, _ := http.NewRequest("POST", testServer.URL+"/transacciones", bytes.NewBuffer(jsonPayload))
	// postReq.Header.Set("Content-Type", "application/json")
	// postReq.Header.Set("x-tenant-id", tenantID)
	// testClient.Do(postReq) // Ignorar respuesta para simplicidad

	// Ahora GET
	getReq, _ := http.NewRequest("GET", config.TestServer.URL+"/transacciones", nil)
	getReq.Header.Set("x-tenant-id", tenantID)

	resp, err := config.TestClient.Do(getReq)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	defer resp.Body.Close()

	var apiResponse my_http.ApiResponse
	var transactions []models.Transaction

	json.Unmarshal(body, &apiResponse)
	log.Printf("Response Data: %+v", apiResponse.Code)

	dataBytes, err := json.Marshal(apiResponse.Data)
	assert.NoError(t, err)
	err = json.Unmarshal(dataBytes, &transactions)
	assert.NoError(t, err, "Error al unmarshalar Data a []Transaction")

	log.Printf("Response type Data: %+v", reflect.TypeOf(transactions))
	log.Printf("Response count Data: %+v", transactions)
	log.Printf("Response count Data: %+v", len(transactions))

	// Ahora valida codigo 200
	assert.Equal(t, 200, apiResponse.Code)
	// Ahora valida la longitud
	assert.GreaterOrEqual(t, len(transactions), 1)
}
