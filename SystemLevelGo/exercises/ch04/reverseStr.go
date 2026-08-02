package main

import (
	"fmt"
	"strings"
)

func byteStringReverse(str string){
	reverseByte := make([]byte, len(str));
	reader := strings.NewReader(str);
	for i := range str {
		b, err := reader.ReadByte()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		reverseByte[len(str) - 1 - i] = b
	}
	fmt.Println(reverseByte)
}

func main(){
	byteStringReverse("Hello")
}
