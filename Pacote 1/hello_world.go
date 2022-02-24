package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello World.")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("teste@gmail.com")
	fmt.Println("\n", erro)
}