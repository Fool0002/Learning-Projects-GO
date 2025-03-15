package main

import "fmt"

func main() {
	var (
		Name   string = "Ramu"
		Age    int    = 18
		Course string = "Science"
		Grade  string = "A+"
		Pass   bool   = true
	)

	//Argument indexing is done with %[index]v
	fmt.Printf("Student %v has scored %v in course %v at the small age of %v. But the question still remainson if he has passed? %v. He has passed which is obvious and achieved such feat at the age of %[4]v",
		Name, Grade, Course, Age, Pass)
}
