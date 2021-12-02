package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput() []string {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		log.Fatalf("failed to open file")
	}

	scanner := bufio.NewScanner(file)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text
}

func strToInt(s string) int {
	val, err := strconv.Atoi(s)

	if err != nil {
		fmt.Println("Could not convert string to int", err)
	}
	return val
}

func answerOne(input []string) int {
	var horizontal int
	var depth int

	for _, value := range input {
		data := strings.Fields(value)

		switch direction := data[0]; direction {
		case "forward":
			horizontal += strToInt(data[len(data)-1])
		case "down":
			depth += strToInt(data[len(data)-1])
		case "up":
			depth -= strToInt(data[len(data)-1])
		default:
			fmt.Println("Could not translate instructions")
		}
	}
	return horizontal * depth
}

func answerTwo(input []string) int {
	var horizontal int
	var depth int
	var aim int

	for _, value := range input {
		data := strings.Fields(value)

		switch direction := data[0]; direction {
		case "forward":
			horizontal += strToInt(data[len(data)-1])
			depth += (strToInt(data[len(data)-1]) * aim)
		case "down":
			aim += strToInt(data[len(data)-1])
		case "up":
			aim -= strToInt(data[len(data)-1])
		default:
			fmt.Println("Could not translate instructions")
		}
	}
	return horizontal * depth
}

func main() {
	inputData := readInput()

	firstAnswer := answerOne(inputData)
	secondAnswer := answerTwo(inputData)
	fmt.Println(firstAnswer, secondAnswer)
}
