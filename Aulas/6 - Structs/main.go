package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {

	enderecoEx := endereco{"Rua qq", 111}

	usuario1 := usuario{"Davi", 16, enderecoEx}
	fmt.Println(usuario1)

	usuario2 := usuario{nome: "Davi2"}
	fmt.Println(usuario2)

	var usuario3 usuario
	usuario3.nome = "davi3"
	usuario3.idade = 12
	fmt.Println(usuario3)
}
