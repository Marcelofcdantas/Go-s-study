package main

import "fmt"

func main() {
	soma := 1 + 2
	subtracao := 4 - 2
	multiplicacao := 3 * 6
	divisao := 10 / 2
	restoDaDivisao := 10 % 3

	fmt.Print(soma, subtracao, multiplicacao, divisao, restoDaDivisao)

	fmt.Println("a" != "b")
	fmt.Println("a" == "b")

	verdade, falso := true, false
	fmt.Println(verdade && falso)
	fmt.Println(verdade || falso)
	fmt.Println(!verdade)

	numero := 10
	numero += 1
	fmt.Println(numero)
	numero++
	fmt.Println(numero)

	numero *= 10
	fmt.Println(numero)
	numero /= 3
	fmt.Println(numero)
	numero %= 3
	fmt.Println(numero)

	// Não existe operador ternário em Go

}