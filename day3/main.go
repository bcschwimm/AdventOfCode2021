package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func convertBinary(str string) int {
	i, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		fmt.Println("binary conversion error", err)
	}
	// convert from int64
	return int(i)
}

// compressBits parses a slice in column
// order and returns a reshaped slice
// with the columns as rows
func compressBits(data []string) (bitArrays []string) {
	for i := 0; i < len(data[0]); i++ {
		var s string
		for _, binaryString := range data {
			s += string(binaryString[i])
		}
		bitArrays = append(bitArrays, s)
	}
	return bitArrays
}

// mostCommonBit parses a string and returns
// a map of the number of occurances of 0 and 1
func mostCommonBit(bitData string) map[string]int {
	m := make(map[string]int)

	for _, char := range bitData {
		m[string(char)]++
	}
	return m
}

// extractValues builds a bit string after
// looking at a map to determine most often
// occurance of 1 or 0
func extractValues(m map[string]int) string {
	valZero, valOne := m["0"], m["1"]

	if valZero > valOne {
		return "0"
	}
	return "1"
}

// inverseValues returns the opposite
// bit used in the string i.e. 1 becomes 0
func inverseValues(s string) (epsilon string) {
	for _, char := range s {
		if string(char) == "0" {
			epsilon += "1"
		} else {
			epsilon += "0"
		}
	}
	return epsilon
}

func answerOne(problemData []string) int {
	var gamma string

	bitArray := compressBits(problemData)

	for _, bit := range bitArray {
		m := mostCommonBit(bit)
		gamma += extractValues(m)
	}
	epsilon := inverseValues(gamma)

	return convertBinary(gamma) * convertBinary(epsilon)
}

// func answerTwo(problemData []string) []string {
// 	var newArr []string
// 	bitArray := compressBits(problemData)

// 	for i := 0; i < len(problemData[0]); i++ {
// 		// most common bit from current index position
// 		mostCommon := extractValues(mostCommonBit(bitArray[i]))
// 		for _, bitStr := range problemData {
// 			if mostCommon == string(bitStr[0]) {
// 				newArr = append(newArr, bitStr)
// 			}
// 		}
// 	}
// 	return newArr
// }

func main() {
	data := readInput()
	firstAnswer := answerOne(data)
	fmt.Println(firstAnswer)
}
