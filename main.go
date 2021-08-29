package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"unicode/utf8"
)

var argFile string
var argUser string

func main() {
	fmt.Println("Hello Me! Let's continue learning...")

	flag.StringVar(&argFile, "file", "", "If given, filename part is printed.")
	flag.StringVar(&argUser, "user", "", "If given, string length is printed.")
	flag.Parse()

	args := os.Args
	fmt.Printf("Args: %#v \n", args)

	if argFile != "" {
		printFilename(argFile)
	}

	if argUser != "" {
		printStrLen(argUser)
	}
}

func printFilename(filename string) {
	_, path := path.Split(filename)
	fmt.Println("Filename:", path, "full-path:", argFile)
}

func printStrLen(s string) {
	fmt.Printf("%v bytes=%v length=%v\n", s, len(s), utf8.RuneCountInString(s))
}
