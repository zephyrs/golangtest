package helper

import (
	"fmt"
	"log"
	"runtime"
)

//Trace start
func Trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

//Un Trace
func Un(s string) {
	fmt.Println("leaving:", s)
}

//Where is called
func Where() {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("%s:%d", file, line)
}
