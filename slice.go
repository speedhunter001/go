package main

import ("fmt"
		"strconv"
		"sort")


func main() {
	var num_str string
	var index int

	slice := make([]int, 3)

	for {
		fmt.Printf("Enter an integer to append to slice : ")
		fmt.Scan(&num_str)

		if num_str == "X" || num_str == "x" {  // Condition to break iterations
			fmt.Println("Exiting Iteration...")
			break
		}

		num, err := strconv.Atoi(num_str) 		// If err then it means it was neither X/x nor any integer
		if err != nil {		
			fmt.Println("Please type the number correctly")
			continue
		}

		if index < 2 {				// Should fill the 3 elemnts first
			slice[index] = num			
		} else {
			slice = append(slice, num)
		}

		index = index + 1   // When index < 2 no need to increment as its not used further but still lets do
		sort.Ints(slice)
		fmt.Println(slice)
	}
}