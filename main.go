package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fname := os.Args[1]
	abs_fname, _ := filepath.Abs(fname)

	seila, _ := os.ReadFile(abs_fname)

	fmt.Println(string(seila))
}
