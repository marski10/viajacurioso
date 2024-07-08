package main

import (
	"fmt"
	"net/http"
)

type get interface {
	url string
	endpoint string;
	connect() json;
}

func main() {

	const(
		url = "XXXX"	
		endpointPais = ""
		endpointEstado = ""
		endpointCidade =  ""
	)



	var httpClient http.Client = http.Client{}


}


