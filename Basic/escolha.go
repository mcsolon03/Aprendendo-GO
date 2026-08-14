package main 

import (
	f "fmt"
	t"time"
)


func main() {
   t:= t.Now()


   switch {
   case t.Hour() < 12: f.Println("Bom Dia!")
   case t.Hour() < 18: f.Println("Boa Tarde!")
   default: f.Println("Boa Noite")	
}
}