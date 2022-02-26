package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	var array1 [5] int
	fmt.Println(array1)

	var array2 [5] string
	fmt.Println(array2)
	array2[2] = "espaço 3"
	fmt.Println(array2)

	array3 := [3]string {"item 1", "item 2", "item 3"}
	fmt.Println(array3)

	// é basicamente um array de tamanho dinâmico
	slice := []int {10, 11, 12, 13, 14}
	fmt.Println(slice)
	slice = append(slice, 15)
	fmt.Println(slice)

	slice2 := array3[1:3]
	fmt.Println(slice2)

	array3[1] = "nada"
	fmt.Println(slice2)
}