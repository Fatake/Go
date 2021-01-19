package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("/home/decodigo/go/archivos/archivo_ejemplo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		fmt.Println(scanner.Bytes())
	}
}
