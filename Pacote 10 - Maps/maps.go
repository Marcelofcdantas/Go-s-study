package main

import "fmt"

func main() {
	fmt.Println("Maps")
	usuario := map[string]string {
		"Nome": "Jorge",
		"Sobrenome": "Nogueira",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["Nome"])

	usuario2 := map[string]map[string]string {
		"Nome": {
			"Prenome": "Joao",
			"Sobrenome": "da Silva",
		},
		"Profissão": {
			"Formação": "T.I.",
			"Cargo": "Devops",
		},
	}
	fmt.Println(usuario2)
	fmt.Println(usuario2["Nome"]["Prenome"])
	delete(usuario2, "Profissão")
	fmt.Println(usuario2)
	
	usuario2["Data de Nascimento"] = map[string]string{
		"Dia": "12",
		"Mes": "12",
		"Ano": "1990",
	}
	fmt.Println(usuario2)
}