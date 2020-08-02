package main

import ("fmt"
		"os"
		"bufio"
		"log"
		"strings")


type full_name struct {
	first_name string
	last_name string
}


func main() {
	var file_name string
	var temp_full_name full_name
	
	fmt.Printf("Enter file name : ")
	fmt.Scan(&file_name)   // I use file names with - between words as I use linux e.g, name-file.txt

	all_full_names := make([]full_name, 0, 1)


	file, err := os.Open(file_name)	// Opening file and using its handle
	if err != nil {
	    log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)		// Passing handle to scanner 
	for scanner.Scan() {					// Reading line by line
		name_read := strings.Split(scanner.Text(), " ")
		temp_full_name.first_name, temp_full_name.last_name = name_read[0], name_read[1]

		all_full_names = append(all_full_names, temp_full_name)
	}

	if err := scanner.Err(); err != nil {
	    log.Fatal(err)
	}

	file.Close()

	fmt.Println("\nThe file contains these names....\n")
	for i := range all_full_names {
		name := all_full_names[i]
		fmt.Println(name.first_name, name.last_name)
	}
}		