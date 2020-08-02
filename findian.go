package main

import ("fmt";"strings")		// Can write on different lines too without ;


func main() {
	var str string

	fmt.Printf("Enter a string : ")
	fmt.Scan(&str)	

	str = strings.ToLower(str)

	if ( strings.HasPrefix(str, "i") && strings.HasSuffix(str, "n") && strings.Contains(str, "a") ) { 
		fmt.Println("Found!")
	} else {						// Writing else on next line will give error
		fmt.Println("Not Found!")		
	}
}