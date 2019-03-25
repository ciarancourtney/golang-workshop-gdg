package main

import "fmt"

func main() {
	const a = "A"
	m := map[string]int{
		a: 1,
		"b": 2,
		"c": 3,
	}
	for key, value := range m {
		fmt.Println("Key:", key, " Value:", value)
	}

	if m["f"] == 1 {
		fmt.Println(true)
	}


	//n := map[string][]int{
	//	"a": [2]int{1,2},
	//}

}