package main 

import "fmt"

func main() {
	var array [5]int
	var number int = 5
	var total int
	
	for i := 0; i < 5; i++{
		array[i] = number * i
		total += array[i]
	}
	fmt.Println(array)
	fmt.Println(total)
}