package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		os.Exit(-1)
	}
	inputFilename := flag.Arg(0)
	log.Println(inputFilename)

	ret, err := GenerateTest(inputFilename)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(ret)
}
