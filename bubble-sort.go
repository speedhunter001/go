package main

import ("fmt"
		"bufio"
		"os"
		"strings"
		"strconv"
		)


func swap(array []int, index int) {
	array[index], array[index+1] = array[index+1], array[index]
}


func bubbleSort(array []int) {
	for i, _ := range(array) {
		for j:=0; j<len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				swap(array, j)
			}
		}
	}
}


func main() {
	array := make([]int, 0, 10)

	fmt.Printf("Enter 10 integers separated with space : ")

	reader := bufio.NewReader(os.Stdin)
	nums_inp, _ := reader.ReadString('\n')  // Reading whole line

	nums_trimmed := strings.TrimSpace(nums_inp)		// If this not done then while converting last index element issues are faced
	nums_str := strings.Split(nums_trimmed, " ")

	for _, element := range nums_str {		
		num, _ := strconv.Atoi(element)
		//fmt.Println(element, num)

		array = append(array, num)
	}

	fmt.Println("Inputted array : ", array)
	bubbleSort(array)
	fmt.Println("After using bubble sort array is : ", array)
}		