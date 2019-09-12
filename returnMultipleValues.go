package main

import "fmt"

func main() {
	myint, bol, name := returnMultipleValues("arindam")
	fmt.Println(myint, bol, name)
}

func returnMultipleValues(name string) (int, bool, string) {
	return 1, true, name
}
