package main 

import "fmt"

func main() {
	
	for i := 1; i <= 10; i++{
		switch i{
		case 1: fmt.Println("one")
		case 2: fmt.Println("two")
		case 3: fmt.Println("three")
		case 4: fmt.Println("four")
		default : fmt.Println("unknown")
		}
		
	}
}