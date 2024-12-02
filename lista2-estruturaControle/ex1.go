/*

Exercício 1 - Letras em uma palavra
A Real Academia Espanhola quer saber quantas letras uma palavra tem e, em seguida, ter cada uma das letras separadamente para soletrá-la. Para isso, eles terão de: 
Criar um aplicativo que receba uma variável com a palavra e imprima o número de letras que ela contém. 
Em seguida, imprima cada uma das letras.
*/

package main

import (
	"fmt"
)

func main() {
	var palavra string
	fmt.Print("Digite uma palavra:")
	fmt.Scan(&palavra)

	letras := len(palavra)
	fmt.Printf("A palavra '%s' tem %d letras. \n", palavra, letras)

	fmt.Println("As letras sao:")
	for i, letra := range palavra {
		fmt.Printf("Letra %d: %c\n", i+1, letra)
	}
}

