package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anderson2093/ejercicios-go/6.CleanArchitecture/config"
	"github.com/anderson2093/ejercicios-go/6.CleanArchitecture/routes"
	"github.com/go-chi/chi/v5"
)

func main() {
	// Cargar configuración y base de datos
	config.LoadEnv()
	config.ConnectDB()
	defer config.CloseDB()

	// Configurar el router
	r := chi.NewRouter()
	routes.SetupRoutes(r)

	// Iniciar servidor
	port := ":8080"
	fmt.Println("🚀 Servidor corriendo en http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, r))
}
