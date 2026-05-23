package main

import "fmt"

func main() {
	fmt.Println("String de uma linha")

	fmt.Println(`
		string de 
		varias linhas
	`)

	fmt.Println("Jonathas tem ", len("Jonathas"), "letras")

	fmt.Println("Jonathas"[0]) // byte, int de 8 bits
	fmt.Println(string("Jonathas"[0]))
}
