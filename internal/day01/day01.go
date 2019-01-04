package day01

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

func Basement(input string) int {
  floor := 0
  for idx, charValue := range input {
    switch charValue {
    case '(':
      floor += 1
    case ')':
      floor -= 1
    }
    if floor == -1 {
      return idx + 1
    }
  }
  return 0
}
