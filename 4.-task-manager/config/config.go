package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // 🔥 IMPORTANTE: Importa el driver de PostgreSQL
)

var DB *sql.DB

func ConnectDB() {
	connStr := "user=anderson password=AmoProgramar dbname=practicadb sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error al conectar con la base de datos: ", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("No se puede hacer ping a la base de datos: ", err)
	}

	fmt.Println("📦 Conectado a PostgreSQL!")
}
