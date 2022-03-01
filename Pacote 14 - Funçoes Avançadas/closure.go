package main

import "fmt"

func closure() func() {
	texto := "Dentro do closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro do main"
	fmt.Println(texto)
	funcaoNova := closure()
	funcaoNova()
}