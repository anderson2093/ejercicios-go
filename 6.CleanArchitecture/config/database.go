package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func ConnectDB() {
	var err error
	DB, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("❌ Error al conectar a PostgreSQL:", err)
	}
	fmt.Println("✅ Conectado a PostgreSQL!")

	// Crear tabla si no existe
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT
	);`
	_, err = DB.Exec(context.Background(), createTableSQL)
	if err != nil {
		log.Fatal("❌ Error al crear la tabla:", err)
	}

	fmt.Println("📌 Tabla 'tasks' verificada/creada")
}

func CloseDB() {
	DB.Close(context.Background())
}
