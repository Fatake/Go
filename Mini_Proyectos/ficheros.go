package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	hello := []byte("Prueba Escritura")
	// 						nombre archi,input,permisos
	err := ioutil.WriteFile("test.txt", hello, 0644)
	if err != nil {
		panic(err)
	}

	// Read test.txt
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Lecura: " + string(data))
}
