package main

import (
	"fmt"
)

//Function test case
const (
	Fibonacci = iota
)

func runTestFunctions() {
	defer un(trace("runTestFunctions"))

	i := Fibonacci
	switch i {
	case Fibonacci:
		fibonacci(10)
	}
}

func fibonacci(n int) {
	defer un(trace("fibonacci"))
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
		fmt.Printf("fibonacci(%2d) is: %4d\n", i, result)
	}
}
