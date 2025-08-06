package main

import (
	"log"
	"my-app-tx/data"
	"my-app-tx/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Conexión a PostgreSQL
	connStr := "host=localhost port=5432 user=postgres password=novopayment dbname=novo_db sslmode=disable"
	if err := data.InitDB(connStr); err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}
	router := mux.NewRouter()

	// Definir rutas
	router.HandleFunc("/transacciones", handlers.GetTransacciones).Methods("GET")
	router.HandleFunc("/transacciones", handlers.CreateTransaccion).Methods("POST")
	router.HandleFunc("/transacciones/{id}", handlers.GetTransaccion).Methods("GET")
	router.HandleFunc("/transacciones/{id}", handlers.DeleteTransaccion).Methods("DELETE")

	// Iniciar servidor
	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
