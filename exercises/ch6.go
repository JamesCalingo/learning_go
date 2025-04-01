package exercises

import "fmt"

//EXERCISE 1: struct pointers
type Person struct {
	FirstName string
	LastName string
	Age int
}

func makePerson(firstName string, lastName string, age int) Person {
	var person Person
	person.FirstName = firstName
	person.LastName = lastName
	person.Age = age
	return person
}

func makePersonPointer(firstName string, lastName string, age int) *Person {
	person := Person {
		FirstName: firstName,
		LastName: lastName,
		Age: age,
	}
return &person
}
/* I had to copy this code into main to get it to work properly with the build flag, because what happens when you run it is that it tells you EXACTLY what it's doing. 

Things are starting to get a little weird and wild here. First, the compiler gives a warning on makePerson should you attempt to call these functions without assigning the return values (because there are no side effects and the return value is ignored, so it's essentially doing nothing). However, since it's just a warning and NOT an error, we can TECHNICALLY ignore it (though doing so is bad practice).

Anyway, if we build the program as written below (with variables to hold the return values because we're good programmers and a Println because creating variables and not using them WILL error), here's the log I got:

# learning_go
./main.go:10:6: can inline makePerson
./main.go:18:6: can inline makePersonPointer
./main.go:28:21: inlining call to makePerson
./main.go:29:28: inlining call to makePersonPointer
./main.go:30:13: inlining call to fmt.Println
./main.go:10:17: leaking param: firstName to result ~r0 level=0
./main.go:10:35: leaking param: lastName to result ~r0 level=0
./main.go:18:24: leaking param: firstName
./main.go:18:42: leaking param: lastName
./main.go:19:2: moved to heap: person
(This is where the non-variable log ends; the following three lines only appear if the variables are used)
./main.go:29:28: moved to heap: person
./main.go:30:13: ... argument does not escape
./main.go:30:14: steve escapes to heap

The thing that is most interesting to me in is that last line: steve escapes to heap. No, steve is not doing that to get away from the angry Smash Bros. community accusing him of ruining the game; steve escapes to heap means...I'm not entirely sure honestly. 

*/

//EXERCISE 2: update and grow
func UpdateSlice(strings []string, update string) {
	strings[len(strings) - 1] = update
	fmt.Println(strings)
}
//This function is changing the slice that is passed in by replacing the value at len-1 to the update

func GrowSlice(strings []string, update string) {
	strings = append(strings, update)
	fmt.Println(strings)
}
//This adds the new value to strings

/* What ends up happening here is rather interesting: if run in the order they're written (update then grow), then the slice we use for this gets changed by update and then THAT slice gets grown. However, as written (i.e. no return values), said growth is only local to the GrowSlice function - if we were to print strings one final time after running GrowSlice, we would only see the result of UpdateSlice!

This can also be seen if we were to reverse the order of the functions, as something VERY DIFFERENT happens: GrowSlice ends up growing the slice by 1, but instead of updating at the new len-1 (i.e. value 4), UpdateSlice is run on the original length 3 slice, and thus occurs on "Three". I'm not sure that this was intended, but I assume it is as it's meant to highlight "slices" as a pointer to a memory address and then the corresponding data after that in memory (using variables changes everything).

*/

func theApproximatePopulationOfNewJersey() {
	var population[]Person
	dummy := Person{"Dummy", "Thicc", 2}
	for i:= 0; i < 10000000; i++ {
		population = append(population, dummy)
	}
	fmt.Println(len(population))
	fmt.Println("DONE")
}
// New Jersey's population, as of a 2024 estimate, was around 9.5 million. Also, don't give functions names like this kids...

func Ch6() {
	steve := makePerson("Steve", "Steeeeeeeeve", 69)
	spike := makePersonPointer("Spike", "Lakilester", 24)
	fmt.Println(steve, spike)

	strings := []string{"One,", "Two,", "Three,"}
	fmt.Println(strings)
	UpdateSlice(strings, "Updated four.")
	fmt.Println(strings)
	GrowSlice(strings, "Grown four.")
	fmt.Println(strings)
	theApproximatePopulationOfNewJersey()
}