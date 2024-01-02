package main

import (
	"fmt"
	"time"
)

// INFO: Esse patern pode atribuir ao "resultados" elementos com o index "desalinhados"
// com os elementos do "tarefas"? Sim... os workers vão executar bem fora de ordem. Então
// numa implementação desse tipo, é provável que eu precise jogar no "resultados" uma structure
// com algum atributo que referencie qual "tarefa" gerou o respectivo resultado.
func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	// INFO: Se eu manter um único worker, então será printado os resultados em ordem. Mas se eu
	// adicionar mais workers, então será printado os resultados fora de ordem. Se eu tirar o sleep,
	// mesmo com mais de um worker, então ele volta a printar ordenadamente... não é tão óbvio o
	// problema

	// TODO: Um desafio interessante: Modificar a structure de forma que eu possa reordená-las
	// (nesse caso seria simplesmente ordenar já que é um array de inteiros), e então usar o pattern de
	// concorrêcia e por fim ordenar os resultados de acordo com a correspondêcia do array de tarefas.

	go worker2(tarefas, resultados)
	go worker2(tarefas, resultados)
	go worker2(tarefas, resultados)
	go worker2(tarefas, resultados)

	for i := 1; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 1; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

}

func worker2(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- randomFunc(numero)
	}
}

func randomFunc(posicao int) int {
	time.Sleep(time.Second / time.Duration(posicao))
	return posicao
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
