package demo

import (
	"fmt"
	"goland-demo/demo/shapes"
)

// GENERAL
// TODO: Search everywhere
// TODO: Find Action
// TODO: Find
// TODO: Recent Files/Locations

// EDITING
// TODO: Navigation
// TODO: Code/Directory Folding
// TODO: Show Context Actions
// TODO: Renaming and Deleting Variables, Functions and Files
// TODO: Change signature
// TODO: Find Usages

func sum(x, y int64) int64 {
	return x + y
}

func avg(x, y int64) int64 {
	return (x + y) / 2
}

type Person struct {
	ID   int    `json:"ID,omitempty" bson:"id"`
	Name string `json:"Name,omitempty" bson:"name"`
}

func NewPerson(ID int, name string) *Person {
	return &Person{ID: ID, Name: name}
}

func (p *Person) Greet(expr string) {
	fmt.Printf("Hi! My Name is %s\n", p.Name)
}

func Demo() {
	var x int64 = -3
	var y int64 = 5

	s := sum(x, y)
	a := avg(x, y)
	fmt.Printf("The sum and average of %d and %d is %d and %d, respectively. \n", x, y, s, a)

	p := NewPerson(1, "jan01")
	p.Greet("rawr")

	rect := shapes.Rectangle{
		Width:  3,
		Height: 4,
	}
	fmt.Println("Area of rectangle is", rect.Area())
}
