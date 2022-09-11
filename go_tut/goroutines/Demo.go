package goroutines

import "fmt"

func Demo() {
	fmt.Println("----- goroutines -----")

	urls := []string{
		"https://www.udemy.com/",
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.amazon.com/",
		"https://www.twitter.com/",
	}

	for _, url := range urls {
		go func(x string) {
			fmt.Println(x)
		}(url)
	}
}
