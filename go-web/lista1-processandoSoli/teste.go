package main

import (
	"encoding/json"
	"log"
	"os"
)

type Aluno struct {
	Nome string `json:"nome"`
	Idade int `json:"idade"`
	Nota float64 `json:"nota"
}

func main() {
	alunos := []Aluno {
		{"Joao", 20, 8.5},
		{"Maria", 22, 9.6},
		{"Pedro", 19, 7.5},
	}

	file, err := os.Create("alunos.json")
	if err != nil {
		log.Fatalf("Erro ao criar arquivo: %v", err)
	}
	defer file.Close()

	enconder := json.NewEncoder(file)

	// Escrever cada aluno como uma linha JSON

	for _, aluno := range alunos {
		err := encoder.Encode(aluno)
		if err != nil {
			log.fatalf("Erro ao codificar alun: %v", err)
		}
	}
}
ear
