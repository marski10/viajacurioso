package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"project/normalize"
	"reflect"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connect struct {
	url     string
	apiKey  string
	apiHost string
}

func newConnect() Connect {

	return Connect{
		url:     "https://countries-cities.p.rapidapi.com/",
		apiKey:  "393041ad93msh827b230c12feec5p1f4dbdjsndde28001dfc8",
		apiHost: "countries-cities.p.rapidapi.com",
	}

}

type Country struct {
	ID struct {
		ISOALPHA2 string `json:"ISO-ALPHA-2"`
	} `json:"id"`
	Nome      string `json:"nome"`
	SubRegiao struct {
		Regiao struct {
			Nome string `json:"nome"`
		} `json:"regiao"`
	} `json:"sub-regiao"`
}

func GetListCoutry() []Country {

	var apiConnect Connect = newConnect()

	apiConnect.url = "https://servicodados.ibge.gov.br/api/v1"

	const (
		endpoint = "/localidades/paises"
	)

	httpClient := &http.Client{}

	//request, _ := http.NewRequest("GET", apiConnect.url+endpoint, nil)
	request, _ := http.NewRequest("GET", "https://a723bc1c-afd1-4c42-b7a6-5a6bc6258baf.mock.pstmn.io/country", nil)

	//request.Header.Add("X-RapidAPI-Key", apiConnect.apiKey)
	//request.Header.Add("X-RapidAPI-Host", apiConnect.apiHost)

	res, _ := httpClient.Do(request)

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

	var countries []Country

	json.Unmarshal(body, &countries)

	fmt.Println(countries)

	for _, x := range countries {
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println(x.Nome)
		fmt.Println(x.ID.ISOALPHA2)
		fmt.Println(x.SubRegiao.Regiao.Nome)
		fmt.Println(strings.Repeat("=", 50))
	}

	return countries

}

type Countries struct {
	gorm.Model
	Name         string
	Abbreviation string
	Continent    string
}

func insertCountry(c []Country) {
	dsn := "root:my-secret-pw@tcp(some-mysql:3306)/country"
	db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if error != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Countries{})

	for _, x := range c {
		db.Create(&Countries{Name: normalize.Text(x.Nome), Abbreviation: normalize.Text(x.ID.ISOALPHA2), Continent: normalize.Text(x.SubRegiao.Regiao.Nome)})
	}

	fmt.Println("type:", reflect.TypeOf(c))
}