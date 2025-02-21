package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/anderson2093/ejercicios-go/5.chi-task/config"
	"github.com/anderson2093/ejercicios-go/5.chi-task/routes"
)

func main() {
	// Cargar configuración
	config.LoadEnv()
	config.ConnectDB()
	defer config.CloseDB()

	// Obtener puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Iniciar servidor
	fmt.Println("🚀 Servidor en http://localhost:" + port)
	http.ListenAndServe(":"+port, routes.SetupRouter())
}
