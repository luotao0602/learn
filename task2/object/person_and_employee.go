package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e *Employee) PrintInfo() {
	fmt.Println("EmployeeID:", e.EmployeeID)
	fmt.Println("Name:", e.Name)
	fmt.Println("Age:", e.Age)
	fmt.Println("Person.Name:", e.Person.Name)
	fmt.Println("Person.Age:", e.Person.Age)
}

func main() {
	re := Employee{
		Person: Person{
			Name: "张三",
			Age:  18,
		},
		EmployeeID: 123,
	}

	re.PrintInfo()
}
