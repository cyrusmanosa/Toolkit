package main

import (
	toolkit "github.com/cyrusmanosa/Toolkit"
)

func main() {
	var tools toolkit.Tools

	tools.CreateDirIfNotExist("./test-dir")
}
