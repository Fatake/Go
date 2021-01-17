package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	/*var edad int
	fmt.Print("Ingrese la edad \n->")
	fmt.Scanf("%d", &edad)*/
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese la operacion (suma, de la forma numero + numero, ej: 2+2): ")
	scanner.Scan()

	operacion := scanner.Text()

	fmt.Println("La operacion ingresada es: ", operacion)

	valores := strings.Split(operacion, "+")

	fmt.Println("Estos son los valores ingresados: ", valores)
	fmt.Println("Primer y segundo valor sumados como texto: ", valores[0]+valores[1])

	// Cast valores from text to number
	operador1, err1 := strconv.Atoi(valores[0])
	operador2, err2 := strconv.Atoi(valores[1])

	if err1 != nil {
		fmt.Println("Se a introducido un error")
	}
	if err2 != nil {
		fmt.Println("Error en el valor 2")
	}

	fmt.Println("Suma de los dos operadores matematicamente: ", operador1+operador2)

}
