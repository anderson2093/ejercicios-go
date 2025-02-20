package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Manejador personalizado en GO")
}

func main() {
	handler := MyHandler{}
	fmt.Println("Servidor en ejecución en http://localhost:8080")
	http.ListenAndServe(":8080", handler)
}
