package main

import (
	"fmt"
	"os"
	"strconv"
)

func parse(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error: format: int '<operator>' int")
		os.Exit(1)
	}
	return value
}

type calculador struct {
	intA     int
	intB     int
	operator string
}

func (self calculador) operate() *int {

	total := 0

	switch self.operator {
	case "+":
		total = self.intA + self.intB
	case "-":
		total = self.intA - self.intB
	case "*":
		total = self.intA * self.intB
	case "/":
		total = self.intA / self.intB
	}

	return &total
}

func main() {

	if len(os.Args) != 4 {
		fmt.Println("Error: format: int '<operator>' int")
		os.Exit(1)
	}

	instance := calculador{
		intA:     parse(os.Args[1]),
		intB:     parse(os.Args[3]),
		operator: os.Args[2],
	}

	response := instance.operate()
	if response == nil {
		fmt.Println("Only '+', '-', '*' and '/' operators are allowed")
		os.Exit(1)
	}

	fmt.Println(*response)
}
