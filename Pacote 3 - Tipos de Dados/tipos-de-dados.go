package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero1 int8 = 100
	fmt.Println(numero1)
	var numero2 int16 = 10000
	fmt.Println(numero2)
	var numero3 int32 = 1000000000
	fmt.Println(numero3)
	var numero4 int64 = 1000000000000000
	fmt.Println(numero4)
	numero5 := 10000000000000
	fmt.Println(numero5)
	// tipo uint é semelhante ao int, todavia não comporta números negativos
	// rune = int32
	var numero6 rune = 1234
	fmt.Println(numero6)
	// uint8 = byte
	var numero7 = 4
	fmt.Print(numero7)
	// float pode ser float32 ou float64
	var numero8 float32 = 1234.56
	fmt.Println(numero8)
	var stringType string = "variavel"
	fmt.Println(stringType)
	// não existe char em Go caso tente usar será pega a tabela ASCII - desde que use aspas simples
	char := 'A'
	fmt.Println(char)
	// é possível declarar uma variável sem atribuir um valor. 
	// No caso de texto será uma string em branco, de numero será 0 e boleano será false
	var texto string
	fmt.Println(texto)
	var numerozerado int64
	fmt.Println(numerozerado)
	var verdade bool = true
	fmt.Println(verdade)
	var verdadeiro bool
	fmt.Println(verdadeiro)
	var erro error
	fmt.Println(erro)
	var erros error = errors.New("Erro do Servidor")
	fmt.Println(erros)
}