package exercises

import "fmt"

func greetings() {
	greetings := []string{"Hello,", "Hola,", "Arabic I believe,", "Konnichiwa,", "I think this is Cyrillic,"}
	sub1 := greetings[:2]
	sub2 := greetings[1:4]
	sub3 := greetings[3:]
	fmt.Println(greetings, sub1, sub2, sub3)
	//NOTE: The book uses various greetings in different , but that ends up being ultimately unimportant to this exercise. Also, I added commas to the greetings as it makes them a lot more readable when you print them.
}

func funWithEmojis() {
	message := "Hi 👧 and 👨"
	fourth := message[3]
	fmt.Println(string(fourth))
	//This one is meant to show that the emojis are "secretly" strings
}

func employeeList() {
	type Employee struct {
		firstName string
		lastName string
		id int
	}

	bob := Employee{"Bob", "Bobberson", 1}
	zucker := Employee{firstName: "Zucker", lastName: "Berg", id: 2}
	var stitches Employee
	stitches.firstName = "Stitches"
	stitches.lastName = "Patchery"
	stitches.id = 0
	fmt.Println(stitches, bob, zucker)
}

func Ch3() {
	greetings()
	funWithEmojis()
	employeeList()
}