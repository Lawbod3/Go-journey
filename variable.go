package main
import "fmt"

func main(){

	var x string 
	x = "hello world"
	fmt.Println(x)
	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)

	var y string = "first"
	fmt.Println(x == y)
	y = "second"
	fmt.Println(x == y)
}