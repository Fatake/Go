package main

import (
	"fmt"

	"paquetes.com/ejemplo/utilerias"
)

func main() {
	var cocheNuevo utilerias.Coche
	cocheNuevo.Marca = "Funciona"
	fmt.Println(cocheNuevo)
}
