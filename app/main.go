package main

import (
	"fmt"

	toolkit "github.com/cyrusmanosa/Toolkit"
)

func main() {
	var tools toolkit.Tools
	s := tools.RandomString(10)

	fmt.Println(s)
}
