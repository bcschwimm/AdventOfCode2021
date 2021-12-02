package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// readInput reads our text file
// and returns an array of ints
func readInput() []int {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		log.Fatalf("failed to open file")
	}

	scanner := bufio.NewScanner(file)
	var text []int

	for scanner.Scan() {
		text = append(text, strToInt(scanner.Text()))
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

// answerOne uses input data and returns the number
// of instances where the depth has increased vs
//the previous measurement
func answerOne(inputData []int) int {
	var depthIncrease int

	for i := 0; i < len(inputData)-1; i++ {
		if inputData[i+1] > inputData[i] {
			depthIncrease++
		}
	}
	return depthIncrease
}

// answerTwo calculates instances where depth
// has increased based on a 3 item sliding window
func answerTwo(inputData []int) int {
	var depthIncrease int

	for i := 0; i < len(inputData)-3; i++ {
		if (inputData[i+3] + inputData[i+2] + inputData[i+1]) > (inputData[i+2] + inputData[i+1] + inputData[i]) {
			depthIncrease++
		}
	}
	return depthIncrease
}

func main() {
	problemData := readInput()
	firstAnswer := answerOne(problemData)
	fmt.Println(firstAnswer)
	secondAnswer := answerTwo(problemData)
	fmt.Println(secondAnswer)
}
