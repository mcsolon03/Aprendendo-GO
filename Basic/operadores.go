package main


import (
	f "fmt"
r "reflect"
)


func main() {
	var b byte = 3
	f.print("Tipo variavel b", r.TypeOf(b))
	 i := 3
	 f.Println ("Tipo variavel i", r.TypeOf(i)) 
}