package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct {
	rua string
	numero int8
}


func main() {
	fmt.Println("structs")
	
	var u usuario
	u.idade = 30
	u.nome = "Eu"
	fmt.Println(u)
	
	exemploEndereco := endereco{"rua x", 3}

	usuario2 := usuario{"Jorge", 12, exemploEndereco}
	fmt.Println(usuario2)
	
	usuario3 := usuario{nome: "Jack"}
	fmt.Println(usuario3)
}