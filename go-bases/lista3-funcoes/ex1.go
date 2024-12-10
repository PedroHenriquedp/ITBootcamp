package main

import "fmt"


func calcularImposto(salario float64) float64 {
    if salario > 150000 {
        return salario * 0.27 // 27%
    } else if salario > 50000 {
        return salario * 0.17 // 17%
    }
    return 0.0
}

func main() {
	salario := 9000.0
	resultado := calcularImposto(salario)
	fmt.Printf("O imposto a ser pago pelo salario e de %.2f\n", resultado)
}
