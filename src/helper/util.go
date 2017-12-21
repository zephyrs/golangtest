package helper

import "fmt"

//Trace start
func Trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

//Un Trace
func Un(s string) {
	fmt.Println("leaving:", s)
}
