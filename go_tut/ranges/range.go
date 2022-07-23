package ranges

import (
	"fmt"
)

func Demo() {
	fmt.Println("----- range ----")
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for i, num := range nums {
		fmt.Println(i, num)
	}

	kvs := map[string]string{"a": "apple", "b": "bear"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println(k)
	}
}
