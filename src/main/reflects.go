package main

import (
	"fmt"
	"reflect"
)

// NotknownType ...
type NotknownType struct {
	S1, S2, S3 string
}

func (n NotknownType) String() string {
	return n.S1 + " - " + n.S2 + " - " + n.S3
}

func (n NotknownType) altString() string {
	return n.S1 + " -- " + n.S2 + " -- " + n.S3
}

// variable to investigate:
var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

func runTestReflects() {
	defer un(trace("runTestReflects"))

	value := reflect.ValueOf(secret)
	typeV := value.Type()
	fmt.Println("typeV:", typeV)
	knd := value.Kind() // struct
	fmt.Println(knd)
	for i := 0; i < value.NumField(); i++ {
		v := value.Field(i)
		fmt.Println(v.Interface())
		fmt.Printf("Field [%v] : %v\n", i, value.Field(i))
	}

	x := 3.14
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float())
	fmt.Println("canset:", v.CanSet())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)
}
