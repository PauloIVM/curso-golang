package main

// INFO: Executa um arquivo go:
// >> go run FILE_NAME.go

// INFO: Cria um go.mod:
// >> go mod init MODULE_NAME

// INFO: Cria um arquivo executável/binário com o nome do módulo:
// >> go build

import (
	"fmt"
	"modulo/auxiliar"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
}
