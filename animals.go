package main

import ("fmt"
		"bufio"
		"os"
		"strings")


type animals struct {
	food string
	locomotion string
	noise string
}		

func(animal animals) Eat() {
	fmt.Println(animal.food)
}

func(animal animals) Move() {
	fmt.Println(animal.locomotion)
}

func(animal animals) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	cow := animals{"grass", "walk", "moo"}
	bird := animals{"worms", "fly", "peep"}
	snake := animals{"mice", "slither", "hsss"}

	for {
		//fmt.Prinf("Enter name of animal i.e cow,snake,bird and its info i.e move,eat,speak with space separating them  ")
		fmt.Printf(">")		// Because its required

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		input_trimmed := strings.TrimSpace(input)	// It would give issue sometimes in case of array if not done but still good practice
		animal_info := strings.Split(input_trimmed, " ")
		
		// Conditions are long i.e hard-coded 
		if animal_info[0] == "cow" {

			if animal_info[1] == "eat"{
				cow.Eat()
			}else if animal_info[1] == "move"{
				cow.Move()
			}else if animal_info[1] == "speak"{
				cow.Speak()
			}

		}else if animal_info[0] == "bird" {

			if animal_info[1] == "eat"{
				bird.Eat()
			}else if animal_info[1] == "move"{
				bird.Move()
			}else if animal_info[1] == "speak"{
				bird.Speak()
			}

		}else if animal_info[0] == "snake" {

			if animal_info[1] == "eat"{
				snake.Eat()
			}else if animal_info[1] == "move"{
				snake.Move()
			}else if animal_info[1] == "speak"{
				snake.Speak()
			}
		}
	}
}