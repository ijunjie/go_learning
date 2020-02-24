package ch5

import (
	"fmt"
	"testing"
)

type person struct {
	age int
	name string
	car Car
}

type Car struct {
	name string
}

func (p *person) setAge(age int) {
	p.age = age
}

func (p person) withAge(age int) person {
	p.age = age
	return p
}


func TestObject(t *testing.T) {
	person := &person{2,"aaa", Car{"mycar"}}
	person.setAge(3)
	fmt.Printf("person: %+v\n", *person)

	newPerson := (*person).withAge(8)
	fmt.Printf("person: %+v\n", *person)
	fmt.Printf("person: %+v\n", newPerson)
}
