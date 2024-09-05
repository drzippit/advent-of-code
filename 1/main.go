package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
  "strconv"
  "strings"
)

func main(){
  filePath := os.Args[1]
  readFile, err := os.Open(filePath)

  sum := 0

  if err != nil {
    fmt.Println(err)
  }
  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)
  var fileLines []string

  for fileScanner.Scan(){
    fileLines = append(fileLines, fileScanner.Text())
  }

  readFile.Close()

  for _, line := range fileLines{
    // print original line
    // fmt.Println(line)
    // regex to find all digits
    re := regexp.MustCompile("[0-9]")
    // get all digits
    allDigits := re.FindAllString(line, -1)
    // get only first digit
    firstDigit := re.FindAllString(line, 1)
    // get only last digit
    lastDigit := string(allDigits[len(allDigits)-1])
    // convert []string to int
    intFirstDigit, err := strconv.Atoi(strings.Join(firstDigit, ""))
    if err != nil {
      panic(err)
    }

    // convert string to int
    intLastDigit, err := strconv.Atoi(lastDigit)
    if err != nil {
      panic(err)
    }
    //combine Ints into a single number
    combinedInts, err := strconv.Atoi(fmt.Sprintf("%d%d", intFirstDigit, intLastDigit))
    if err != nil {
      panic(err)
    }
    // fmt.Println(combinedInts)
    
    sum += combinedInts

  }
  fmt.Println("Sum:", sum)
}
