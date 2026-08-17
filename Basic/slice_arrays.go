package main

import (
	"fmt"
	"slices"
)


func main() {
//Array
//numeros := [5]int {10,20,30,40}
//numeros[4] = 50
//fmt.Println(numeros)
//Slice
frutas := []string {"Laranja", "Pera", "Uva"}
frutas = append(frutas, "Melancia")
frutas = append(frutas, "Caju")
frutas =append(frutas, "Abacaxi")
fmt.Println("Frutas", frutas)
fmt.Println("Tamanho do array: ", len(frutas))
fmt.Println("Apresentar 02 primeiros: ", frutas[:2])
fmt.Println("Apresentar a partir das posiçoes: ", frutas[3:])
//frutas = slices.Delete(frutas, 0, 1)
//fmt.Println("Sem a uva: ", frutas)
slices.Sort(frutas)
fmt.Println(frutas) //ordem alfabetica
fmt.Println("Tem  abacaxi?: ", slices.Contains(frutas, "Abacaxi"))
slices.Reverse(frutas)
fmt.Println(frutas) //ordem reverse
frutas = slices.Insert(frutas, 3, "Abacate")
fmt.Println(frutas)
}