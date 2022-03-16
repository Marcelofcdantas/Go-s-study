package main

import (
	"fmt"
	"time"
)

func main() {
	// concorrencia != paralelismo
	// no paralelismo duas tarefas sao executadas ao mesmo tempo
	// na concorrencia as tarefas PODEM ser executadas de forma alternada ou não
	go escrever("ola mundo")
	escrever("programando em Go")
}

func escrever(texto string) {
	for{
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}