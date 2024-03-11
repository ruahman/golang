package logs

import (
	"log"
	"os"
)

func Run() {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.Print("Hello, World!")
}
