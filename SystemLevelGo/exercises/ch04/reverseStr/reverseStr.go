package main

import (
	"fmt"
)

func byteStringReverse(str string) string{
	b := []byte(str);
	reverseByte := make([]byte, len(b));
	for i := range b {
		reverseByte[len(str) - 1 - i] = b[i]
	}
	return string(reverseByte)
}

func runeStringReverse(str string) string{
	r := []rune(str);
	reverseRune := make([]rune, len(r));
	for i := range r {
		reverseRune[len(r) - 1 - i] = r[i]
	}
	return string(reverseRune)
}

func main(){
	byteStr := byteStringReverse("Hello Babyy, I love youuu 💗")
	runeStr := runeStringReverse("Hello Babyy, I love youuu 💗")

	fmt.Printf("%v\n%v", byteStr, runeStr)
}
