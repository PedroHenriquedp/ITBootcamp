/*
Um banco deseja conceder empréstimos a seus clientes, mas nem todos têm acesso a eles. Ele tem certas regras para determinar quais clientes podem receber empréstimos. O banco só concede empréstimos a clientes com mais de 22 anos de idade, que estejam empregados e que estejam em seu emprego há mais de um ano. Dentro dos empréstimos concedidos, o banco não cobrará juros daqueles que tiverem um salário superior a US$ 100.000.
Você precisa criar um aplicativo que receba essas variáveis e imprima uma mensagem de acordo com cada caso.
📌Dica: seu código deve ser capaz de imprimir pelo menos três mensagens diferentes.
*/

package main

import (
	"fmt"
)

func main() {
	var idadeCliente int
	var salario int
	var empregado bool
	var empregoTemp int

	fmt.Print("Digite a idade do cliente:\n")
	fmt.Scan(&idadeCliente)

	fmt.Print("Digite o salario do cliente\n")
	fmt.Scan(&salario)

	fmt.Print("Esta empregado?\n")
	fmt.Scan(&empregado)

	fmt.Print("Ha quanto tempo esta trabalhando?\n")
	fmt.Scan(&empregoTemp)

	if(idadeCliente >= 22 && empregado == true && empregoTemp >= 1 && salario >= 100000) {
		fmt.Print("Aprovado sem juros\n")
		return
	}

	if(idadeCliente >= 22 && empregado == true && empregoTemp >= 1) {
		fmt.Print("Aprovado com juros\n")
		return
	}
	fmt.Print("Nao esta apto!\n")
}
