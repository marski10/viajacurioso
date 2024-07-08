package main

import(
"fmt"
"maps"
)

func main(){
	m := make(map[string]string)
	
	m["firstName"] = "Marcos"
	m["lastName"] = "Borges"
	

	fmt.Println("map with appresentation :", m)




	delete(m,"firstname")

n := map[string]string{"foo":"1","bar": "2" }

if maps.Equal(m,n){
	fmt.Println("m == n")
}else{

	fmt.Println("this very hard")
}

}

