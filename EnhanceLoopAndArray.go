package main 

import "fmt"

func main() {
	 array := [5]int{0, 5, 10, 15, 20}
	var total int
	 for _, value := range array{
		  total += value
	 }
	 fmt.Println(total)
	 fmt.Println(total / len(array))
}