package main

import("fmt")

func main()  {
 limite := 80 
 velocidades := []int{72, 80, 95, 110, 65, 84, 130, 78, 90, 81}
	
 var  dentro int
 var  leves int
 var  graves int
 var  gravissimas int

 var maior = velocidades[0]
 var posicaoMaior = 0

for i := 0; i < len(velocidades); i++ {

    velocidade := velocidades[i]

    if velocidade <= limite {
        dentro++
    } else if velocidade <= 96 {
        leves++
    } else if velocidade <= 120 {
        graves++
    } else {
        gravissimas++
    }

    if velocidade > maior {
        maior = velocidade
        posicaoMaior = i
    }
}


 fmt.Println("Total dentro do limite: ", dentro)
 fmt.Println("Total de infraçoes leves: ", leves)
 fmt.Println("Total de infraçoes graves: ", graves)
 fmt.Println("Total de infraçoes gravissimas: ", gravissimas)
 fmt.Println("Maior velocidade: ", maior)
 fmt.Println("Posiçao da maior velocidade: ", posicaoMaior)
}