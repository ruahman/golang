package http

import (
	"fmt"
	"net/http"
	"os"
)

func Run() {
	fmt.Println("--- http ---")
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
