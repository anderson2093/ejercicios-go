```
/your_project/
 ├── /cmd/                    # Punto de entrada de la aplicación
 │   ├── main.go               # Inicialización principal
 ├── /config/                  # Configuración de la aplicación
 │   ├── config.go             # Carga de variables de entorno
 ├── /internal/
 │   ├── /domain/              # Entidades y modelos de negocio
 │   │   ├── user.go           # Modelo de usuario
 │   ├── /usecase/             # Casos de uso / lógica de negocio
 │   │   ├── user_usecase.go   # Casos de uso para usuarios
 │   ├── /handler/             # Manejo de HTTP (controllers)
 │   │   ├── user_handler.go   # Handlers para usuarios
 │   ├── /repository/          # Interfaces de acceso a datos
 │   │   ├── user_repository.go
 │   ├── /infrastructure/      # Infraestructura externa
 │   │   ├── /database/        # Bases de datos
 │   │   │   ├── postgres.go   # Conexión a PostgreSQL
 │   │   │   ├── redis.go      # Conexión a Redis
 │   │   │   ├── mongodb.go    # Conexión a MongoDB
 │   │   ├── /http/            # Servidor HTTP
 │   │   │   ├── router.go     # Definición de rutas con `chi`
 │   │   │   ├── middleware.go # Middlewares de autenticación
 │   │   ├── /external_api/    # Conexión con APIs externas
 │   │   │   ├── whatsapp.go   # WhatsApp API
 │   │   │   ├── billing.go    # API de facturación electrónica
 ├── /pkg/                     # Código reutilizable
 │   ├── logger/               # Implementación de logs
 ├── go.mod
 ├── go.sum
```
