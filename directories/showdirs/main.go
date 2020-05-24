package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("No .env file found")
	} else {
		log.Println("Loaded .env file.")
	}
}

func readFileIntoByteArray(path string) ([]byte, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
		return []byte{}, err
	}
	return dat, nil
}

func printBytesAsString(filecontent []byte) {
	s := string(filecontent)
	fmt.Print(s)
}

func main() {
	var dirRCfilePath string

	rcFileDir := flag.String("rcdir", "", "Using DIR_RC path provided in runtime flag instead...")
	flag.Parse()

	if *rcFileDir != "" {
		dirRCfilePath = *rcFileDir
	} else {

		dirRClocation, exists := os.LookupEnv("DIR_RC_PATH")
		if exists {
			dirRCfilePath = dirRClocation
		} else {
			log.Fatal("dir_rc file location not specified. Please set DIR_RC_PATH envvar.")
		}
	}

	fmt.Println("DIR_RC_PATH: ", dirRCfilePath)

	dirRCfileContent, err := readFileIntoByteArray(dirRCfilePath)

	if err != nil {
		log.Fatal("Could not read dir_rc file.")
	}

	printBytesAsString(dirRCfileContent)
}
