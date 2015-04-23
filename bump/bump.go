package main

import (
	"flag"
	"log"

	"github.com/olivoil/bump"
)

func main() {
	flag.Parse()

	command := flag.Arg(0)
	files := flag.Args()[1:]

	for _, file := range files {
		err := bump.File(command, file)
		if err != nil {
			log.Fatalf("Error updating %s file: %s", file, err.Error())
		}
	}
}
