package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args) == 2 {
		// gets the filename from the argument passed in the command line
		fn := os.Args[1]
		// opens the file name in current working directory
		f, err := os.Open(fn)
		if err != nil {
			fmt.Println("Error!", err)
		}
		// reads the file and writes to stdout
		io.Copy(os.Stdout, f)
	} else {
		fmt.Println("Error: expected 2 arguments but got", len(os.Args))
		os.Exit(1)
	}
}
