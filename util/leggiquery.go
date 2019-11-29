package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// LeggiQuery ....
func LeggiQuery(f string) []string {
	var s []string

	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		if len(t) > 5 && !strings.HasPrefix(t, "-- #") {
			s = append(s, t)
			fmt.Println(t)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return s
}
