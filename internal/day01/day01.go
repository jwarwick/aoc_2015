package day01

import (
  // "fmt"
)

func FloorCount(input string) int {
  floor := 0
  for _, charValue := range input {
    switch charValue {
    case '(':
      floor += 1
    case ')':
      floor -= 1
    }
  }
  return floor
}
