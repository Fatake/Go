package main

import (
	"fmt"
)

// ejemplo dead lock
func main() {
	/*
		Los canales permite comunicacion entre go routunes
	*/

	// Concurrente
	fmt.Printf("\nCanales \n<------------------------->\n")
	canal1 := make(chan string)
	canal2 := make(chan int)

	go func(canal chan string, contador chan int) {
		var stop int
		for {
			stop = <-canal2
			if stop > 5 {
				break
			}
			var nombre string
			fmt.Scanln(&nombre)

			// el canal sirve como pipe
			// entra algo y sale
			canal <- nombre
		}
	}(canal1, canal2)
	i := 0
	for {
		canal2 <- i
		i++
		mensage := <-canal1
		fmt.Printf("E recivido : %s \n", mensage)

	}
}
