package main

import "fmt"

var simpleMap map[string]int

func main() {
	simpleMap = map[string]int{
		"diego": 40,
		"andy":  37,
	}

  fmt.Println(simpleMap)

	currency := map[string]int{
		"USD": 55,
		"BTC": 77,
	}

  fmt.Println(currency)

  for k,v := range simpleMap {
    fmt.Println(k,v)
  }

}
