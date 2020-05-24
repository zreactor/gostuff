package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {

	flag.Parse()
	clargs := flag.Args()

	if flag.NArg() != 1 {
		log.Println("Takes exactly one argument")
		os.Exit(1)
	}

	ut, err := strconv.Atoi(clargs[0])

	if err != nil {
		log.Println("Could not parse unixtime")
		os.Exit(2)
	}

	jstTimestamp := time.Unix(int64(ut), 0)
	fmt.Println(jstTimestamp)
}
