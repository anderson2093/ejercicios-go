package db

import (
	"fmt"
	"log"

	"github.com/anderson2093/ejercicios-go/4.-task-manager/config"
)

// Crear una tabla si no existe (ejemplo inicializacion)
func InitTables() {
	query := `
CREATE TABLE IF NOT EXISTS tasks (
	id SERIAL PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	description TEXT,
	status VARCHAR(50) NOT NULL DEFAULT 'pending'
	);`

	_, err := config.DB.Exec(query)
	if err != nil {
		log.Fatal("❌ Error al crear tabla: ", err)
	}
	fmt.Println("✅ Tabla 'tasks' verificada!")
}
