package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // Importa el driver de PostgreSQL
)

var DB *sql.DB

func ConnectDB() {
	// Obtener valores desde variables de entorno
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	// Construir la cadena de conexión de forma segura
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("❌ Error al conectar con la base de datos:", err)
	}

	// Probar conexión
	if err := DB.Ping(); err != nil {
		log.Fatal("❌ No se puede hacer ping a la base de datos:", err)
	}

	fmt.Println("✅ 📦 Conectado a PostgreSQL!")
}
