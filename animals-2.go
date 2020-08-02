package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)



type Animal interface {
	Eat()
	Move()
	Speak()
}


type Cow struct {
	food,walk,voice string
}

func(animal Cow) Eat() {
	fmt.Println(animal.food)
}

func(animal Cow) Speak() {

	fmt.Println(animal.voice)
}
func(animal Cow) Move() {

	fmt.Println(animal.walk)
}


type Bird struct {
	food,walk,voice string
}

func(animal Bird) Eat() {
	fmt.Println(animal.food)
}

func(animal Bird) Speak() {
	fmt.Println(animal.voice)
}

func(animal Bird) Move() {
	fmt.Println(animal.walk)
}


type Snake struct {
	food,walk,voice string
}

func(animal Snake) Eat()  {
	fmt.Println(animal.food)
}

func(animal Snake) Speak()  {
	fmt.Println(animal.voice)
}

func(animal Snake) Move()  {
	fmt.Println(animal.walk)
}


func main() {
	var animal Animal

	data := make(map[string]Animal)

	for {

		fmt.Printf(">")		// Because > is required

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		full_command := strings.TrimSpace(input)		// In case of array it may give error
		command_and_details := strings.Split(full_command, " ")

		if command_and_details[0] == "newanimal" { 	

			switch command_and_details[2] {		
				case "cow":
					animal = Cow{"Grass","Walk","moo"}
				
				case "snake":
					animal = Snake{"mice","slither","hsss"}
				
				case "bird":
					animal = Bird{"worm","fly","peep"}
			}

			data[command_and_details[1]] = animal 	// ..[1] has the name so it should be mapped to name
			fmt.Println("Created it!")

		} else if command_and_details[0] == "query" {

			switch command_and_details[2] {
				case "eat":
					data[command_and_details[1]].Eat()		// ..[1] is name so thats why its used as key in map

				case "move":
					data[command_and_details[1]].Move()

				case "speak":
					data[command_and_details[1]].Speak()
			}

		}

	}
}