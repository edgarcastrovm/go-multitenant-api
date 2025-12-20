package config

import (
	"log"
	"my-app-tx/utils/logger"
	"net/http"
	"net/http/httptest"
	"time"
)

var TestServer *httptest.Server
var TestClient *http.Client

func SetupTestEnv() {
	// Inicializar logger global
	logger.InitLogger()

	// Carga variables de entorno de prueba
	errenv := LoadEnv(".test_env")
	if errenv != nil {
		log.Fatalf("Error al cargar las variables de entorno: %v", errenv)
	}

	// Configurar la base de datos
	err := SetUpDb()
	if err != nil {
		log.Fatalf("Error al configurar la base de datos: %v", err)
	}

}

func SetupTestServer(routerOpt string) {
	router, err := SetUpHttpServer(routerOpt)
	if err != nil {
		log.Fatalf("Error configurando router %s: %v", routerOpt, err)
	}

	TestServer = httptest.NewServer(router)
	TestClient = TestServer.Client()

	// Esperar un poco para que el servidor esté listo
	time.Sleep(100 * time.Millisecond)
}

func TearDownTestServer() {
	if TestServer != nil {
		TestServer.Close()
		TestServer = nil
		TestClient = nil
	}
}
