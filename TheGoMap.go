package main

import (
	"fmt"
	"sort"
	"strings"
)

func main()  {
	fmt.Println(strings.Title("head first go"))
	fmt.Printf("This is GoLand")
	input := [] string{"bob","arindam","arindam","arindam", "arindam", "ironman","ironman","ironman"}
	sort.Strings(input)
	myResultMap := make(map[string]int)
	for _, can := range input {
		myResultMap[can]++
	}

	for i, can1 := range myResultMap {
		fmt.Println(i, " ", can1 )
	}
	fmt.Printf("%v", myResultMap)

	// Sort the map by values
	for _, value :=range input {
		fmt.Println("key", value, myResultMap[value])
	}

}

