package main

import "fmt"

func main() {
	x := []int{
	  48,96,86,68,
	  57,82,63,70,
	  37,34,83,27,
	  19,97, 9,17,
	}

	min:=min(x)
	
	fmt.Printf("El numero minimo es: %d\n",min)
}

func min(x []int) int{
	var min int
	min = x[0]

	for _, e := range x {
		if(e<min){
			min = e
		}
	}

	return min
}