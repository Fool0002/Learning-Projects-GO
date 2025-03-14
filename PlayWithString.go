package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	//Accept input from user when compiling
	str := os.Args[1]

	//Add ! according to the length of String
	NewStr := strings.Repeat("!", len(str)) + str + strings.Repeat("!", len(str))
	fmt.Println(strings.ToUpper(NewStr))

	//Trimming ! on the right side
	NS := strings.TrimRight(NewStr, "!")
	fmt.Println(NS)
}
