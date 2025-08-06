package main

import (
	"log"
	"my-app-tx/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
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
