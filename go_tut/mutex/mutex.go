package mutex

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	balance int
	lock    sync.Mutex
}

func (a *Account) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

func (a *Account) WithDraw(v int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if v > a.balance {
		fmt.Println("to big")
	} else {
		a.balance = a.balance - v
		fmt.Printf("withdaw %d\n", v)
	}
}

func Demo() {
	var acct Account
	acct.balance = 100
	fmt.Println("get balance: ", acct.GetBalance())
	for i := 0; i < 12; i++ {
		go acct.WithDraw(10)
	}
	time.Sleep(2 * time.Second)
}
