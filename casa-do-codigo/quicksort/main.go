package main

import "os"

import "fmt"

import "strconv"

func main() {
	entrada := os.Args[1:]
	numeros := make([]int, len(entrada))

	for indice, valor := range entrada {
		numero, err := strconv.Atoi(valor)
		if err != nil {
			fmt.Printf("%s não é um número válido\n", valor)
			os.Exit(1)
		}
		numeros[indice] = numero
	}

	fmt.Println(quicksort(numeros))
}

func quicksort(numeros []int) []int {
	if len(numeros) <= 1 {
		return numeros
	}

	copiaNumeros := make([]int, len(numeros))
	copy(copiaNumeros, numeros)
	indicePivo := len(copiaNumeros) / 2
	pivo := copiaNumeros[indicePivo]
	copiaNumeros = append(copiaNumeros[:indicePivo], copiaNumeros[indicePivo+1:]...)

	menores, maiores := particionar(copiaNumeros, pivo)

	return append(
		append(quicksort(menores), pivo), quicksort(maiores)...,
	)
}

func particionar(numeros []int, indice int) (menores []int, maiores []int) {

	for _, numero := range numeros {
		if numero <= indice {
			menores = append(menores, numero)
		} else {
			maiores = append(maiores, numero)
		}
	}

	return menores, maiores
}
