package main

import "fmt"

func fib(prev int, current int, limit int) int {
  next := prev + current
  result := 0

  if current % 2 == 0 {
    result = current
  }

  if next < limit {
    return result + fib(current, next, limit)
  } else {
    return result
  }
}

func main() {
  res := fib(1, 2, 4000000)
  fmt.Println(res)
}
