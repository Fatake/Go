package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg es utilizado para indicar al programa que espere a
// que finalicen las rutinas
var wg sync.WaitGroup

func main() {
	wg.Add(2) // indicamos que esperara dos rutinas
	fmt.Println("[m] Iniciando go rutines")

	go imprimirCantidad("A")

	go imprimirCantidad("B")

	fmt.Println("[m] Esperando rutinas...")
	wg.Wait()
	fmt.Println("\r\n[m] end")

}

func imprimirCantidad(etiqueta string) {
	defer wg.Done() // indicamos cuando ya termina la go rutine
	for i := 0; i < 10; i++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("[%s] Trabajando %d\n", etiqueta, (i + 1))
	}
}
