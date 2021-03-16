package main

import (
	"fmt"
)

func main() {
	/*
		Los canales permite comunicacion entre go routunes
	*/

	// Concurrente
	fmt.Printf("\nCanales \n<------------------------->\n")
	canal := make(chan string)

	go func(canal chan string) {
		for {
			var nombre string
			fmt.Scanln(&nombre)

			// el canal sirve como pipe
			// entra algo y sale
			canal <- nombre
		}
	}(canal)

	mensage := <-canal
	fmt.Printf("E recivido : %s \n", mensage)

	mensage = <-canal
	fmt.Printf("E recivido : %s", mensage)
}
