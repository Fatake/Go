package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const ( //Raw strings para el control con expresiones regulares de las operaciones que recibira la calculadora.
	operationRegexp string = `[0-9]+(\.[0-9]+)?[\+\-\*/][0-9]+(\.[0-9]+)?` //Comprueba que inserte un float operador float
	operatorsRegexp string = `[\+\-\*/]`                                   //Comprueba si se trata de un operador permitido.
)

func getInput() string {
	fmt.Print("[*] Ingresa una operacion (ejem:5+2)\n->")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}

func obtainOperation() (string, error) {
	operation := getInput()
	expr := operationRegexp
	patern := regexp.MustCompile(expr)
	if patern.MatchString(operation) { //Comprobamos si la entrada concuerda con nuestra expresion, en caso contrario devolvemos un error.
		return operation, nil
	}
	return "", errors.New("[!]Lexical: La entrada no es un operador")
}

func obtainNumbers(operation string) (float64, float64) {
	expr := operatorsRegexp
	patern := regexp.MustCompile(expr)
	numbers := patern.Split(operation, 2)
	operand1, _ := strconv.ParseFloat(numbers[0], 32)
	operand2, _ := strconv.ParseFloat(numbers[1], 32)
	return operand1, operand2
}

func obtainOperator(operation string) byte {
	expr := operatorsRegexp
	patern := regexp.MustCompile(expr)
	indexes := patern.FindIndex([]byte(operation))[0]
	return operation[indexes]
}

func operate() float64 {
	operation, err := obtainOperation() //Obtener la operacion del usuario
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	value1, value2 := obtainNumbers(operation)                  //Obtener los valores
	return calculate(value1, value2, obtainOperator(operation)) //Devolver el calculo obtenido a partir de los valores y el operador.
}

func calculate(value1 float64, value2 float64, operator byte) float64 {
	var ans float64
	switch operator {
	case '+':
		ans = value1 + value2
	case '-':
		ans = value1 - value2
	case '*':
		ans = value1 * value2
	case '/':
		ans = value1 / value2
	}
	return ans
}

func main() {
	ans := operate() //Funcion que realizara el calculo
	fmt.Printf("[*] El resultado es %f\n", ans)
}
