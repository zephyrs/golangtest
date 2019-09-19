package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// IO
const (
	IOInput = iota
	IOFile
)

// NewLine ...
var NewLine = flag.Bool("n", false, "print newline") // echo -n flag, of type *bool

func runTestIOs() {
	defer un(trace("runTestIOs"))

	flag.PrintDefaults()
	flag.Parse() // Scans the arg list and sets up flags

	var s string
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *NewLine { // -n is parsed, flag becomes true
				s += "Newline"
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)

	i := IOInput
	switch i {
	case IOInput:
		tryInput()
	case IOFile:
		fallthrough
	default:
		fmt.Println("unknown test item")
	}
}

func tryInput() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	fmt.Printf("Your name is %s", input)
}
