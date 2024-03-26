package generics

import (
	"fmt"
)

// with generics you can specify the datatype at a latter time

// no type check with this approach
// func print(output ...interface{}) {
// 	fmt.Println(output)
// }

// define your generic type
type MyConstraint interface {
	int | float64
}

func getSum[T MyConstraint](x T, y T) T {
	return x + y
}

type Type interface {
	int | string
}

// make sure all types are the same
func print_generic[T Type](output ...T) {
	fmt.Println(output)
}

// you can also use any if you want anything
func print_any[T any](output ...T) {
	fmt.Println(output)
}

/***
 * this take a map and returns all the keys and values as a slice
 *
 * note: the comparable contraint is used to make sure the key is comparable
 *
 * @param m map[K]V - a map with key and value
 * @return []K - a slice of keys
 */
func MapKey[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// link lists are a usful example for generics
type element[T any] struct {
	next  *element[T]
	value T
}
type List[T any] struct {
	head *element[T]
	tail *element[T]
}

// append to link list
func (l *List[T]) Push(v T) {
	e := &element[T]{value: v}
	if l.head == nil {
		l.head = e
		l.tail = l.head
	} else {
		l.tail.next = e
		l.tail = l.tail.next
	}
}

// get all the values in the list
func (l *List[T]) GetAll() []T {
	var result []T
	for e := l.head; e != nil; e = e.next {
		result = append(result, e.value)
	}
	return result
}

func Run() {
	fmt.Println("----- generics -----")
	print(1, 2, 3, 4, 5)
	print(1, "2", 3.3)
	print_generic(6, 7, 8, 9)

	// can't do this since Type only work for int and string
	// print_generic(6.6, 7.7, 8.8, 9.9)
	print_generic("6", "7", "8", "9")

	print_any(6.6, 7.7, 8.8, 9.9)

	// this causes a problem
	// print_generic(6, 7.7, "8", 9)

	fmt.Println(getSum(5, 4))
	fmt.Println(getSum(5.5, 4.4))

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(MapKey(m))

	l := List[int]{}
	l.Push(1)
	l.Push(2)
	l.Push(3)
	fmt.Println(l.GetAll())
}
