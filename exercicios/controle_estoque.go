package main


import ("fmt")

func main() {
 estoque := []int{10, 0, 3, 25, 7, 1, 0, 15}
 
 var semEstoque int
 var estoqueCritico int
 var estoqueBaixo int
 var estoqueNormal int

 var maiorQuantidade = estoque[0]
 var maiorPosicao = 0

 for posicao, quantidade := range estoque {
	if quantidade == 0 {
		semEstoque++
	}else if quantidade >= 1 && quantidade <= 3 {
		estoqueCritico++
	} else if quantidade >= 4 && quantidade  <=10 {
		estoqueBaixo++
	}else{
		estoqueNormal++
	}
	
  if quantidade > maiorQuantidade{
	maiorQuantidade = quantidade
	maiorPosicao = posicao
  } 

}
fmt.Println("Total sem estoque: ", semEstoque)
fmt.Println("Total em estoque baixo: ", estoqueBaixo)
fmt.Println("Total em estoque critico: ", estoqueCritico)
fmt.Println("Total em estoque normal: ", estoqueNormal)

fmt.Println("Maior quantidade em estoque: ", maiorPosicao)
fmt.Println("Maior quantidade: ", maiorQuantidade )

}