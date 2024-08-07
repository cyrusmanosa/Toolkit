package main

import (
	"log"

	toolkit "github.com/cyrusmanosa/Toolkit"
)

func main() {
	toSlug := "NOW!!? is the time 123"
	var tools toolkit.Tools

	slugifed, err := tools.Slugify(toSlug)
	if err != nil {
		log.Println(err)
	}

	log.Println(slugifed)
}
