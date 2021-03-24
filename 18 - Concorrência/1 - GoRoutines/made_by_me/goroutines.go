package main

import (
	"fmt"
	"time"
)

func main() {

	go escrever("Raphael")
	go escrever("Laila")
	escrever("Lara")

}

func escrever(texto string) {

	for i := 0; i <= 3; i++ {

		fmt.Printf("[%d] %s\n", i, texto)
		time.Sleep(time.Second)

	}
	fmt.Println(" ")
}
