package main

import (
    "fmt"
    "os"
)

func main() {
    // Passo 1: Criar um arquivo
    file, err := os.Create("dados.txt")
    if err != nil {
        fmt.Println("Erro ao criar o arquivo:", err)
        return
    }
    defer file.Close() // Garante que o arquivo será fechado ao final

    // Passo 2: Escrever dados no arquivo
    _, err = file.WriteString("Olá, este é o meu primeiro arquivo em Go!\n")
    if err != nil {
        fmt.Println("Erro ao escrever no arquivo:", err)
        return
    }

    _, err = file.WriteString("Manipular arquivos é fácil com Go.\n")
    if err != nil {
        fmt.Println("Erro ao escrever no arquivo:", err)
        return
    }
    fmt.Println("Dados escritos no arquivo com sucesso!")

    // Passo 3: Abrir e ler o arquivo
    data, err := os.ReadFile("dados.txt")
    if err != nil {
        fmt.Println("Erro ao ler o arquivo:", err)
        return
    }

    // Passo 4: Exibir os dados no console
    fmt.Println("\nConteúdo do arquivo:")
    fmt.Println(string(data))
}
