package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func main() {
	p := "copyfrom.sql"
	p, _ = filepath.Abs(p)
	b, err := ioutil.ReadFile(p) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	tmp := string(b)
	examples := strings.Split(tmp, ";")
	fmt.Printf("%d esempi", len(examples))
}
