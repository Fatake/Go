package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Lectura de Linea a Linea\n<------------------------->")

	file, err := os.Open("./hola.txt")
	if err != nil {
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		linea := scanner.Text()
		fmt.Println(linea)
		time.Sleep(250 * time.Millisecond)
	}

	file.Close()
}
