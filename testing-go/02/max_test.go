package main

import "testing"

func Max(numeros []int) int {
	if len(numeros) == 0 {
		return -1
	}

	var max int

	for _, numero := range numeros {
		if numero > max {
			max = numero
		}
	}

	return max
}

func TestMaxInvalid(t *testing.T) {
	actual := Max([]int{1, 2, 3, 4})
	if actual != 4 {
		t.Errorf("Expected %d, got %d", 4, actual)
	}
}

func TestMaxEmpty(t *testing.T) {
	actual := Max([]int{})
	if actual != -1 {
		t.Errorf("Expected %v, got %d", -1, actual)
	}
}
