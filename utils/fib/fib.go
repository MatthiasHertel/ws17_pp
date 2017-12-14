package main

import (
	"flag"
	"log"
)

// intentionally slow
func fib(n int) int {
	if n < 3 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

var n int

func main() {
	flag.IntVar(&n, "n", 10, "argument for fib")
	flag.Parse()

	for i := 1; i <= n; i++ {
		log.Printf("%v => %v\n", i, fib(i))
	}
}
