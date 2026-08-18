package main 

import("fmt")

func main() {
 vendas := []float64{350, 120, 980, 450, 75, 1500, 620, 300, 2100, 180}

 var pequeno int
 var medio int
 var grande int
 var premium int

var soma float64
var comissao float64

var menorVenda = vendas[0]
var maiorVenda = vendas[0]


 for i := 0; i < len(vendas); i++ {

    venda := vendas[i]

    if venda <= 200 {
        pequeno++
    } else if venda <= 500 {
        medio++
    } else if venda <= 1000 {
        grande++
    } else {
        premium++
	    comissao += venda * 0.05
}

		soma += venda

		if venda < menorVenda {
			menorVenda = venda
}

		if venda > maiorVenda {
			maiorVenda = venda
}

     ticketMedio := soma / float64(len(vendas))

 
	fmt.Println("Serviços pequenos: ", pequeno)	
	fmt.Println("Serviços médios:", medio)
	fmt.Println("Serviços grandes:", grande)
	fmt.Println("Serviços premium:", premium)
	fmt.Println("Ticket médio:", ticketMedio)
	fmt.Println("Menor venda:", menorVenda)
	fmt.Println("Maior venda:", maiorVenda)
	fmt.Println("Faturamento total:", soma)
	fmt.Println("Total de comissão:", comissao)
 }



}