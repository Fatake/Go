package main

import "fmt"

type animal interface {
	mover() string
}

type dog struct {
}

type fish struct {
}

type bird struct {
}

func (dog) mover() string {
	return "Soy un perro y corro"
}

func (fish) mover() string {
	return "Soy un pez y nado"
}

func (bird) mover() string {
	return "Soy un pajaro y vuelo"
}

func moverAnimal(a animal) {
	fmt.Println(a.mover())
}

func main() {
	d := dog{}
	f := fish{}
	b := bird{}

	moverAnimal(d)
	moverAnimal(f)
	moverAnimal(b)
}
