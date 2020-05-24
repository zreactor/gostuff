package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	timestamp := strconv.FormatInt(time.Now().UTC().Unix(), 10)
	fmt.Println(timestamp)
}
