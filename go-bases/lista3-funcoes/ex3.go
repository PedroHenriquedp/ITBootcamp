package main

import "fmt"

func calcularSalario(minutosTrabalhados int, categoria string) float64 {
	horasTrabalhadas := float64(minutosTrabalhados) / 60.0
	var salario float64

	if categoria == "C" {
		salario = horasTrabalhadas * 1000.0
		return salario
	}
	if categoria == "B" {
		salarioBase := horasTrabalhadas * 1500.0
		salario = salarioBase + (salarioBase * 0.20)
		return salario
	}
	if categoria == "A" {
		salarioBase := horasTrabalhadas * 3000.0
		salario = salarioBase + (salarioBase * 0.50)
		return salario
	}
    return salario
}

func main() {
	minutosTrabalhadas := 1200
	categoria := "C"

	resultado := calcularSalario(minutosTrabalhadas, categoria)

	fmt.Printf("O valor do resultado de quanto o trabalhador ira receber e %.2f\n", resultado)
}
