// package main

// import (
// 	"fmt"
// 	"os"
// 	"regexp"
// )

// func basename(command string, extention string) (string, error) {
// 	if command == "/" {
// 		return "/", nil
// 	}


// 	pattern := `([^/]+)/?$`
// 	re := regexp.MustCompile(pattern)
// 	dir := re.FindAllStringSubmatch(command, -1)
	
// 	if len(dir) <= 0 {
// 		return "/", nil
// 	}

// 	res := dir[len(dir)-1][0]

// 	if extention != "" {
// 		re = regexp.MustCompile(extention)
// 		res = re.ReplaceAllString(dir[len(dir)-1][0], "")
// 	}
// 	if len(dir) > 0 {
// 		return res, nil
// 	}
	
// 	return "", fmt.Errorf("Nothing at base.")
// }

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Fprintf(os.Stderr, "Please pass the argument.")
// 		os.Exit(1)
// 	}

// 	command := os.Args[1]
// 	ext := ""
// 	if len(os.Args) > 2 {
// 		ext = os.Args[2]
// 	}

// 	base, err := basename(command, ext)
// 	if err != nil {
// 		fmt.Printf("Error: ", err)
// 		return
// 	}

// 	fmt.Println(base)
// }
