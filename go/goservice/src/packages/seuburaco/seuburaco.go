package seuteste

import "fmt"

type seuteste interface {
	Printpint()
}

type pint struct {
	Nome string
}

func (a pint) Meuteste() {

	fmt.Println("meuteste")

}
