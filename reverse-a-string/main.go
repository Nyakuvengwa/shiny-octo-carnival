package main

import (
  "fmt"
  "strings"
)

func main() {
  words := strings.Split("Hello, World!", " ")
  var swappedStrings []string

  for _, word := range words {
    newWord := swapLetters(word)
    swappedStrings = append(swappedStrings, newWord)
  }

  fmt.Println(strings.Join(swappedStrings, " "))
}

func swapLetters(word string) string {
  start, end := 0, len(word)-1
  swapped := []byte(word)
  for start < end {
    if !isLetter(swapped[start]) {
      start++
      continue
    }
    if !isLetter(swapped[end]) {
      end--
      continue
    }
    swapped[start], swapped[end] = swapped[end], swapped[start]
    start++
    end--
  }

  return string(swapped)
}

func isLetter(b byte) bool {
  return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}