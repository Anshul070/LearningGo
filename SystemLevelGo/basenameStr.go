package main

import (
	"fmt"
	"os"
	"strings"
)

func basename(path string, extention ...string) (string) {
	if path == "/" {
		return path
	} 

	path = strings.TrimRight(path, "/")
	if path == "" {
		return "/"
	}

	lastIndex := strings.LastIndex(path, "/")
	if lastIndex >= 0 {
		path =path[lastIndex + 1: ]
	}

	for _, ext := range extention {
		path = strings.TrimSuffix(path, ext)
	}

	return path
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Please pass the argument.")
		os.Exit(1)
	}

	path := os.Args[1]
	ext := ""
	if len(os.Args) > 2 {
		ext = os.Args[2]
	}

	base := basename(path, ext)
	
	fmt.Println(base)
}
