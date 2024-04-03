package pointers

import (
	"errors"
	"fmt"
)

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func Run() {
	fmt.Println("----- pointer demo -----")

	var x int = 3
	var px *int = &x
	fmt.Println(x, px, *px)

	var py *int = new(int)
	*py = 3
	fmt.Println(py, *py)

	i := 1
	fmt.Println(i)

	// doen't change i
	zeroval(i)
	fmt.Println(i)

	// changes i
	zeroptr(&i)
	fmt.Println(i)

	fmt.Println(&i)

	a := 5
	b := &a
	fmt.Println(a, b, *b)
	*b = 10
	fmt.Println(*b)

	// a pointer holds the memory address of a value
	var p *int
	c := 42
	p = &c
	fmt.Println(p, *p)
}
