package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Set the sum to 0
	sum := 0

	// Read the file
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	for _, line := range fileLines {
		// print original line
		fmt.Println(line)
		// regex to find all digits and words
		re := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|\d)`)
		// get all digits
		allDigits := re.FindAllString(line, -1)
		// fmt.Println(allDigits)
		firstDigit := re.FindAllString(line, 1)
		if len(firstDigit) > 0 {
			if intValue, ok := wordToInt(firstDigit[0]); ok {
				fmt.Println(intValue)
			} else {
				intFirstDigit, _ := strconv.Atoi(strings.Join(firstDigit, ""))
				fmt.Println(intFirstDigit)
			}
		}
		lastDigit := string(allDigits[len(allDigits)-1])
		if len(lastDigit) > 0 {
			if intValue, ok := wordToInt(lastDigit); ok {
				fmt.Println(intValue)
			} else {
				intLastDigit, _ := strconv.Atoi(lastDigit)
				fmt.Println(intLastDigit)
			}
		}

	}
	fmt.Println("Sum:", sum)
}

func wordToInt(word string) (int, bool) {
	// Mapping of number words to Intergers
	numberMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	value, exists := numberMap[strings.ToLower(word)]
	return value, exists
}
