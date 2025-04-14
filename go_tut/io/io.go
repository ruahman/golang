package io

import (
	"fmt"
	"log"
	"os"
)

func Io() {
	content, err := os.ReadFile("./foo.txt")
	if err != nil {
		fmt.Println("error no file")
		log.Fatal(err)
		panic("aaaaa")
	} else {
		str := string(content)
		fmt.Println(str)
	}
	dir, _ := os.Getwd()
	fmt.Println(dir)

	// write file
	err = os.WriteFile("foo.txt", []byte("bar from go"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
