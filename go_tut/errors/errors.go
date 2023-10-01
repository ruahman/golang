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
}
