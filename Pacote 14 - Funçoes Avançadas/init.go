package main

import "fmt"

var i int

func main() {
	fmt.Println("main")
	println(i)
}

// funcao init é SEMPRE a primeira a ser iniciada, independentemente da ordem
func init() {
	fmt.Println("init")
	i = 10
}