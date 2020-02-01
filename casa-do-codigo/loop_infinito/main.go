package main

import "math/rand"

import "time"

import "fmt"

func main() {

	rand.Seed(time.Now().UnixNano())
	n := 0

	for {
		n++

		i := rand.Intn(4200)
		fmt.Println(i)
		if i%42 == 0 {
			break
		}
	}

	fmt.Printf("Saída após %d iterações.\n", n)
}
