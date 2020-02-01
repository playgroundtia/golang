package main

import "fmt"

func max(numeros []int) int {
	var max int

	for _, numero := range numeros {
		if numero > max {
			max = numero
		}
	}

	return max
}

func testMax(numeros []int, expectativa int) string {
	saida := "Pass"

	atual := max(numeros)
	if atual != expectativa {
		saida = fmt.Sprintf("Esperador %v, mas obteve %v", expectativa, atual)
	}
	return saida
}

func main() {
	fmt.Println(testMax([]int{1, 2, 3, 4, 5}, 5))
	fmt.Println(testMax([]int{1, 2, 3, 4, 6}, 5))
	fmt.Println(testMax([]int{1, 2, 3, 4, 7}, 7))

	// Pass
	// Esperador 5, mas obteve 6
	// Pass

}
