package states

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"project/utilities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type JsonState struct {
	Id     int    `json:"id"`
	Sigla  string `json:"sigla"`
	Nome   string `json:"nome"`
	Regiao struct {
		Id int `json:"id"`
	} `json:"regiao"`
}

type State struct {
	gorm.Model
	Id     int
	Sigla  string
	Nome   string
	Regiao int
}

func GetListStates() []JsonState {

	var connect utilities.Connect = utilities.NewConnect()

	const (
		endpoint = "/state"
	)

	connect.Url = "https://a723bc1c-afd1-4c42-b7a6-5a6bc6258baf.mock.pstmn.io"
	connect.ApiKey = ""

	httpClient := &http.Client{}

	request, _ := http.NewRequest("GET", connect.Url+endpoint, nil)

	request.Header.Add("X-Api-Key", connect.ApiKey)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, _ := httpClient.Do(request)

	body, _ := ioutil.ReadAll(response.Body)

	var States []JsonState

	json.Unmarshal(body, &States)

	for _, x := range States {

		fmt.Println(x.Sigla)
		fmt.Println(x.Id)
		fmt.Println(x.Sigla)
		fmt.Println(x.Nome)
		fmt.Println(x.Regiao.Id)

	}

	return States

}

func InsertStates() {

	states := GetListStates()

	dsn := "root:my-secret-pw@tcp(some-mysql:3306)/country"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("faleid to connect database")

	}

	db.AutoMigrate(&State{})

	for _, x := range states {
		db.Create(&State{Id: x.Id, Sigla: x.Sigla, Nome: utilities.Text(x.Nome), Regiao: x.Regiao.Id})
	}

}
