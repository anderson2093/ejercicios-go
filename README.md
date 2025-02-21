# Ejercicios en Go

Este repositorio contiene una serie de ejercicios en Go relacionados con el desarrollo de aplicaciones web y gestión de tareas. Se divide en tres carpetas principales:

## 📂 Estructura del Repositorio

- **Servidor HTTP**: Implementa un servidor HTTP básico en Go.
- **HTTP Handler, GET**: Manejo de solicitudes HTTP con el método GET.
- **Task Manager**: Gestión de tareas con operaciones CRUD básicas.

## 🚀 Requisitos

Para ejecutar los ejercicios, necesitas tener instalado:

- [Go](https://golang.org/doc/install) (versión 1.18 o superior)

## 📌 Uso

1. Clona el repositorio:

   ```bash
   git clone https://github.com/anderson2093/ejercicios-go.git
   cd ejercicios-go
   ```

2. Ve a la carpeta del ejercicio que deseas ejecutar y compila el código:

   ```bash
   cd "Nombre-de-la-carpeta"
   go run main.go
   ```

## 📜 Descripción de cada carpeta

### 1️⃣ Servidor HTTP

Este ejercicio implementa un servidor HTTP básico en Go que responde a las solicitudes en un puerto específico.

Ejecuta:

```bash
 go run main.go
```

Accede en tu navegador a `http://localhost:8080` para probarlo.

---

### 2️⃣ HTTP Handler, GET

Manejo de peticiones HTTP con el método GET. Permite capturar parámetros desde la URL y responder con datos dinámicos.

Ejemplo de uso:

```bash
 go run main.go
```

Luego, en el navegador o con `curl`:

```bash
 curl http://localhost:8080/saludo?nombre=Anderson
```

Salida esperada:

```
Hola, Anderson!
```

---

### 3️⃣ Task Manager

Ejercicio que implementa una API para la gestión de tareas. Permite agregar, listar y eliminar tareas.

Ejecuta:

```bash
 go run main.go
```

Prueba la API con `curl` o Postman:

- Obtener todas las tareas:
  ```bash
  curl http://localhost:8080/tareas
  ```
- Agregar una tarea:
  ```bash
  curl -X POST -d '{"titulo": "Aprender Go"}' http://localhost:8080/tasks
  ```
- Eliminar una tarea:
  ```bash
  curl -X DELETE http://localhost:8080/tasks/1
  ```

## 🛠 Contribuciones

Si deseas contribuir, ¡siéntete libre de hacer un fork y enviar un pull request! ✨

## 📄 Licencia

Este proyecto está bajo la licencia MIT. Consulta el archivo `LICENSE` para más detalles.

