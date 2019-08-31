package main

//package is a workspace or like a folder
//so if I had more files belonging to this workspace
//I'd need all of them to include package main

//in go there are 2 kinds of packages
//1. executable package --> generates an executable file
//2. reusabe package --> code used as 'helpers' or reusable logic

//specifically 'package main' is used to make an executable type package
//'package some-random-str' is used to make a reusabe type package. When run 'go build' it doesn't create an executable file.

import "fmt"

//if 'package main' is declared, the file must have a func called main()
func main() {
	fmt.Println("Hi there!")
}

/****************************************************
go cli

go build 		--> Compiles a bunch of go source files
go run 			--> Compiles one or two files
go fmt 			--> Formars all the code in each file in the current directory
go install 	--> Compiles and installs a package
go get 			--> Downloads the raw source code of someone elses package
go test 		--> Runs any test associated with the current project
*****************************************************/
