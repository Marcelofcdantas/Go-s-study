package main

import "fmt"

func main() {
	func () {
		fmt.Println("Ola mundo")
	}()
	retorno:= func (texto string) string {
		return fmt.Sprintf("Ola mundo do %s", texto)
	}("parametro")

	fmt.Println(retorno)
}