package os_open

import (
	"fmt"
	"io"
	"os"
)

// go test -v ./os_open -args /home/ruahman/source/repos/golang/go_tut/os_open/open.go
func Run() {
	fmt.Println("--- os.Open ---")
	fmt.Println(os.Args)
	f, err := os.Open(os.Args[4])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, _ = io.Copy(os.Stdout, f)
}
