package main

import (
	"fmt"
	"os"
	"strconv"
)
import (
	"github.com/Leozamboni/Go/Screen"
	"github.com/Leozamboni/Calc"
)

var manter float64

func main() {
	for {
		screen.Clear()
		var (
			primeiro,
			segundo,
			resultado 	float64
			operador,
			aux 		string
		)
		if manter != 0 {
			primeiro = manter
		}
		fmt.Println("\tGolang calculadora!")
		fmt.Println()
		fmt.Println("[S] [L] [  /  ]", "     Legenda:", "                     Como usar:")
		fmt.Println("[7] [8] [9] [*]", "     ^ - Potência", "                 Digite 'N¹' 'espaço ou enter' 'Operador' 'espaço ou enter' 'N²'")
		fmt.Println("[4] [5] [6] [-]", "     . - Operações com vírgula", "    Exemplo: 1 * 2")
		fmt.Println("[1] [2] [3] [+]", "     L - Limpar")
		fmt.Println("[  0  ] [.] [^]", "     S - Sair")
		if manter == 0 {
			fmt.Scan(&aux)
			primeiro, _ = strconv.ParseFloat(aux, 10)
			if aux == "s" {
				operador = "s"
			} else if aux == "l" {
				operador = "l"
			} else {
				fmt.Scan(&operador, &segundo)
			}
		} else {
			fmt.Print(primeiro, " ")
			fmt.Scan(&operador)
			if operador == "s" {
			} else if operador == "l" {
			} else {
				fmt.Scan(&segundo)
			}
		}
		switch operador {
		case "+":
			resultado = calc.Somar(primeiro, segundo)
		case "-":
			resultado = calc.Subtrair(primeiro, segundo)
		case "*":
			resultado = calc.Multiplicar(primeiro, segundo)
		case "/":
			resultado = calc.Divisao(primeiro, segundo)
		case "^":
			resultado = calc.Potenciacao(primeiro, segundo)
		case "s":
			os.Exit(0)
		case "l":
			manter = 0
		}
		fmt.Println("Resultado:", resultado)
		manter = resultado
	}
}
