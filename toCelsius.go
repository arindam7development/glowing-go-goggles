package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("##################################################")
	fmt.Println("Convert the temp from Farenheit to Celsius")
	fmt.Print("Enter the temp ")
    temperature_f, err := getInput()
    if err!=nil {
    	log.Fatal("Something went wrong")
	}
	celsius := (temperature_f-32)*5/9
	fmt.Printf("%0.2f tempearture in cel", celsius)
}

func getInput() (float64, error) {
	input := bufio.NewReader(os.Stdin)
	number_temp, err := input.ReadString('\n')
	if err!=nil {
		log.Fatal(err)
		return 0, err
	}
	number_temp = strings.TrimSpace(number_temp)
	temp, err := strconv.ParseFloat(number_temp, 64)
	if err!=nil {
		log.Fatal(err)
		return 0, err
	}
	return temp, nil
}
