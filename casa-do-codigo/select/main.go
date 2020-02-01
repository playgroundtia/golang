package main

import "fmt"

func separar(nums []int, impar, par chan<- int, pronto chan<- bool) {
	for _, n := range nums {
		if n%2 == 0 {
			par <- n
		} else {
			impar <- n
		}
	}
	pronto <- true
}

func main() {
	impar, par := make(chan int), make(chan int)
	pronto := make(chan bool)

	nums := []int{1, 22, 34, 55, 76, 77, 89, 88, 90, 115, 117, 132}

	go separar(nums, impar, par, pronto)

	var impares, pares []int
	fim := false

	for !fim {
		select {
		case n := <-impar:
			impares = append(impares, n)
		case n := <-par:
			pares = append(pares, n)
		case fim = <-pronto:
		}
	}

	fmt.Printf("Ímpares: %v | Pares: %v\n", impares, pares)

}
