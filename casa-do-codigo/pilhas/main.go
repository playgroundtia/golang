package main

import "errors"

import "fmt"

// Pilha pilha
type Pilha struct {
	valores []interface{}
}

// Tamanho tamanho
func (pilha Pilha) Tamanho() int {
	return len(pilha.valores)
}

// Vazia vazia
func (pilha Pilha) Vazia() bool {
	return pilha.Tamanho() == 0
}

// Empilhar empilhar
func (pilha *Pilha) Empilhar(valor interface{}) {
	pilha.valores = append(pilha.valores, valor)
}

// Desempilhar desempilhar
func (pilha *Pilha) Desempilhar() (interface{}, error) {
	if pilha.Vazia() {
		return nil, errors.New("Pilha vazia")
	}

	valor := pilha.Peek()
	pilha.valores = pilha.valores[:pilha.Tamanho()-1]

	return valor, nil
}

// Peek peek
func (pilha Pilha) Peek() interface{} {
	ultimoValor := pilha.Tamanho() - 1
	return pilha.valores[ultimoValor]
}

// SimNao simnao
func SimNao(condicional bool) string {
	var valor string

	if condicional {
		valor = "SIM"
	} else {
		valor = "Não"
	}
	return valor
}

func main() {

	pilha := Pilha{}
	fmt.Println("Pilha criada com tamanho", pilha.Tamanho())
	fmt.Println("Vazia?", SimNao(pilha.Vazia()))

	pilha.Empilhar("Golang")
	pilha.Empilhar(2020)
	pilha.Empilhar("Aracaju")
	pilha.Empilhar(3.1456789)

	fmt.Println("Tamanho após empilhar valores:", pilha.Tamanho())
	fmt.Println("Vazia?", SimNao(pilha.Vazia()))

	for !pilha.Vazia() {
		valor, _ := pilha.Desempilhar()
		fmt.Println("Desempilhando ", valor)
		fmt.Println("Tamanho: ", pilha.Tamanho())
		fmt.Println("Vazia?", SimNao(pilha.Vazia()))
	}

	_, err := pilha.Desempilhar()
	if err != nil {
		fmt.Println("Erro: ", err)
	}
}
