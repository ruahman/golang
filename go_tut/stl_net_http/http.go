package stl_net_http

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

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--- handlerFunc ---")
	// w is the response writer which is used to send data back to the client
	// r is the request which contains all the information about this request
	fmt.Println(w, r)
	fmt.Fprint(w, "<h1>Hello, World!</h1>")
}

func Handlers() {
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

func Get() {
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

func Static() {
	// this is the file server
	fs := http.FileServer(http.Dir("static"))
	// In order to serve files correctly, we need to strip away a part of the url path.
	// Usually this is the name of the directory our files live in.
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// http://localhost:3000/static/hello_world.txt
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(err)
	}
}
