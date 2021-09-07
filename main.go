package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"sync"
	"time"
	"unicode/utf8"
)

var argFile string

func main() {
	fmt.Println("Hello Me! Let's continue learning !")

	flag.StringVar(&argFile, "file", "", "If given, filename part is printed.")
	flag.Parse()

	args()

	splitFilename(argFile)

	utf8Sample()

	arrays()

	sliceGotcha()

	clojure().Wait()
}

func args() {
	fmt.Printf("\n=== Args ===\n")
	args := os.Args
	fmt.Printf("Args: V:%#v \n", args)
}

func splitFilename(filename string) {
	fmt.Printf("\n=== Split Filename ===\n")

	if filename == "" {
		filename = "/home/dev/img/avatar.jpg"
		fmt.Println("Filename is set to: \"/home/img/avatar.jpg")
		fmt.Println("Different filename can be provided with --file param.")
	}

	_, path := path.Split(filename)
	fmt.Println("Filename:", filename, "full-path:", path)
}

func utf8Sample() {
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

func arrays() {
	fmt.Printf("\n=== Arrays ===\n")

	// Arrays are value types, and their types also include length.
	// So their size can not be changed in runtime. Length belongs compile time.
	seasons := [4]string{"summer", "fall", "winter", "sprint"}

	// Since arrays are values, in assignments they are copied.
	// Also in for...range operations, they are copied again;
	copy := seasons
	copy[1] = "autumn"

	fmt.Printf("Org: %#v\n", seasons)
	fmt.Printf("Copy: %#v\n", copy)
}

func sliceGotcha() {
	fmt.Printf("\n=== Slice Gotcha ===\n")

	const s = "You need to find me somehow!"

	b, err := ioutil.ReadFile("./main.go")
	if err != nil {
		panic("Could not read main.go")
	}

	x := regexp.MustCompile(`(?m:^.*somehow.*$)`)
	b2 := x.Find(b)

	/*
  When you create a slice, they both point to the same array.
  If original slice is much bigger than child slice, and
  if you don't need original slice but smaller child slice, 
  you better copy smaller slice to new slice then return it.
  */

	// Gotcha here is 'b2' slice' is created from 'b' slice.
	// Underneath, they point point to the same array, which hold all file content.
	// So for a single line, we keep whole document in the memory.
	fmt.Printf("Matched line: %q\n", string(b2))
	fmt.Printf("capacity: %v\n", cap(b))

	// Solution is, before returing the matched line, we need to copy it.
	b3 := append([]byte(nil), b2...)
	fmt.Printf("Copied line: %q (Actually hold all main.go lines in the memory)\n", string(b3))
	fmt.Printf("capacity: %v (Only a single line remains in memory)\n", cap(b3))
}

func clojure() *sync.WaitGroup {
	fmt.Printf("\n=== Clojures in Go ===\n")
  fmt.Println()
  
  var wgx sync.WaitGroup
	wgx.Add(1)

	go func() {
		local := 1
		var wg sync.WaitGroup

		// Anonymous function with local closure
		f := func() {
			fmt.Println(local)
		}

		// Print local in each second
		wg.Add(1)
		go func() {
			defer wg.Done()
      done := make(chan bool, 1)
			ticker := time.NewTicker(1 * time.Second)
      for {
        select {
          case <- ticker.C:
            f()
            if local > 5 {
              ticker.Stop()
              done <- true
            }
          case <- done:
            return
        }
      }
		}()

		// Increment local in each second
		wg.Add(1)
		go func() {
			defer wg.Done()
      done := make(chan bool, 1)
			ticker := time.NewTicker(1 * time.Second)
      // When ticker.Stop is called; ticker is stopped but 
      // it does not close the ticker channel. so "range" not works
      // This is why we use just "for"
      for {
        select {
          case <- ticker.C:
            local += 1
            if local > 5 {
              ticker.Stop()
              done <- true
            }
          case <- done:
            return
        }
      }
		}()

		wg.Wait()
		wgx.Done()
	}()

	return &wgx
}
