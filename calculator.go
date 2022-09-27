package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(prompt string) float64 {
	fmt.Printf("%v", prompt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message, _ := fmt.Scanf("%v must be number only", prompt)
		panic(message)
	}
	return value
}

func add(value1, value2 float64) float64 {
	return value1 + value2
}

func subtract(value1, value2 float64) float64 {
	return value1 - value2
}

func multiply(value1, value2 float64) float64 {
	return value1 * value2
}

func divide(value1, value2 float64) float64 {
	return value1 / value2
}

func getOperator() string{
	fmt.Print("should your operation (+ - * /) : ") 
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func main() {
	fmt.Println("Welcome to calculator")
	value1 := getInput("value1 : ")
	value2 := getInput("value2 : ")

	var result float64

	switch operator := getOperator(); operator {
	case "+" : result = add(value1, value2)
	case "-" : result = subtract(value1, value2)
	case "*" : result = multiply(value1, value2)
	case "/" : result = divide(value1, value2)
	}
	fmt.Printf("result is  %v", result)
}