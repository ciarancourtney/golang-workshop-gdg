package main

import "fmt"

func main() {

	finger := 1

	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	default:
		fmt.Println("Huh?")
	}
	
}
