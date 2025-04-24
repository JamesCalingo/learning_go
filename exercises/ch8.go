package exercises

import "fmt"

// EXERCISE 1: Double
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func double[T Integer](val T) T {
	return val * 2
}

// EXERCISE 2: Printable
type Printable[T any] interface{ 
	fmt.Stringer
	~int | ~float64
}


func printLine[T Printable[T]](val T) {
	fmt.Println(val)
}

//EXERCISE 3: Linked List

type List[T any] struct {
	head *Node[T]
	tail *Node[T]
}

type Node[T any] struct {
	Val T
	Next *Node[T]
}

func (l *List[T]) Add(v T) {
	n := &Node[T]{Val: v}
	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}
	l.tail.Next = n
	l.tail = l.tail.Next
}

func (n *Node[T]) Insert(val T, index int) {

}

func(n Node[T]) Index()int {
	return -1
}

func Ch8() {
	fmt.Println(double(3))
	fmt.Println(double(2.1))
	printLine(3)
}
