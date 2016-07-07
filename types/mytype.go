package types

import "fmt"

type OpenType struct {
	Name string
	Age  int
}

type CloseType struct {
	name string
	age  int
}

func (t OpenType) Display() {
	fmt.Println(t.Name, t.Age)
}

func (t OpenType) display() {
	fmt.Println(t.Name, t.Age)
}

func (t CloseType) Display() {
	fmt.Println(t.name, t.age)
}

func (t CloseType) display() {
	fmt.Println(t.name, t.age)
}

func NewCloseType(name string, age int) CloseType {
	a := CloseType{name, age}
	fmt.Printf("%#p\n", &a)
	return a
}
