package main

import (
	"os"
	"fmt"
)

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if fi.Mode() & os.ModeNamedPipe == 0 {
		fmt.Println("no pipe :(")
	} else {
		fmt.Println("hi pipe!")
	}
}
