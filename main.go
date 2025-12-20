package main

import (
	"flag"
	"log"
	"my-app-tx/config"
	"my-app-tx/utils/logger"
	"net/http"
)

func main() {
	// Inicializar logger global
	logger.InitLogger()

	//Carga variables de entorno
	errenv := config.LoadEnv()
	if errenv != nil {
		log.Fatalf("Error al cargar las variables de entorno: %v", errenv)
		return
	}

	// Configurar la base de datos
	err := config.SetUpDb()
	if err != nil {
		log.Fatalf("Error al configurar la base de datos: %v", err)
		return
	}

	router_opt := flag.String("router", "gin", "Especifica el enrutador HTTP a usar: 'mux' (Gorilla Mux) o 'gin' (Gin Gonic)")
	flag.Parse()

	// Configurar el enrutador
	router, err := config.SetUpHttpServer(*router_opt)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	// Iniciar servidor
	log.Fatal(http.ListenAndServe(":8080", router))
}
