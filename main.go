package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	fname := os.Args[1]
	abs_fname, errAbs := filepath.Abs(fname)

	if errAbs != nil {
		log.Fatal(errAbs.Error())
	}

	file, err := os.ReadFile(abs_fname)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(file))
}
