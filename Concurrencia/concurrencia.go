package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	canal := make(chan string)
	servidores := []string{
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
		"localhost",
	}

	// Manera secuencial
	timer := time.Now()
	for _, v := range servidores {
		revisarServidorSec(v)
	}

	fmt.Println("[time Secuencial] ", time.Since(timer))
	fmt.Println("\n<------------------->")
	// Manera Concurrente
	timer = time.Now()
	for _, v := range servidores {
		go revisaServidorCon(v, canal)
		fmt.Print(<-canal)
	}

	fmt.Println("[time Paralela] ", time.Since(timer))

	fmt.Printf("\n\n[*] End\n\n")
}

func revisarServidorSec(servidor string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Printf("[!] Servidor: %s está caido\n", servidor)
	} else {
		fmt.Printf("[+] Servidor: %s está Activo\n", servidor)
	}
}

func revisaServidorCon(servidor string, canal chan string) {
	// go usa la palabra reservada go para los hilos
	_, err := http.Get(servidor)
	if err != nil {
		canal <- "[!] Servidor: " + servidor + " está Caido\n"
	} else {
		canal <- "[+] Servidor: " + servidor + " está Activo\n"
	}
}
