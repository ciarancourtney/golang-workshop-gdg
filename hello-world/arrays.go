package main

import "fmt"

var print = fmt.Println


func appendSlice() {
	sa := []int{1, 2, 3}
	newSa := append(sa, []int{12, 13}...)

	fmt.Println(newSa)
}

func main() {

	appendSlice()

	
}
