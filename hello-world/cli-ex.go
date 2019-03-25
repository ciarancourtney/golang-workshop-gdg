package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	cmdArgs()
}

func cmdArgs() {
	InputString := strings.ToLower(os.Args[1])
	fmt.Println("Input is:", InputString)

	var VowelCount uint
	//var DigitCount uint

	for i, char := range InputString {
		fmt.Println(i, char)

		switch char {
		case 'a', 'e', 'i', 'o', 'u':
			VowelCount++
		//case char >= '0' && char <= '9':
		//	DigitCount++
		}

	}

	flag.VisitAll(func (f *flag.Flag) {
		if f.Value.String() == "vowels" {
			fmt.Println(f.Name, "not set!")
		}
	})

	vowels := flag.Bool("vowels", false, "vowels")
	flag.Parse()

	if *vowels {
		fmt.Println("vowels:", VowelCount)
	}

}

