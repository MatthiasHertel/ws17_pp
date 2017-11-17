package main

import ( "fmt" )

//
func fib(n int) int {
  if n<3 {
    return 1
  } else {
    return fib(n-1)+fib(n-2)
  }
}

func main() {
  for i := 1; i <= 10; i++ {
    fmt.Printf("%v => %v\n", i, fib(i))
  }
}
