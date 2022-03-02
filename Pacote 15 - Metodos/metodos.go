package main

import "fmt"

type usuario struct {
	nome string
	idade uint
}

func (u usuario) salvar() {
	fmt.Printf("Metodo salvar %s no BD\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func main() {
	usuario1 := usuario{"Jorge", 14}
	fmt.Println(usuario1)
	usuario1.salvar()

	fmt.Println(usuario1.maiorDeIdade())	

}