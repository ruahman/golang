package errors

import (
	"errors"
	"fmt"
)

// an error is anything that implements the Error() method

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

// a sentinel error is a predefined error that we can compare against
var (
	ErrOutOfTea   = errors.New("out of tea")
	ErrOutOfPower = errors.New("out of power")
)

func makeTea(arg int) error {
	if arg == 0 {
		return ErrOutOfTea
	} else if arg == 4 {
		// we can wrap an error with other errors to provide context with fmt.Errorf
		return fmt.Errorf("making tea: %w", ErrOutOfPower)
	}
	return nil
}

// you define custom error types by implementing the Error() method
type MyError struct {
	arg     int
	message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &MyError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func Exec() {
	fmt.Println("----- errors -----")

	if r, e := f1(11); e != nil {
		fmt.Println("f1 failed", e)
		fmt.Println(e.Error())
	} else {
		fmt.Println(r)
	}

	if r, e := f1(42); e != nil {
		fmt.Println("f1 failed", e)
		fmt.Println(e.Error())
	} else {
		fmt.Println(r)
	}

	for i := range []int{0, 1, 2, 3, 4} {
		if e := makeTea(i); e != nil {
			fmt.Println("makeTea failed", e)
		} else {
			fmt.Println("tea is ready")
		}
	}

	// return custom error type
	if r, e := f2(42); e != nil {
		fmt.Println("f2 failed", e)
		fmt.Println(e.Error())
	} else {
		fmt.Println(r)
	}
}
