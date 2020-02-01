package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*
		Array: tamanho fixo
		Slice: tamanho variável
	*/
	var unidadeDestino string
	numerosDeArgumentos := len(os.Args)
	ultimoElemento := numerosDeArgumentos - 1
	celsius := "celsius"
	fahrenheit := "fahrenheit"
	quilometros := "quilometros"
	milhas := "milhas"

	if numerosDeArgumentos < 3 {
		fmt.Println("Uso: conversor <valores> <unidade>")
		os.Exit(1)
	}

	// último argumento passado
	unidadeOrigem := os.Args[ultimoElemento]
	valoresOrigem := os.Args[1:ultimoElemento]

	if unidadeOrigem == celsius {
		unidadeDestino = fahrenheit
	} else if unidadeOrigem == quilometros {
		unidadeDestino = milhas
	} else {
		unidadeDestino = unidadeOrigem
		fmt.Printf("%s não é uma unidade conhecida! \n", unidadeDestino)
		os.Exit(1)
	}

	for indice, valor := range valoresOrigem {

		var valorDestino float64

		valorOrigem, err := strconv.ParseFloat(valor, 64)
		if err != nil {
			fmt.Printf("O valor %s na posição %d não é um número válido!\n", valor, indice)
			os.Exit(1)
		}

		if unidadeOrigem == celsius {
			valorDestino = convertCelsiusFahrenheit(valorOrigem)
		} else {
			valorDestino = convertQuilometrosMillhas(valorOrigem)
		}

		fmt.Printf("%.fº %s = %.fº %s\n", valorOrigem, unidadeOrigem, valorDestino, unidadeDestino)
	}

}

func convertCelsiusFahrenheit(valor float64) float64 {
	return valor*1.8 + 32
}

func convertQuilometrosMillhas(valor float64) float64 {
	return valor / 1.60934
}
