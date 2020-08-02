package main

import ("bufio"
		"fmt"
		"os"
		"strconv"
		"strings")


func GenDisplaceFn(a, v0, s0 float64) func(t float64) float64 {

	return func(t float64) float64 {
					return ( (1/2) * (a * (t*t)) ) + ( v0 * t ) + s0
			}
}

func main() {
	var t float64

	fmt.Printf("Enter values for ACCELERATION, INITIAL VELOCITY and INITIAL DISPLACEMENT separated with spaces : ")
	
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	nums_inp := strings.TrimSpace(input)	// Would cause issue on last value if not done
	nums_str := strings.Split(nums_inp, " ")

	a, _ := strconv.ParseFloat(nums_str[0], 64)		// Converting to floats
	v0, _ := strconv.ParseFloat(nums_str[1], 64)
	s0, _ := strconv.ParseFloat(nums_str[2], 64)

	fmt.Printf("Enter TIME value: ")
	fmt.Scan(&t)	

	function := GenDisplaceFn(a, v0, s0)
	fmt.Println(function(t))
}		