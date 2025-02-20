package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error en la petición:", err)
		return
	}
	defer resp.Body.Close() // Cierra el cuerpo de la respuesta

	body, _ := io.ReadAll(resp.Body) // Leer el cuerpo
	fmt.Println(string(body))
}
