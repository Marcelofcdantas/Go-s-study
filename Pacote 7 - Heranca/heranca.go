package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura float32
}

type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main() {

	p1 := pessoa{"Joao", "Henrique", 24, 1.78}
	fmt.Println(p1)

	e1 := estudante{p1, "ADS", "UFF"}
	fmt.Println(e1)
	fmt.Println(e1.altura)
}