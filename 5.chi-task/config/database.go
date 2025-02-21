package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var DB *pgx.Conn

// Cargar variables de entorno
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando .env")
	}
}

// Conectar a PostgreSQL con pgx
func ConnectDB() {
	var err error
	DB, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Error al conectar a PostgreSQL:", err)
	}
	fmt.Println("📦 Conectado a PostgreSQL!")

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT
	);`

	_, err = DB.Exec(context.Background(), createTableSQL)
	if err != nil {
		log.Fatal("Error al crear la tabla:", err)
	}

	fmt.Println("✅ Tabla 'tasks' verificada/creada")
}

// Cerrar la conexión
func CloseDB() {
	DB.Close(context.Background())
}
