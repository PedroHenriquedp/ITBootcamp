/*
Exercício 3 - Qual mês corresponde a


Crie um aplicativo que receba uma variável com o número do mês.

Dependendo do número, imprima o mês correspondente no texto.
Você consegue pensar em maneiras diferentes de resolver isso? Qual delas você escolheria e por quê?
*/

package main

import (
	"fmt"
)

func main() {
	// O codigo ficara feio, mas eu quero me acostumar com a sintaxe de map e for
	var mesEscolhido int
	meses := [12]string{"Janeiro", "Fevereiro", "Marco", "Abril", "Maio", "Junho", "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"}
	mesMap := make(map[int]string)
	var i int = 1

	for _, v := range meses {
		mesMap[i] = v
		i++
	}

	fmt.Print("Escreva o numero do mes\n")
	fmt.Scan(&mesEscolhido)



	fmt.Print("O mes escolhido foi:\n")
	fmt.Printf("%s\n", mesMap[mesEscolhido])
}
