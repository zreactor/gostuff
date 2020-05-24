package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
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

func readToLastNewLine(dirRCfileContent []byte) []byte {
	filelen := len(dirRCfileContent) - 1
	for i := filelen; i >= 0; i-- {
		if dirRCfileContent[i] == 10 {
			return dirRCfileContent[i+1:]
		}
	}
	return dirRCfileContent
}

func readFromEnd(dirRCfileContent []byte) ([]byte, bool) {
	var _f []byte
	startWithNL := false
	if len(dirRCfileContent) == 0 {
		fmt.Println("File was empty.")
		return []byte{}, startWithNL
	}

	if dirRCfileContent[len(dirRCfileContent)-1] == 10 {
		fmt.Println("ends with enter")
		_f = dirRCfileContent[:(len(dirRCfileContent) - 1)]
	} else {
		fmt.Println("ends with not-enter")
		startWithNL = true
		_f = dirRCfileContent
	}
	lastline := readToLastNewLine(_f)
	return lastline, startWithNL
}

func calculateNextAliasName(dirRCfileContent []byte) (string, bool) {
	var nr string
	lastline, startWithNL := readFromEnd(dirRCfileContent)

	if len(lastline) == 0 {
		nr = "01"
	} else {
		nr = getLatestNumber(lastline)
	}
	return ("DIR" + nr), startWithNL
}

func getLatestNumber(inputtext []byte) string {
	res1 := strings.SplitN(string(inputtext), "DIR", -1)
	res2 := strings.SplitN(res1[1], "=", -1)
	number := res2[0]
	i, err := strconv.Atoi(number)
	if err != nil {
		log.Println(err)
	}
	nextIndex := i + 1
	if nextIndex < 10 {
		return "0" + strconv.Itoa(nextIndex)
	}
	return strconv.Itoa(nextIndex)

}

func composeEntryString(filename string, path string, startWithNL bool) string {
	entryString := "export " + filename + "=\"" + path + "\"\n"
	if startWithNL == true {
		return "\n" + entryString
	}
	return entryString
}

func appendLineToFile(entry string, filename string) {
	fmt.Println("Will append:", entry, filename)
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	if _, err := file.WriteString(entry); err != nil {
		log.Fatal(err)
	}
}

func getPWD() string {
	pwdPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return pwdPath
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

	nextDirAlias, startWithNL := calculateNextAliasName(dirRCfileContent)
	nextRecord := composeEntryString(nextDirAlias, getPWD(), startWithNL)

	appendLineToFile(nextRecord, dirRCfilePath)
	fmt.Println("Done! Make sure to source your profile file before using. Ex. source .bash_profile")
}
