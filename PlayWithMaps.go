package main

import "fmt"

func main() {

	mapList := map[int]int{
		1: 2,
		3: 4,
		5: 6}

	fmt.Println(mapList)
	delete(mapList, 1)
	fmt.Println(mapList)
	mapList[7] = 8
	mapList[7]++
	fmt.Println(mapList)
}
