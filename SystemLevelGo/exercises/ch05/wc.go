package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

var (
	countLines = flag.Bool("l", false, "print the newline counts")
	countWords = flag.Bool("w", false, "print the word counts")
	countChars = flag.Bool("c", false, "print the chars")
)

func countOld(scanner *bufio.Scanner)(int, int, int){
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

func countWordsInString(line string) (wordCount int) {
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wordCount++
	}
	return
}

func count(r io.Reader)(lineCount int, wordCount int, charCount int){
	b := bufio.NewReader(r)
	for {
		line,err := b.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("Error read: ", err)
			return
		}

		if len(line) > 0 {
			lineCount++
			trimmed := strings.TrimSuffix(line, "\r\n")
			wordCount += countWordsInString(trimmed)
			charCount += utf8.RuneCountInString(trimmed) + 1
		}

		if err == io.EOF {
			break
		}
	}
	return
}

func printCount(lineCount int, wordCount int, charCount int, filename string) {
	if !*countLines && !*countWords && !*countChars {
		fmt.Printf("%7d %7d %7d %s\n", lineCount, wordCount, charCount, filename)
		return
	}

	if *countLines {
		fmt.Printf("%7d ", lineCount)
	}
	if *countWords {
		fmt.Printf("%7d ", wordCount)
	}
	if *countChars {
		fmt.Printf("%7d ", charCount)
	}
	fmt.Printf("%s\n", filename)
}


func main(){
	flag.Parse()
	Args := flag.Args() 
	fmt.Println(Args)
	if len(Args) == 0 {
		fmt.Println("wc <filename>")
		return
	}

	for _, f := range Args {
		file, err := os.Open(f)
		if err != nil {
			fmt.Println( "Error opeinig file: ", err)
			continue
		}
		lc, wc, cc := count(file)
		file.Close()

		printCount(lc, wc, cc, f)
	}


	// // It will work with countOld
	// f, err := os.Open("../../Notes.txt")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error : %v", err)
	// 	os.Exit(1)
	// }
	// defer f.Close()
	// scanner := bufio.NewScanner(f)
	// scanner.Split(bufio.ScanLines)
	// lines, words, chars := countOld(scanner)
	// fmt.Printf("Lines: %v	Words: %v	Chars: %v",lines, words, chars)
}
