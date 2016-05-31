package main

import "fmt"

func main() {
  i := 2
  n := 600851475143

  for i < n {
    if n % i == 0 {
      n /= i
    } else {
      i++
    }
  }

  fmt.Println(i)
}
