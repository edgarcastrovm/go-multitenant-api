package main

import (
	"flag"
	"fmt"
	"log"
	ctl_gin "my-app-tx/controller/gin"
	ctl_mux "my-app-tx/controller/mux"
	"my-app-tx/data"
	"my-app-tx/utils/logger"
	md "my-app-tx/utils/middleware"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Inicializar logger global
	logger.InitLogger()
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("No se encontró archivo .env")
		return
	}
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
		return
	}

	// Construir la cadena de conexión
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)

	// Valida Conexión a PostgreSQL
	err := data.InitDB(connStr)
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}

	router_opt := flag.String("router", "mux", "Especifica enrutador: 'gorilla/mux' o 'gin'")

	flag.Parse()

	// Usar un switch para usar mux o gin
	switch *router_opt {
	case "mux":
		router := mux.NewRouter()
		// Agregamos middleware para interceptar las peticiones

		router.Use(md.RequestIDMiddleware)
		// Definir rutas
		router.HandleFunc("/transacciones", ctl_mux.GetTransactions).Methods("GET")
		router.HandleFunc("/transacciones", ctl_mux.CreateTransaction).Methods("POST")
		// router.HandleFunc("/transacciones/{id}", controller.GetTransaccion).Methods("GET")
		// router.HandleFunc("/transacciones/{id}", controller.DeleteTransaccion).Methods("DELETE")

		// Iniciar servidor
		log.Println("Servidor corriendo en http://localhost:8080")
		log.Fatal(http.ListenAndServe(":8080", router))
	case "gin":
		// Crear enrutador de Gin
		router := gin.Default()

		// Agregar middleware para interceptar las peticiones
		router.Use(md.RequestIDMiddlewareGin)

		// Definir rutas
		router.GET("/transacciones", ctl_gin.GetTransactions)
		router.POST("/transacciones", ctl_gin.CreateTransaction)
		// router.GET("/transacciones/:id", controller.GetTransaccionGin)
		// router.DELETE("/transacciones/:id", controller.DeleteTransaccionGin)

		// Iniciar servidor
		log.Println("Servidor corriendo en http://localhost:8080")
		log.Fatal(router.Run(":8080"))
		// Aquí iría la lógica para el servidor B
	default:
		log.Fatalf("Error: el valor de 'router' debe ser 'mux' o 'gin', recibido: %s", *router_opt)
	}

}
