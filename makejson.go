package main

import ("fmt"
		"encoding/json")

type person struct {
	name string
	address string
}		

func main() {
	var p1 person

	dict := make( map[string]string )

	fmt.Printf("Enter the name of the Person : ")
	fmt.Scan(&p1.name)

	fmt.Printf("Enter the address of this Person : ")
	fmt.Scan(&p1.address)

	dict[p1.name] = p1.address
	obj, _ := json.Marshal(dict)
	fmt.Println( string(obj) )
}