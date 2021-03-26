package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Olá mundo!")
	auxiliar.Escrever()
	//auxiliar.escrever2()

	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)
}
