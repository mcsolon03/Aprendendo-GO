package main

import ("fmt")



func main() {

notas := []float64{7.5, 4.0, 8.2, 6.5, 9.0, 3.5, 5.0, 7.0, 10.0, 6.0}

var aprovados int
var reprovados int
var recuperacao int

var soma float64
var menorNota = notas[0]
var maiorNota = notas[0]


    for _, nota := range notas {
        if nota >= 7 {
            aprovados++
        } else if nota >= 5 {
            recuperacao++
        } else {
            reprovados++
        }

        soma += nota

        if nota < menorNota {
            menorNota = nota
        }

        if nota > maiorNota {
            maiorNota = nota
        }
    }

    media := soma / float64(len(notas))

    fmt.Println("Alunos aprovados:", aprovados)
    fmt.Println("Alunos reprovados:", reprovados)
    fmt.Println("Alunos recuperação:", recuperacao)
    fmt.Println("Menor nota:", menorNota)
    fmt.Println("Maior nota:", maiorNota)
    fmt.Println("Média:", media)
}
