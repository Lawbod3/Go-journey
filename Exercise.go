package main 

import "fmt"

func main() {
	fmt.Print("Enter a number in farenheit: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input - 32 * 5/9
	fmt.Printf("The convertion in celsius is %f" , output)
	
}