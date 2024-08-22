package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := range 100 {
		output := ""
		if (i+1)%3 == 0 {
			output += "Fizz"
		}
		if (i+1)%5 == 0 {
			output += "Buzz"
		}
		if output == "" {
			output = strconv.Itoa(i + 1)
		}
		fmt.Println(output)
	}
}
