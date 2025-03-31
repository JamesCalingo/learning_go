package exercises

import (
	"fmt"
	"errors"
	"strconv"
)

func add(i int, j int) (int, error){ return i + j, nil}
func sub(i int, j int) (int, error) { return i - j, nil}
func mul(i int, j int) (int, error) { return i * j, nil}
func div(i int, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("divide by 0")
	}
	return i / j, nil}

var opMap = map[string]func(int, int) (int, error) {
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}
//I'll admit that it took me a little bit to understand what THIS part is for, but it's a way to map the operands to their functions since we're going to be inputting strings


func calculator(expression []string) (int, error) {
	
	if len(expression) != 3 {
		return 0, errors.New("invalid expression")
	}
	i1, err := strconv.Atoi(expression[0])
	if err != nil {
		return 0, err
	}
	op := expression[1]
	opFunc, ok := opMap[op]
	if !ok {
		return 0, errors.New("unsupported operator: " + op)
	}
	i2, err := strconv.Atoi(expression[2])
	if err != nil {
		return 0, err
	}
	result, err := opFunc(i1, i2)
	if err != nil {
		return 0, err
	}
	return result, nil
}

var inputs = []string{"2", "/", "0"}

func prefixer(prefix string) func(string) string {
	return func(suffix string) string {
		return fmt.Sprintf("%s %s", prefix, suffix)
		//Using Sprintf to make sure the prefix and suffix are spaced properly
	}
}

func Ch5() {
	fmt.Println(calculator(inputs))
helloPrefix := prefixer("Hello")
fmt.Println(helloPrefix("Bob"))
fmt.Println(helloPrefix("Maria"))
/*I think we should break down what on earth is going on here, cause it can get CONFUSING.

First, we use the prefixer function to create helloPrefix - a variable of type func(suffix string) string. If we didn't use a closure, declaring helloPrefix would look like this:

func helloPrefix (suffix string) string {
return "Hello " + suffix
}

Now, using that helloPrefix function, we create two strings: "Hello Bob" and "Hello Maria".

So what on earth was the point? Why couldn't we have just declared helloPrefix without the whole closure situation?

Well...you've got a point there. However, say you wanted to change the prefix to something else like "YO WASSUP" - this makes it easier (I think?) to change the prefix (i.e. once someone sees how unprofessional you've been and corrects your "Gen-something-or-other slang")
*/
}