package main

import (
	"cap4/modulo"
	"fmt"
)

func main() {
	fmt.Println(modulo.Somar(1, 2))

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	resultado := f("texto digitado")
	fmt.Println(resultado)

	rSum, rSub := modulo.CalculosMatematicos(2, 2)
	fmt.Println(rSum, rSub)

	_, rSub2 := modulo.CalculosMatematicos(15, 20)
	fmt.Println(rSub2)
}
