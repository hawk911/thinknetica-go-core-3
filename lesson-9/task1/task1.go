package main

import "fmt"

type Employee struct {
	years int
}

type Customer struct {
	years int
}

type person interface {
	age() int
}

func (e Employee) age() int {
	return e.years
}

func (c Customer) age() int {
	return c.years
}
func oldest(persons []person) int {
	years := 0
	for _, e := range persons {
		if e.age() > years {
			years = e.age()
		}
	}
	return years
}

func main() {
	persons := []person{
		&Customer{years: 23},
		&Employee{years: 45},
		&Customer{years: 19},
		&Employee{years: 55},
	}
	age := oldest(persons)
	fmt.Println("The oldest person is", age, "years old.")
}
