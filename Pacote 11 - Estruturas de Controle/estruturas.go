package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := 10

	if numero < 15 {
		fmt.Println("menor")
	} else {
		fmt.Println("maior ou igual")
	}

	// variavel criada dentro do escopo do if é local
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("maior")
	} else if outroNumero == 0{
		fmt.Println("zero")
	} else {
		fmt.Println("menor")
	}
}