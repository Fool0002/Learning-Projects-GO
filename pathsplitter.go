package main

import (
	"fmt"
	"path"
)

func main() {
	//Program to split the filepath

	var file_Name, dir_Name string
	var PATH string

	PATH = "Hero/Personality/Anger.jsp"
	dir_Name, file_Name = path.Split(PATH)
	fmt.Println("The Directory is : ", dir_Name)
	fmt.Println("The File is : ", file_Name)

	//Using shortcut variable declaration and only taking the file_name
	_, FILE_NAME := path.Split(PATH)
	fmt.Println("Another File Name : ", FILE_NAME)

}
