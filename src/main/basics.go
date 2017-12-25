package main

import "unsafe"
import "fmt"
import "container/list"

//Person ...
type Person struct {
	firstName string
	lastName  string
	list.List
}

func (p *Person) iter() {
	// ...
}

func runTestBasics() {
	defer un(trace("runTestBasics"))

	pp := new(Person)
	pp.firstName = "clerk"
	pp.lastName = "..."
	pp.PushBack("boss")
	fmt.Printf("person is %v\n", pp.Front().Value)

	pers3 := &Person{"Chris", "Woodward", list.List{}}
	fmt.Printf("person name is %v\n", pers3)

	var i int = 1
	p := unsafe.Sizeof(i)
	fmt.Printf("size of int is %v\n", p)

	l := list.New()
	//l.Init()
	l.PushBack(101)
	l.PushBack(102)
	l.PushBack(103)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("list e is %v\n", e.Value)
	}
}
