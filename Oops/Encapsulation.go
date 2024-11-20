package oops

import "fmt"

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{name: name, age: age}
}
func (p *Person) GetName() string {
	return p.name
}
func (p *Person) SetName(name string) {
	p.name = name
}
func RunEncapsulation() {
	persom := NewPerson("Nitin", 25)

	fmt.Println(persom.GetName())
	persom.SetName("Test")
	fmt.Println(persom.GetName())
}
