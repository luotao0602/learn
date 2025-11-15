package main

import (
	"fmt"
)

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}

type Circle struct {
}

func (re *Rectangle) Area() {
	fmt.Println("Rectangle.Area()")
}

func (re *Rectangle) Perimeter() {
	fmt.Println("Rectangle.Perimeter()")
}

func (c *Circle) Area() {
	fmt.Println("Circle.Area()")
}

func (c *Circle) Perimeter() {
	fmt.Println("Circle.Perimeter()")
}

func main() {
	re := &Rectangle{}
	c := &Circle{}
	re.Area()
	re.Perimeter()
	c.Area()
	c.Perimeter()
}
