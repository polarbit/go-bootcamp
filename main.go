package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"unicode/utf8"
)

var argFile string

func main() {
	fmt.Println("Hello Me! Let's continue learning !")

	flag.StringVar(&argFile, "file", "", "If given, filename part is printed.")
	flag.Parse()

	printArgs()

	printFilename(argFile)

	printUtf8Sample()
}

func printArgs() {
	fmt.Printf("\n=== Args ===\n")
	args := os.Args
	fmt.Printf("Args: V:%#v \n", args)
}

func printFilename(filename string) {
	fmt.Printf("\n=== Split Filename ===\n")

	if filename == "" {
		filename = "/home/dev/img/avatar.jpg"
		fmt.Println("Filename is set to: \"/home/img/avatar.jpg")
		fmt.Println("Different filename can be provided with --file param.")
	}

	_, path := path.Split(filename)
	fmt.Println("Filename:", filename, "full-path:", path)
}

func printUtf8Sample() {
	// Go source code is always utf-8.
	// A string in go is read-only slice of arbitrary bytes; not have to unicode text.
	// A string literal always holds valid UTF-8 sequences.
	// len(s) returns number of bytes; not length of text.
	// using 'for i,r := range s ...' iterates over runes, not bytes.

	fmt.Printf("\n=== UTF-8 ===\n")

	const nihongo = "日本語"
	fmt.Println("Sample string:", nihongo)

	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}

	// Above for loop can also be write simply as
	/*
		for i, r := range nihongo {
			fmt.Printf("%#U starts at byte position %d\n", r, i)
		}
	*/

	fmt.Printf("日本語 => len: %d runes: %d\n", len(nihongo), utf8.RuneCountInString(nihongo))
}
