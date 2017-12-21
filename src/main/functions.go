package main

import (
	"fmt"
	"io"
	"log"
	"time"
)

//Function test case
const (
	FuncFibonacci = iota
	FuncDefer
	FuncUnknown
)

func runTestFunctions() {
	where()
	defer un(trace("runTestFunctions"))

	i := FuncDefer
	switch i {
	case FuncFibonacci:
		fibonacci(25)
	case FuncDefer:
		deferLog("DEFER")
	case FuncUnknown:
		fallthrough
	default:
		fmt.Println("an unknown function to test")
	}
}

func fibonacci(n int) {
	defer un(trace("fibonacci"))

	start := time.Now()

	last1 := 1
	last2 := 1
	fib := func(i int) (result int) {
		if i < 2 {
			result = 1
		} else {
			result = last1 + last2
			last2 = last1
			last1 = result
		}
		return
	}

	for i, result := 0, 0; i <= n; i++ {
		result = fib(i)
		fmt.Printf("fibonacci(%2d) is: %6d\n", i, result)
	}

	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("fibonacci took time: %s\n", delta)
}

func deferLog(s string) (n int, err error) {
	defer func() {
		fmt.Printf("deferLog(%q) = %d, %v\n", s, n, err)
		log.Printf("deferLog(%q) = %d, %v", s, n, err)
	}()

	return 200, io.EOF
}
