package main

import (
	"fmt"
	"strings"
)

func main() {
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println("Arreglo: ", array, "\nTamaño: ", len(array), "\nCapacidad:", cap(array))

	// Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("\nSlice: ", slice, "\nTamaño: ", len(slice), "\nCapacidad:", cap(slice))

	// Métodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 11)
	fmt.Println(slice)

	// Append
	newSlice := []int{12, 13, 14}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	fmt.Println(isPalindromeSinRange("Anita lava la tina"))

	fmt.Println(isPalindromeConRange("Anita lava la tina"))
}

func isPalindromeSinRange(text string) string {
	var textReverse string
	text = strings.ToLower(text)               // Pasamos el texto a minúsculas
	text = strings.ReplaceAll(" ", text, text) // Eliminamos los espacios
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		return "No es palindrome"
	} else {
		return "Es palindrome"
	}

}

func isPalindromeConRange(palabra string) string {
	for i := range palabra {
		if palabra[i] != palabra[len(palabra)-1-i] {
			return "No es palindrome con Range"
		}
	}
	return "Es palidrome con Range"
}
