package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	args := os.Args
	fmt.Println("Hello Me!")
	fmt.Println(getFilename(args[1]))
}

func getFilename(filename string) string {
	_, path := path.Split(filename)
	return path
}
