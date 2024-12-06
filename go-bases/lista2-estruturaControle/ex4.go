package main

import (
	"fmt"
)

func main() {
	var maiorIdade int
	var  
	empregadosMap = make(map[string]int)
	empregadosMap["Benjamin"] = 20
	empregadosMap["Nahuel"] = 26
	empregadosMap["Brenda"] = 19
	empregadosMap["Dario"] = 44
	empregadosMap["Pedro"] = 30

	for _, v := range empregadosMap {
		if v >= 21 {
			maiorIdade++		
		}
	}

	fmt.Printf("Pessoas com mais de 21 anos: %d\n",  maiorIdade)
	fmt.Printf("Adicionando Federico\n")
	empregadosMap["Frederico"] = 25
	fmt.Printf("Removendo Pedro...\n")
	delete(empregadosMap, "Pedro")
	
	fmt.Printf("Os empregados sao:\n")
	
	for k, v := range empregadosMap {
		fmt.Printf("%s %d anos\n", k, v)
	}
	
}
