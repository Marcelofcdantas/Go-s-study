package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Ola", canal)
	mensagem := <- canal
	fmt.Println(mensagem)
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}