package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println("Slice Vacio:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("Slice s:", s)
	fmt.Println("Elemento s[2]:", s[2])

	fmt.Println("len(s):", len(s))
	fmt.Println("cap(s):", cap(s))

	s = append(s, "d")      // agreda d
	s = append(s, "e", "f") // agrega ef
	fmt.Println("append(s,\"elemenets\"...):", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copia de s make(type[],len(s)):", c)

	l := s[2:5]
	fmt.Println("Paticion s[2:5]:", l)

	l = s[:5]
	fmt.Println("Paticion s[:5]:", l)

	l = s[2:]
	fmt.Println("Paticion s[2:]:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("t := []string{\"elem\",...}:", t)

	twoD := make([][]int, 3)
	fmt.Println("Matriz m := make([][]int, 3):", twoD)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("Matris llenada: ", twoD)
}
