package main

import "fmt"

//go:generate filestr -trim VERSION version.go myVersion

func main() {
	fmt.Println(myVersion)
}
