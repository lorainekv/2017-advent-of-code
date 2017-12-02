package main

import (
  "strings"
  "strconv"
  "bufio"
  "fmt"
  "os"
)

func main() {
  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()

  sum := 0
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()
    diff := difference(line)
    sum += diff
  }
  fmt.Println(sum)
}

func difference(input string) int {
  slice := strings.Fields(input)
  firstValue := slice[0]

  // Largest and smallest values are initially 
  // the first value
  largest, _ := strconv.Atoi(firstValue)
  smallest, _ := strconv.Atoi(firstValue)

  for _, num := range slice {
    current, _ := strconv.Atoi(num)

    if current > largest {
      largest = current
    }

    if current < smallest {
      smallest = current
    }
  }

  diff := largest - smallest
  return diff
}
