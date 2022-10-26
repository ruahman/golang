package functions

import (
	"fmt"
)

func newCard() string {
	return "Five of Diamonds"
}

// func sum(nums ...int) {
// 	fmt.Println(nums, " ")
// 	total := 0
// 	for _, num := range nums {
// 		total += num
// 	}
// 	fmt.Println(total)
// }

func multiple_return() (int, bool) {
	return 123, false
}

func lastHi() {
	fmt.Println("lastHi")
}

// multiple consecutive params
func plusPlus(a, b, c int) int {
	return a + b + c
}

// unlimited params
func sum(nums ...int) {
	fmt.Println(nums)
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func closure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func learnDefer() (ok bool) {
	// A defer statement pushes a function call onto a list. The list of saved
	// calls is executed AFTER the surrounding function returns.
	defer fmt.Println("deferred statements execute in reverse (LIFO) order.")
	defer fmt.Println("\nThis line is being printed first because")
	// Defer is commonly used to close a file, so the function closing the
	// file stays close to the function opening the file.
	return true
}

// Functions can have variadic parameters.
func learnVariadicParams(myStrings ...interface{}) {
	// Iterate each value of the variadic.
	for idx, param := range myStrings {
		fmt.Println("index: ", idx, "param:", param)
	}

	// Pass variadic value as a variadic parameter.
	fmt.Println("params:", fmt.Sprintln(myStrings...))

}

func learnFunctionFactory() {
	fmt.Println(sentenceFactory("summer")("A beautiful", "day!"))

	d := sentenceFactory("summer")
	fmt.Println(d("A beautiful", "day!"))
	fmt.Println(d("A lazy", "afternoon!"))
}

// Decorators are common in other languages. Same can be done in Go
// with function literals that accept arguments.
func sentenceFactory(mystring string) func(before, after string) string {
	return func(before, after string) string {
		return fmt.Sprintf("%s %s %s", before, mystring, after) // new string
	}
}

/*
Functions can have parameters and (multiple!) return values.
Here `x`, `y` are the arguments and `sum`, `prod` is the signature (what's returned).
Note that `x` and `sum` receive the type `int`.
*/
func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y // Return two values.
}

// recursion
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func closureRecursion() {
	fmt.Println(fact(5))

	// if you want to make a clojure recursive
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}

// variadic function can be called with any number of arguments.
func sum2(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// func PanicRecover() {
// 	panic("I'm panicing")

// 	defer func() {
// 		if error := recover(); error != nil {
// 			fmt.Println("caught the panic")
// 		}
// 	}()
// }

// // closures
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func Demo() {
	fmt.Println("----- functions -----")

	card := newCard()
	i, b := multiple_return()

	// call this at the end with top fuction is about to close
	defer lastHi()

	// push on defer stack
	defer func() {
		fmt.Println("Almost last??")
	}()

	fmt.Println("multiple retrun: ", i, b)
	fmt.Println(card)
	fmt.Println("demo from functions")

	nums := []int{1, 2, 3}
	sum(nums...)

	c := closure()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())

	learnDefer()

	learnVariadicParams("great", "learning", "here!")

	fmt.Println(sentenceFactory("summer")("A beautiful", "day!"))

	d := sentenceFactory("summer")
	fmt.Println(d("A beautiful", "day!"))
	fmt.Println(d("A lazy", "afternoon!"))

	fmt.Println(learnMultiple(3, 4))

	fmt.Println(plusPlus(1, 2, 3))
	fact(7)
	closureRecursion()

	myFunc := func() {
		fmt.Println("anomous")
	}

	myFunc()

	fmt.Println("variadic sum", sum2(1, 2, 3, 4, 5))

	// PanicRecover()

	//// clojures
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

}
