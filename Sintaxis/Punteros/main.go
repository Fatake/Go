package main

import (
	"fmt"

	pk "paquetes.com/puntero/Utilerias"
)

func main() {
	myPc := pk.New(12, 200, "HP")
	myPc.SetRam(16)
	myPc.FormatPrint()
	fmt.Println("Se duplica la ram")
	myPc.DuplicateRAM()
	myPc.FormatPrint()
}

type Pc struct {
	ram   int
	disk  int
	brand string
}

func New(rm, dsk int, brn string) Pc {
	myPc := Pc{ram: rm, disk: dsk, brand: brn}
	return myPc
}

func (self Pc) FormatPrint() {
	fmt.Printf("Esta pc marca %s cuenta con una ram de %dGB y un disco de %dGB.\n", self.brand, myPc.ram, myPc.disk)
}

func (myPc *Pc) DuplicateRAM() {
	myPc.ram = myPc.ram * 2
}

func (myPc Pc) GetRam() int {
	return myPc.ram
}

func (myPc *Pc) SetRam(rm int) {
	myPc.ram = rm
}
