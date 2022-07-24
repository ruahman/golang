package err

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

func Demo() {
	fmt.Println("----- errors -----")

	if r, e := f1(11); e != nil {
		fmt.Println("f1 failed", e)
	} else {
		fmt.Println(r)
	}
}
