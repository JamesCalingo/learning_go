package main

import "fmt"



type Person struct {
	FirstName string
	LastName string
	Age int
}

func main() {
	var population[]Person
	dummy := Person{"Dummy", "Thicc", 2}
	for i:= 0; i < 10000000; i++ {
		population = append(population, dummy)
	}
	fmt.Println(len(population))
	fmt.Println("DONE")
}
