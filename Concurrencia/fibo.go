package main

import (
	"fmt"
	"time"
)

func main() {
	go animacion(100 * time.Millisecond)
	const n = 45
	resultado := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, resultado)
}

func animacion(retraso time.Duration) {
	anima := "\\|/-"
	for {
		for _, v := range anima {
			// retorno de carro en la misma linea
			fmt.Printf("\r%c", v)
			time.Sleep(retraso)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
