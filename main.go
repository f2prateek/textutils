package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var titleCmd = flag.Bool("title", false, "title case the input")

func main() {
	flag.Parse()

	if *titleCmd {
		title()
	}
}

func title() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(strings.Title(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}
}
