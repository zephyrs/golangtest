package main

import (
	"container/list"
	"fmt"
	"reflect"
	"unsafe"
)

//Info ...
type Info interface {
	GetSalary() int
}

//Any ...
type Any interface{}

//Person ...
type Person struct {
	firstName string
	lastName  string
	salary    int
	colleage  list.List
}

type per Person

//Manager ...
type Manager struct {
	Person
	department string
}

func (p *Person) giveRaise() {
	p.salary += 500
}

//GetSalary ...
func (p *Person) GetSalary() int {
	return p.salary
}

func (p *Person) String() string {
	return p.firstName + " " + p.lastName
}

func (p *Manager) String() string {
	return "Sir " + p.firstName + " " + p.lastName
}

//LastName ...
func (p *Person) LastName() string {
	return p.lastName
}

//SetLastName ...
func (p *Person) SetLastName(s string) {
	p.lastName = s
}

func runTestBasics() {
	defer un(trace("runTestBasics"))

	getInfo1 := func(it Info) {
		fmt.Printf("getInfo1 print %v\n", it.GetSalary())
	}

	getInfo2 := func(any Any) {
		if v, ok := any.(Info); ok {
			fmt.Printf("getInfo2 print %v\n", v.GetSalary())
		} else {
			fmt.Println("getInfo2 not work")
		}
	}

	var magic interface{}
	magic = 500

	pp := new(Person)
	pp.firstName = "Jason"
	pp.lastName = "Bison"
	pp.salary = magic.(int)
	pp.giveRaise()
	fmt.Printf("person salary is %v\n", pp.salary)

	getInfo1(pp)
	getInfo2(pp)

	mm := &Manager{Person{"Jackie", "Chan", 5000, list.List{}}, "Film"}
	fmt.Printf("manager is %v, salary is %v\n", mm, mm.salary)

	v := reflect.ValueOf(*mm)
	val := v.FieldByIndex([]int{0, 2})
	fmt.Println("FieldByIndex val", val)

	pers3 := &Person{"Chris", "Woodward", 1000, list.List{}}
	var info Info
	info = pp
	fmt.Printf("interface is %v\n", info.GetSalary())

	if sv, ok := interface{}(pp).(Info); ok {
		fmt.Printf("pp implements Info(): %v\n", sv.GetSalary())
	}

	pers3.colleage.PushBack(pp)
	pers3.colleage.PushBack(mm)

	fun := func(p *per) {
		fmt.Printf("fun get name is %v\n", p.lastName)
	}
	fun((*per)(pers3))

	col := pers3.colleage
	for e := col.Front(); e != nil; e = e.Next() {
		var salary int
		/*
			if _, ok := e.Value.(*Person); ok {
				salary = e.Value.(*Person).salary
			} else {
				salary = e.Value.(*Manager).salary
			}
		*/
		switch t := e.Value.(type) {
		case *Person:
			salary = t.salary
		case *Manager:
			salary = t.GetSalary()
		default:
			salary = -1
		}
		fmt.Printf("colleage name is %v, %v\n", e.Value, salary)
	}

	i := 1
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
