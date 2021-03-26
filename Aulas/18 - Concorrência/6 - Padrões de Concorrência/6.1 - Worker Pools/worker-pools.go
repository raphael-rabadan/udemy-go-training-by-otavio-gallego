package main

import "fmt"

func main() {

	qtdVezes := 45
	tarefas := make(chan int, qtdVezes)
	resultados := make(chan int, qtdVezes)

	for i := 0; i < qtdVezes; i++ {
		go worker(tarefas, resultados)
	}

	// go worker(tarefas, resultados)
	// go worker(tarefas, resultados)
	// go worker(tarefas, resultados)
	// go worker(tarefas, resultados)

	for i := 0; i < qtdVezes; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < qtdVezes; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

}

func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
