package io

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Demo() {
	fmt.Println("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		fmt.Println("hello ", name)
	} else {
		log.Fatal(err)
	}
}
