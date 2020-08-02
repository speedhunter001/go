package main

import "fmt"


func main() {
	var num_float float64
	var num_trunc int64 

	fmt.Printf("Enter a floating point number : ")  
	fmt.Scan(&num_float)

 	num_trunc = int64(num_float)
 	fmt.Printf("\nThis is the truncated number : ")
 	fmt.Println(num_trunc)  // Println prints a \n too on its own
}

