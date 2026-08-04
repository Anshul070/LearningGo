package main

import (
	"errors"
	"fmt"
	"os"
)

func main(){
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "program.go <filename to create>\n");
		os.Exit(1)
	}

	filename := os.Args[1]
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		if errors.Is(err, os.ErrExist){
			fmt.Fprintf(os.Stderr, "File Already Exists: %v\n", filename)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "There is some error creating: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	fmt.Println("File created")
}
