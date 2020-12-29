package main

import (
	"flag"
	"fmt"
	"json_comparator/comparator"
	"log"
	"os"
)

const helpDescription = `The comparator receives two json files each with an array of json objects and
compares them to see if they are similar. In order for two files to be similar
they must have the same keys/values. The order of the elements is not important.

When run, 0 is returned if both files are similar, 1 if they are different and
a negative value is returned in case of an error.

  Usage: ./comparator <file1> <file2>
`

func main() {
	help := flag.Bool("help", false, "Show the comparator explanation")
	flag.Parse()

	if *help {
		fmt.Println(helpDescription)
		os.Exit(-2)
	}

	if len(os.Args) != 3 {
		fmt.Println("Usage: ./comparator <file1> <file2>")
		os.Exit(-1)
	}

	file1 := os.Args[1]
	file2 := os.Args[2]

	areSimilar, err := comparator.AreSimilar(file1, file2)
	if err != nil {
		log.Println(err)
		os.Exit(-3)
	}

	if areSimilar {
	  fmt.Println("Files are similar")
		os.Exit(0)
	}

  fmt.Println("Files are different")
	os.Exit(1)
}
