package main

import "fmt"

func calcularMedia(notas ...float64) (float64) {
    var soma float64
    for _, nota := range notas {
        soma += nota
    }

    media := float64(soma) / float64(len(notas))
    return media
}

func main() {
	var notas = []float64{6.2, 7.7, 3.3}

	resultado := calcularMedia(notas...)
	fmt.Printf("A media do aluno foi %.2f\n", resultado)
}
