package main

import (
	"fmt"
)

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var searchNum int
	fmt.Println("Enter the number you want to search")
	fmt.Scan(&searchNum)

	//Initializing first and last
	last := len(arr) - 1
	first := 0

	for first <= last {
		mid := (first + last) / 2

		if searchNum == arr[mid] {
			fmt.Println("The number found at index : ", mid)
			return
		}

		if searchNum < arr[mid] {
			last = mid - 1
		}

		if searchNum > arr[mid] {
			first = mid + 1
		}

	}
	fmt.Println("Number not Found.")
}
