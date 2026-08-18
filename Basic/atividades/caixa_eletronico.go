package main


import ("fmt")

func main() {
	saldo := 1500.0 
	movimentacoes := []float64{-200, 500, -100, -1200, 300, -900}
 

	var saques int
	var depositos int
	var saquesRecusados int
	var totalMovimentado float64

	for i := 0; i < len(movimentacoes); i++ {

		movimentacao := movimentacoes[i]
    
		if movimentacao > 0 {
			saldo += movimentacao
			depositos++
         totalMovimentado += movimentacao

		} else {
           saque := -movimentacao

			if saque <= saldo {
            saldo -= saque
				saques++
			totalMovimentado += saque

			} else {
              saquesRecusados++
			}
		}
	}

	fmt.Println("Saldo final: ", saldo)
	fmt.Println("Saques realizados: ", saques)
	fmt.Println("Quantidade de depósitos: ", depositos)
	fmt.Println("Saques recusados: ", saquesRecusados)
	fmt.Println("Total efetivamente movimentado: ", totalMovimentado)
}
