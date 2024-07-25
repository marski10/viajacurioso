package main

import (
	"fmt"
	_ "os"
)

type pessoa struct {
	nome   string
	idade  int
	teste  string
	altura float64
}

type nulo interface{}

func (p pessoa) getIdade() int {
	return p.idade
}

func (p pessoa) getNome() string {

	return p.nome

}

func main() {

	var marcos = pessoa{"marcos", 20, "masculino", 1.80}

	var existe interface{}
	fmt.Println(marcos.getIdade())
	fmt.Println(existe)

}
