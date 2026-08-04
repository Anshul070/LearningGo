package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	// "io"
)

func count(scanner *bufio.Scanner)(int, int, int){
	lines, words, chars := 0,0,0;
	for reading := true; reading ;{
		
		reading = scanner.Scan()
		line := scanner.Text()
		lines++
		
		wrds := strings.Split(line, " ")
		words += len(wrds)
		
		chrs := strings.Split(line, "")
		chars += len(chrs)
		if err := scanner.Err(); err != nil {
			break
		}
	}
	return lines, words, chars
}


func main(){
	f, err := os.Open("../../Notes.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error : %v", err)
		os.Exit(1)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	lines, words, chars := count(scanner)

	fmt.Printf("Lines: %v	Words: %v	Chars: %v",lines, words, chars)
}
