package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0	

	for i <= 10 {
		fmt.Println(i)
		time.Sleep(time.Second)
		i++
	}
	fmt.Println("Acabou")

	for j:= 10; j >=0; j-- {
		fmt.Println("Decrementando", j)
		time.Sleep(time.Second)
	}
	fmt.Println("Acabou 2")

	nomes := [3]string{ "Joao", "Jorge", "Jack"}

	for indice, nome := range(nomes) {
		fmt.Println(indice, nome)
	}

	for _, nome := range(nomes) {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRAS"{
		fmt.Println(indice, letra)
		fmt.Println(string(letra))
	}

	usuario := map[string]string {
		"nome": "Eu",
		"sobrenome": "mesmo",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}