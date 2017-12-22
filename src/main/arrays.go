package main

import "fmt"
import "bytes"

//Array test case
const (
	ArrayAppend = iota
	ArrayInsert
	ArrayRemove
	ArrayFilter
)

func runTestArrays() {
	defer un(trace("runTestArrays"))

	i := ArrayFilter
	switch i {
	case ArrayAppend:
		src := []byte{'p', 'o', 'e', 'm'}
		data := src[3:4]
		data = append(data, make([]byte, 10-len(data))...)
		data = src[:2]
		begin := 7
		sl := make([]byte, begin, 10)
		sl[begin-1] = '-'
		sl = appendBytes(sl, data)
		fmt.Printf("new string: %s, len: %v, cap: %v\n", sl, len(sl), cap(sl))
	case ArrayInsert:
		src := "abcdedf"
		data := "xyz"
		insertStringSlice(&src, data, 4)
		fmt.Println(src)
	case ArrayRemove:
		src := "abcdedf"
		removeStringSlice(&src, 2, 4)
		fmt.Println(src)
	case ArrayFilter:
		src := []int{18, 20, 15, 22, 16}
		result := filterArray(src, func(i int) bool {
			return i%2 == 0
		})
		fmt.Printf("filterArray result:\n %v\n", result)
	default:
		fmt.Println("an unknown array to test")
	}
}

func appendBytes(slice, data []byte) (newSlice []byte) {
	buffer := bytes.NewBuffer(slice)
	n, err := buffer.Write(data)
	if err != nil {
		newSlice = slice
		fmt.Printf("appendBytes error:\n %v\n", err)
		return
	}

	//test split
	tempBuffer := bytes.NewBufferString("abcdefg")
	sep := 2
	sl1, sl2 := tempBuffer.Next(sep), tempBuffer.Next(tempBuffer.Len())
	fmt.Printf("new slice1: %s\n", sl1)
	fmt.Printf("new slice2: %s\n", sl2)

	fmt.Printf("new string write: %v bytes\n", n)
	return buffer.Bytes()
}

func insertStringSlice(s *string, data string, index int) {
	ss := []byte(*s)
	dd := []byte(data)
	ss = append(ss[:index], append(dd, ss[index:]...)...)
	*s = string(ss)
}

func removeStringSlice(s *string, begin, end int) {
	ss := []byte(*s)
	ss = append(ss[:begin], ss[end:]...)
	*s = string(ss)
}

func filterArray(sl []int, fn func(int) bool) (result []int) {
	for _, data := range sl {
		if fn(data) {
			result = append(result, data)
		}
	}

	return
}
