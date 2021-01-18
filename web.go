package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Structura escritor

type escritorWeb struct {
}

// hacemos la "herencia" de la interface write de go
func (escritorWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	sitio := "http://google.com"
	respuesta, err := http.Get(sitio) // hace una peticion
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	// respuesta es un paquete http
	fmt.Println(respuesta)
	fmt.Println("\n\n")
	// body
	fmt.Println(respuesta.Body)

	// Pasamor a generar un escritor y lector
	// creamos la funcion copy
	escritor := escritorWeb{}
	io.Copy(escritor, respuesta.Body)
}
