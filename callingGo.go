package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	var now time.Time = time.Now()
	year := now.Year()
	fmt.Println(year)
	Replacement()
}

// String replacment
//Go doesn't allow the users to simply decalare the vars and then don't use them
func Replacement() {
	broken := "G# r#cks"
	fmt.Println("the broken string", broken)
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println("Go has replaced #s with os .. ", fixed)
}
