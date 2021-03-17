package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("\rIniciando servidor")
	listener, err := net.Listen("tcp", "localhost:9878")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	for {
		// Acepta nueva coneccion
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// Crea una coneccion
		go manejaCliente(conn)
	}
}

func manejaCliente(coneccion net.Conn) {
	defer coneccion.Close()
	for {
		_, err := io.WriteString(coneccion, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
