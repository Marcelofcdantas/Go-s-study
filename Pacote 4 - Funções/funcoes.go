package main

import "fmt"

func somar(n1 int16, n2 int16) int16 {
	return n1 + n2
}

// Go permite que as funções tenham mais de um retorno
func calculosMatematicos(n1, n2 int32) (int32, int32) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}


func main() {
	soma := somar(16, 24)
	fmt.Println(soma)

	var f = func(text string) string {
		fmt.Println(text)
		return "0"
	}

	f("texto")

	resultado1, resultado2 := calculosMatematicos(50, 30)
	fmt.Println(resultado1, resultado2)

	// o _ serve para receber e "apagar" uma variavel que não se deseje usar quando tiver mais de 1 retorno
	_, resultado3 := calculosMatematicos(50, 30)
	fmt.Println(resultado3)
}