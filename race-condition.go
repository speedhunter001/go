package main

import ("fmt"
		"time")


func print_hello() {
	fmt.Println("Hello ")
}

func print_world() {
	fmt.Println("World")
}


func main() {
	
	for i:=0; i<10; i++ {
		go print_hello()
		go print_world()
	}

	time.Sleep(50 * time.Millisecond)  /* If this is removed then main funnction exits before the goroutines execute.You can see that when program is executed 'Hello' and 'World' are printed in a random manner which shows the race condition when the code is not executed in sequence as written*/
}