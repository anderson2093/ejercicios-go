package main

import (
	"fmt"
	"net/http"

	"github.com/anderson2093/ejercicios-go/4.-task-manager/config"
	"github.com/anderson2093/ejercicios-go/4.-task-manager/db"
)

func main() {
	config.ConnectDB()
	db.InitTables()

	fmt.Println("Servidor en ejecucion http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
