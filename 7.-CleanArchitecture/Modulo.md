```

📂 marketplace
 ├── 📂 domain               # Entidades, Value Objects, Interfaces
 │   ├── entity.go           # Definición de entidades (e.g., Product, Order)
 │   ├── repository.go       # Interfaces para persistencia
 │   ├── service.go          # Interfaces de servicios de dominio
 ├── 📂 usecase              # Casos de uso (Application Layer)
 │   ├── create_product.go   # Caso de uso para crear productos
 │   ├── list_products.go    # Caso de uso para listar productos
 ├── 📂 infrastructure       # Adaptadores externos (DB, API, etc.)
 │   ├── repository          # Implementaciones de repositorios
 │   │   ├── pg_product.go   # Persistencia en PostgreSQL con pgx
 │   ├── cache               # Implementaciones de caché
 │   │   ├── redis_cache.go  # Uso de Redis para caching
 ├── 📂 delivery             # Interfaces de entrada (HTTP, gRPC, eventos)
 │   ├── http                # Handlers HTTP con Chi
 │   │   ├── product_handler.go  # Handlers para productos
 │   ├── event               # Consumo de eventos (SNS, SQS, Kafka, etc.)
 ├── 📂 config               # Configuración específica del módulo
 │   ├── config.yaml         # Variables de configuración
```
