package main

import ("fmt"
		"bufio"
		"os"
		"sort"
		"strconv"
		"strings"
		"sync"
		"math")


func sort_array(array []int, wait_group *sync.WaitGroup) {
	fmt.Println("Now sorting : ", array)
	sort.Ints(array)
	fmt.Println("Sorted : ", array, "\n")
	
	wait_group.Done()  // This goroutine is done so +1 which was added is decremented
}

func sort_array_by_goroutines(array []int, wait_group sync.WaitGroup) {
	parts := 4
	num_for_subarrays := int( math.Max( math.Ceil( float64(len(array))/float64(parts) ), 1 ) )   // A number to cound from and to
	fmt.Print("h")
	
	for i:=0; i<parts; i++ {
		from := int( math.Min( float64(i*num_for_subarrays), float64(len(array)) ) )
		to := int( math.Min( float64((i+1)*num_for_subarrays), float64(len(array)) ) )

		wait_group.Add(1)	// One goroutine added
		go sort_array(array[from:to], &wait_group)  // Need to pass by reference otherwise deadlock will happen
	}

	wait_group.Wait()	// Wait for all goroutines to finish
	
	fmt.Println("Now finally sorting: ", array)
	sort.Ints(array)
	fmt.Println("Finally Sorted", array, "\n")
} 

func main() {
	fmt.Printf("Enter integers separated by space : ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	array_str := strings.Split(strings.TrimSpace(input), " ")   // If TrimSpace is not used then issues may be faced
	array := make([]int, 0, len(array_str))

	for i := range array_str {
		num, _ := strconv.Atoi(array_str[i])
		array = append(array, num)
	}

	var wait_group sync.WaitGroup	

	sort_array_by_goroutines(array, wait_group)

	fmt.Println("Sorted by goroutines : ", array)
}		