package http

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println("--- logWriter ---")
	fmt.Println(string(bs))
	return len(bs), nil
}

func handlerFunc(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("--- handlerFunc ---")
	fmt.Fprint(w, "<h1>Hello, World!</h1>")
}

func RunHandlers() {
	// 100:  information, still loading
	// 200:  success, 200, 202
	// 300:  redirection
	// 400:  client error
	// 500:  server error
	http.HandleFunc("/", handlerFunc)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(err)
	}
}

func Run() {
	fmt.Println("--- http ---")
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// fmt.Println("--- fmt.Println ---")
	// bs := make([]byte, 99999)
	// _, _ = resp.Body.Read(bs)
	// fmt.Println(string(bs))

	fmt.Println("--- io.Copy ---")
	// _, _ = io.Copy(os.Stdout, resp.Body)
	lw := logWriter{}
	_, _ = io.Copy(lw, resp.Body)
}
