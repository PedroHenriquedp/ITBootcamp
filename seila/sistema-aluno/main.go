package main

import (
	"fmt"

	"github.com/google/uuid"
)

// func main() {
// 	var aiiMessi uuid.UUID = uuid.New()
// 	fmt.Printf("O valor do uuid e %s", aiiMessi)
// 	// type Aluno struct {
// 	// 	UUID
// 	// }
// }

func incluiAluno() {

}

func main() {
	var numberSelected string

	type Aluno struct {
		uuid     uuid.UUID
		nome     string
		idade    uint8
		telefone string
	}

	for i := 0; i != 1; {
		fmt.Printf("=== Menu de Opções ===\n 1 - Incluir Aluno\n 2 - Listar Todos os Alunos\n 3 - Pesquisar Aluno por Matrícula\n 4 - Alterar Aluno\n 5 - Excluir Aluno\n 6 - Sair\n")
		fmt.Scanf("%s", &numberSelected)
		fmt.Printf("Numero selecionado foi %s\n", numberSelected)

		if numberSelected == "1" {

		}

		if numberSelected == "6" {
			i++
		}
	}
}
