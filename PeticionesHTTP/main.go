package main

import (
	"server.go"
)

func main() {
	servidor := server.NewServidor(":3000")
	servidor.Listen()
}
