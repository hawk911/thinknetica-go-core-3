package main

import "fmt"

type Employee struct {
	years int
}

type Customer struct {
	years int
}

func oldest(persons []interface{}) interface{} {
	var cAge, age = 0, 0
	var p interface{}
	for _, per := range persons {
		switch t := per.(type) {
		case Customer:
			cAge = t.years
		case Employee:
			cAge = t.years
		default:
			cAge = -1
		}
		if cAge > age {
			age = cAge
			p = per
		}
	}
	return p
}

func main() {
	persons := []interface{}{
		Customer{years: 10},
		Employee{years: 40},
		Employee{years: 20},
		Employee{years: 25}}
	per := oldest(persons)
	// Per will be used for something
	fmt.Println("object is ", per)
}
