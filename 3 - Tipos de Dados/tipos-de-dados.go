package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	//int8, int16, int32, int64
	//int = usa arquitetura do pc. 32 bits ou 64bits
	var numero int64 = 10000000000000
	fmt.Println(numero)

	//
	var numero2 uint32 = 1000
	fmt.Println(numero2)

	//int32 = rune
	var numero3 rune = 123456
	fmt.Println(reflect.TypeOf(numero3))

	//uint8 = byte

	//float32
	var algumafloat32 float32 = 10.1
	//float64
	var algumfloat64 float64 = 10.1
	//float
	numeroReal3 := 123412.12
	fmt.Println(algumafloat32, algumfloat64, numeroReal3)

	var texto string
	var intvazio int16
	var realvazio float32

	fmt.Println(texto, intvazio, realvazio)

	var bol1 bool = true
	var bol2 bool
	fmt.Println(bol1, bol2)

	var erro error
	erro2 := errors.New("Erro interno")
	fmt.Println(erro, erro2)

}
