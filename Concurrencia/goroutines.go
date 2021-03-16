package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// Secuencial
	fmt.Printf("Iniciando secuencial \n<------------------------->\n")
	inicio := time.Now()
	miNombre("Paulo")

	fin := time.Now()
	fmt.Printf("Tiempo %v \n", fin.Sub(inicio))

	// Concurrente
	fmt.Printf("\nIniciando Concurrente \n<------------------------->\n")
	inicio = time.Now()
	go miNombre("Paulo")

	fin = time.Now()
	fmt.Printf("Tiempo %v \n", fin.Sub(inicio))

	// Funciones anonimas
	go func() {
		fmt.Println("Funcion anonima en concurrente")
	}()
}

func miNombre(name string) {
	letras := strings.Split(name, "")
	for _, letra := range letras {
		fmt.Println(letra)
		time.Sleep(1000 * time.Millisecond)
	}
}
