# Task Manager API

## Descripción
Task Manager API es una aplicación REST en Go que permite gestionar tareas. Los usuarios pueden crear, listar, obtener y eliminar tareas mediante endpoints HTTP.

## Tecnologías Utilizadas
- **Go (Golang)**
- **PostgreSQL**
- **net/http** (Manejo de servidores HTTP en Go)
- **encoding/json** (Serialización y deserialización de datos JSON)
- **github.com/lib/pq** (Driver de PostgreSQL para Go)

## Estructura del Proyecto
```
/4.-task-manager
├── config/        # Configuración de la base de datos
├── db/            # Inicialización de la base de datos
├── handlers/      # Lógica para manejar las solicitudes HTTP
├── models/        # Definición de modelos y operaciones en la base de datos
├── routes/        # Definición de rutas
├── main.go        # Punto de entrada de la aplicación
```

## Instalación
### Prerrequisitos
1. Instalar [Go](https://go.dev/doc/install)
2. Instalar [PostgreSQL](https://www.postgresql.org/download/)
3. Configurar las variables de entorno:
   ```sh
   export DB_USER=tu_usuario
   export DB_PASSWORD=tu_contraseña
   export DB_NAME=task_manager
   export DB_HOST=localhost
   export DB_PORT=5432
   ```

### Clonar el repositorio
```sh
git clone https://github.com/anderson2093/ejercicios-go.git
cd ejercicios-go/4.-task-manager
```

### Ejecutar la aplicación
```sh
go run main.go
```

## Endpoints
### 1. Obtener todas las tareas
- **GET /tasks**
- **Respuesta:** JSON con la lista de tareas

### 2. Crear una nueva tarea
- **POST /tasks**
- **Cuerpo de la petición:**
  ```json
  {
    "title": "Nombre de la tarea",
    "description": "Descripción de la tarea"
  }
  ```
- **Respuesta:** Estado 201 si se crea con éxito

### 3. Obtener una tarea por ID
- **GET /tasks/{id}**
- **Respuesta:** JSON con la tarea solicitada o error 404 si no existe

### 4. Eliminar una tarea
- **DELETE /tasks/{id}**
- **Respuesta:** Estado 204 si se elimina con éxito

## Base de Datos
La base de datos se inicializa automáticamente con la siguiente estructura:
```sql
CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    status VARCHAR(50) NOT NULL DEFAULT 'pending'
);
```

## Autores
- **Anderson Cusma Vasquez**

## Licencia
Este proyecto está bajo la licencia MIT.

