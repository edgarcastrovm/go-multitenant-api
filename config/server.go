package config

import (
	"fmt"
	"log"
	ctl_gin "my-app-tx/controller/gin"
	ctl_mux "my-app-tx/controller/mux"
	"my-app-tx/data"
	md "my-app-tx/utils/middleware"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func SetUpHttpServer(router_opt string) (http.Handler, error) {
	// Usar un switch para usar mux o gin
	switch router_opt {
	case "mux":
		router := mux.NewRouter()
		// Agregamos middleware para interceptar las peticiones

		router.Use(md.RequestIDMiddleware)
		// Definir rutas
		router.HandleFunc("/transacciones", ctl_mux.GetTransactions).Methods("GET")
		router.HandleFunc("/transacciones", ctl_mux.CreateTransaction).Methods("POST")

		// Iniciar servidor
		log.Println("Servidor corriendo con mux en http://localhost:8080")

		return router, nil
	case "gin":
		// Crear enrutador de Gin
		router := gin.Default()

		// Agregar middleware para interceptar las peticiones
		router.Use(md.RequestIDMiddlewareGin)

		// Definir rutas
		router.GET("/transacciones", ctl_gin.GetTransactions)
		router.POST("/transacciones", ctl_gin.CreateTransaction)

		// Iniciar servidor
		log.Println("Servidor corriendo con gin en http://localhost:8080")

		return router, nil
	default:
		log.Fatalf("Error: el valor de 'router' debe ser 'mux' o 'gin', recibido: %s", router_opt)
		return nil, fmt.Errorf("enrutador inválido: %s", router_opt)
	}
}

func SetUpDb() error {
	// Obtener las variables de entorno
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	// Validar que todas las variables estén definidas
	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" || dbSSLMode == "" {
		log.Println("una o más variables de entorno de la base de datos no están definidas")
		return fmt.Errorf("variables de entorno de la base de datos no definidas")
	}

	// Construir la cadena de conexión
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)

	// Valida Conexión a PostgreSQL
	err := data.InitDB(connStr)
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
		return err
	}

	log.Println("Conexión a la base de datos establecida")
	return nil
}

func LoadEnv(envFile ...string) error {

	if len(envFile) > 0 {
		// Se habilita ambiente test
		fmt.Println("Variables entorno test:", envFile[0])
		if err := godotenv.Load(envFile[0]); err != nil {
			log.Printf("No se encontró archivo: %v", envFile[0])
			return err
		}
	} else {
		// Si habilita ambiente real
		fmt.Println("Variables entorno real")
		// Cargar variables de entorno
		if err := godotenv.Load(); err != nil {
			log.Println("No se encontró archivo .env")
			return err
		}
	}
	return nil
}
