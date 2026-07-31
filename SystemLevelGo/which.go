// package main

// import (
// 	"fmt"
// 	"os"
// 	"path/filepath"
// 	"runtime"
// 	"strings"
// )

// func which(command string) (string, error) {
// 	envPaths := os.Getenv("PATH")
// 	paths := strings.Split(envPaths, string(os.PathListSeparator))

// 	var extentions []string
// 	if runtime.GOOS == "windows" {
// 		ext := os.Getenv("PATHEXT")
// 		if ext == "" {
// 			extentions = []string{".cmd", ".exe", ".bat", ".com"}
// 		} else {
// 			extentions = strings.Split(strings.ToLower(ext), string(os.PathListSeparator))
// 		}
// 	} else {
// 		extentions = []string{}
// 	}

// 	for _, path := range paths {
// 		for _, ext := range extentions {
// 			var target string;
// 			if !strings.HasSuffix(command, ext) {
// 				target = command + ext
// 			}

// 			filePath := filepath.Join(path, string(target))
// 			fileInfo, err := os.Stat(filePath)
// 			if err != nil {
// 				if os.IsNotExist(err) {
// 					continue
// 				} else {
// 					fmt.Fprintf(os.Stderr, "Some error occured %v \n", err)
// 					os.Exit(1)
// 				}
// 			}

// 			if !fileInfo.Mode().IsRegular() {
// 				continue;
// 			}

// 			if runtime.GOOS != "windows" && fileInfo.Mode().Perm()&0111 == 0 {
// 				fmt.Fprintf(os.Stderr, "Do not have permission to execute file \n")
// 				os.Exit(1)
// 			}
			
// 			return filePath, nil
// 		}
// 	}

// 	return "", fmt.Errorf("File not found")
// }

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Fprintf(os.Stderr, "Please give the command also:\ngo run which.go <command> \n")
// 		os.Exit(1)
// 	}

// 	command := os.Args[1]
// 	path, err := which(command)
// 	if err != nil {
// 		fmt.Printf("Error: %v", err)
// 		return
// 	}

// 	fmt.Printf("Found: %v", path)
// }
