package states

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"project/utilities"
)

type states struct {
	id     int    `json:"id"`
	sigla  string `json:"sigla"`
	nome   string `json:"nome"`
	regiao struct {
		id    int    `json:"id"`
		sigla string `json:"sigla"`
		nome  string `json:"nome"`
	} `json:"regiao"`
}

func GetListStates() {

	var connect utilities.Connect = utilities.NewConnect()

	const (
		endpoint = "/state"
	)

	connect.Url = "https://a723bc1c-afd1-4c42-b7a6-5a6bc6258baf.mock.pstmn.io"
	connect.ApiKey = "apikey"

	httpClient := &http.Client{}

	request, _ := http.NewRequest("GET", connect.Url+endpoint, nil)

	request.Header.Add("X-Api-Key", connect.ApiKey)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, _ := httpClient.Do(request)

	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))

}
