package main

import "fmt"

func recuperarExecucao() {
	// o recover impede que o panic trave o programa, dando continuidade
	if r := recover(); r != nil {
		fmt.Println("Execucao recuperada")
	}
}

func alunoAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	// o panic trava o programa
	panic("Média é 6")
}

func main() {
	fmt.Println(alunoAprovado(6, 7))
	fmt.Println("Pós execução")
}