package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {

	quant := 5
	fmt.Println("Arindam Go")
	strings.Title("head first go")
	fmt.Println(reflect.TypeOf(quant), strings.Title("head first go"))
}
