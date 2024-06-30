package main

import (
	"fmt"
)

func main() {
	a := make(map[int]int)

	c := 0
	for i := 0; i <= 9; i++ {

		fmt.Scan(&c)

		if v, ok := a[c]; ok {

			fmt.Print(v, " ")

		} else {
			a[c] = work(c)
			fmt.Print(a[c], " ")
		}

	}
}
