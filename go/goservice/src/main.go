package main

import (
	"project/country"
	"project/states"
)

func main() {
	listCountry := country.GetListCoutry()
	country.InsertCountry(listCountry)

	states.GetListStates()
}
