package main

import "fmt"

func main(){
 list := []float64{1,2,3,4,5,6}

 
  fmt.Println(average(list))


}

func average(variable []float64) float64{
	var total float64
	var answer float64
	for _, value := range variable{
		total +=  value
	}
	answer = total / float64(len(variable))
	return answer
 }