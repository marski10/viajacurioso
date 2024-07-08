package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

type Connect struct {
	url     string
	apiKey  string
	apiHost string
}

func newConnect() Connect {

	return Connect{
		url:     "url",
		apiKey:  "key",
		apiHost: "host",
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

	apiConnect.url = "endpoint"

	const (
		endpoint = "/localidades/paises"
	)

	httpClient := &http.Client{}

	//request, _ := http.NewRequest("GET", apiConnect.url+endpoint, nil)
	request, _ := http.NewRequest("GET", "endpoint", nil)

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

// func getDetailCountry(codeCountry string) {

// 	var apiConnect Connect = newConnect()

// 	const (
// 		endpoint = ""
// 	)

// }

type Countries struct {
	gorm.Model
	Name         string
	Abbreviation string
	Continent    string
}

func insertCountry(c Country) {
	dsn := "user:password@tcp(127.0.0.0:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if error != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Countries{})

	db.Create(Countries{Name: c.Nome, Abbreviation: c.ID.ISOALPHA2, Continent: c.SubRegiao.Regiao.Nome})

}

func main() {
	var listCountry []Country = GetListCoutry()

	for _, x := range listCountry {
		insertCountry(x)
	}
}
