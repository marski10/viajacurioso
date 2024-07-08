package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type marcos interface {
	Meuteste()
}

type cara struct {
	teste string
	saco  int
	idade int
}

func (c cara) Meuteste() {

	fmt.Println("teste")

}

func seuteste2(a marcos) {

	fmt.Printf(" meu teste no seu %T", a)
	fmt.Println("seuteste")

}

func healthCheck() string {

	const (
		url      = "https://google.com"
		endpoint = "/health"
	)

	response, _ := http.Get(url + endpoint)
	body, _ := io.ReadAll(response.Body)
	response.Body.Close()

	fmt.Println(&response)

	return string(body)
}

func testemul() (string, string) {

	return "s", "ss"

}

func variacao(a ...int) {

	fmt.Printf("%d \n", a[1])
	fmt.Printf("%d \n", a[0])
	fmt.Printf("%d \n ", a[3])

}

func slic(a ...int) {

	var chsco []int

	chsco = append(chsco, 1)

	chsco = append(chsco, 3)

	fmt.Println(len(chsco))

}

func mass() {

	claudio := make(map[string]int)

	claudio["x"] = 1
	claudio["teste"] = 2

	fmt.Println(claudio)

	if value, existe := claudio["x"]; existe && value == 1 {

		fmt.Println("cara")

	} else {
		fmt.Println("aa")
	}

}

var waitGroup sync.WaitGroup

var m sync.Mutex
var xxt int

func main() {
	/* 	var response = "nothing here"

	   	var xxx *string = &response

	   	*xxx = "meuteste"

	   	//response = healthCheck()
	   	fmt.Println(response)
	   	fmt.Println(xxx)
	   	teste := *xxx

	   	fmt.Println(teste)
	   	fmt.Println(&teste)

	   	xx, teste := testemul()

	   	fmt.Println(xx)
	   	fmt.Println(teste)

	   	variacao(1, 2, 3, 4)
	   	//fmt.Printf("Response: %s", response)

	   	aa := 0

	   	pipiu := func() int {

	   		aa += 2

	   		return aa
	   	}

	   	fmt.Println(pipiu())
	   	fmt.Println(pipiu())
	   	fmt.Println(pipiu())
	   	fmt.Println(pipiu())
	   	fmt.Println(pipiu())
	   	fmt.Println(pipiu())

	   	slic()

	   	mass()

	   	amaeu := cara{"teste", 1, 2}

	   	fmt.Println(amaeu.teste)
	   	testeJson()

	   	xupera := cara{"22", 1, 2}
	   	seuteste2(xupera)

	   	popo := seuteste.pint{"marcos"}
	   	seuteste2(popo)

	   	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))

	   	//waitGroup.Add(2)
	   	//
	   	//var teste chan
	   	//
	   	//
	   	////go xoteste("PRETA", 100)
	   	////go xoteste("BRANCA", 100)
	   	//
	   	//waitGroup.Wait()
	   	fmt.Println("found xxts ", xxt)
	*/
	tempo := time.Now().UnixNano() / 1e6
	//for nx := 0; nx < 10000; nx++ {
	//	printachora()
	//}

	p := new(int)

	tempo2 := time.Now().UnixNano() / 1e6
	fmt.Println(tempo2 - tempo)
	fmt.Println(*p)
}

func printachora() {
	ovario := make(chan string)
	ok := make(chan bool)
	go func() {

		for i := 0; i <= 10; i++ {
			ovario <- fmt.Sprintf("estacao teste:%d  ", i)

		}

		ok <- true

	}()

	go func() {

		ovario <- "estacao teste"

		ok <- true
	}()

	go func() {

		ovario <- "etestacao teste"

		ok <- true
	}()

	go func() {
		<-ok
		<-ok
		<-ok
		close(ovario)
		close(ok)
	}()

	for number := range ovario {
		fmt.Printf(number)
	}

}

func xoteste(elasenta string, numerodesentadas int) {

	for x := 0; x < numerodesentadas; x++ {

		//	m.Lock()
		var xpto int
		//	m.Unlock()
		fmt.Println("SENOCA ", elasenta)
		m.Lock()
		xpto = xxt
		xpto++
		fmt.Println("SENICA ", elasenta)
		xxt = xpto
		//	m.Unlock()
	}

	waitGroup.Done()

}

type Marcos struct {
	Servidor  string
	Porta     string `json:"amaeu"`
	Reiniciar bool
	Status    string
}

func testeJson() {

	cara := Marcos{"marcos", "2001", true, "online"}
	server := Marcos{}
	fmt.Println((cara.Porta))
	result, err := json.Marshal(cara)

	//xpto := []byte(`{"Servidor":"Ana","amaeu":"2001","Reiniciar":false,"Status":"offline"}`)

	var test any

	test = "meuteste"

	fmt.Println(test)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))

	json.Unmarshal(result, &server)

	fmt.Println(result)
	fmt.Println(server)
}
