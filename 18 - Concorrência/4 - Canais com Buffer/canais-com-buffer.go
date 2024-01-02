package main

import "fmt"

func main() {
	// INFO: O canal com buffer só bloqueia se atingir a capacidade máxima de escrita ou
	// mínima de leitura.
	canal := make(chan string, 4)
	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"
	canal <- "Programando em Go De Novo!"
	canal <- "Programando em Go De Novo!"

	mensagem := <-canal
	mensagem2 := <-canal
	mensagem3 := <-canal
	mensagem4 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
	fmt.Println(mensagem3)
	fmt.Println(mensagem4)
}
