package enums

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func Enums() {
	fmt.Println("enums...")
	fmt.Println(StateIdle)
	fmt.Println(StateConnected)
	fmt.Println(StateError)
}
