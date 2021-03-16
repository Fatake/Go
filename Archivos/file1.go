package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Lectura de archivos\n<------------------------->")
	// ReadFile carga todo en ram   ruta relativa
	file, err := ioutil.ReadFile("./hola.txt")
	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}
	fmt.Println("En bytes")
	fmt.Println(file)
	fmt.Println("\nEn String")
	fmt.Println(string(file))
}
