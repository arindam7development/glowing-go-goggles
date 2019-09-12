package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	sec := time.Now().Unix()
	rand.Seed(sec)
	target := rand.Intn(100) + 1
	fmt.Println(target)
	success := false
	for i := 0; i < 3; i++ {
		fmt.Println("Ready to guess the number, you have ", 3-i, " chances. Enter it")
		input := bufio.NewReader(os.Stdin)
		enteredValue, er := input.ReadString('\n')
		enteredValue = strings.TrimSpace(enteredValue)
		newconvValue, er2 := strconv.ParseFloat(enteredValue, 64)
		if er != nil || er2 != nil {
			fmt.Println("Got error")
		}
		if newconvValue == float64(target) {
			success = true
			fmt.Println("Yeah got it correct")
			break
		} else if newconvValue > float64(target) {
			fmt.Println("Guess is greater the target")
			//fmt.Println("The target was :", target)
		} else {
			fmt.Println("Guess is less than the target")
			//fmt.Println("The target was :", target)
		}
	}
	if !success {
		fmt.Println("You ran out of chances. ", "The target was : ", target)
	}

}
