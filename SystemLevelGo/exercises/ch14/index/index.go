package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type Storage struct {
	Path string `json:"path,omitempty"`
	Op   string `json:"op"`
}

type Journal struct {
	File *os.File
}

type Node struct {
	Children map[string]*Node
	Files    []string
}

type Store struct {
	Journal   *Journal
	Registary map[string]bool
	Total     int
	Root      *Node
}

func NewJournal(filepath string) *Journal {
	f, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Sprintf("Error : ", err))
	}

	return &Journal{f}
}

func NewStore(filepath string) *Store {
	Journal := NewJournal(filepath)

	return &Store{Journal: Journal, Registary: make(map[string]bool), Root: new(Node)}
}

func Tokenizer(path string) []string {
	replaced := strings.ReplaceAll(path, "\\", "/")
	return strings.FieldsFunc(replaced, func(r rune) bool {
		return r == '/' || r == '.' || r == '_' || r == '-'
	})
}

func InternalAdd(s *Store, path string) {
	if s.Registary[path] {
		return
	}
	// s.Journal.UpdateJournal("index", path)
	s.Registary[path] = true
	s.Total++

	tokens := Tokenizer(path)

	for _, token := range tokens {
		node := s.Root

		token = strings.ToLower(token)

		for _, char := range token {
			charStr := string(char)

			if node.Children == nil {
				node.Children = make(map[string]*Node)
			}
			if _, exists := node.Children[charStr]; !exists {
				node.Children[charStr] = &Node{Children: make(map[string]*Node)}
			}
			node = node.Children[charStr]
		}

		if !slices.Contains(node.Files, path) {
			node.Files = append(node.Files, path)
		}
	}
}

func (s *Store) FindItem(path string) []string {
	tokens := Tokenizer(path)
	var results []string
	seen := make(map[string]bool)
	for _, token := range tokens {
		node := s.Root

		token = strings.ToLower(token)

		for _, char := range token {
			charStr := string(char)

			if next, exists := node.Children[charStr]; exists {
				node = next
			} else {
				return results
			}
		}
		var collect func(n *Node)
		collect = func(n *Node) {
			for _, f := range n.Files {
				if !seen[f] {
					results = append(results, f)
					seen[f] = true
				}
			}

			for _, nn := range n.Children {
				collect(nn)
			}
		}
		collect(node)
	}

	return results
}

func InternalDelete(s *Store, path string) {
	if !s.Registary[path] {
		return
	}

	// s.Journal.UpdateJournal("delete", path)

	delete(s.Registary, path)
	s.Total--

	tokens := Tokenizer(path)

	for _, token := range tokens {
		token = strings.ToLower(token)
		node := s.Root
		for _, char := range token {
			charStr := string(char)
			if next, exists := node.Children[charStr]; exists {
				node = next
			} else {
				node = nil
			}
		}
		if node != nil {
			newFiles := make([]string, 0)
			for _, f := range node.Files {
				if f != path {
					newFiles = append(newFiles, f)
				}
			}
			node.Files = newFiles
		}
	}
}

func (j *Journal) UpdateJournal(op, path string) {
	Item := Storage{Op: op, Path: path}
	jsonData, err := json.Marshal(Item)
	if err != nil {
		panic(err)
	}
	j.File.Write(append(jsonData, '\n'))
	j.File.Sync()
}

func (s *Store) replayJournal(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Probelm retrieving items from storage.....")
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var item Storage
		json.Unmarshal(scanner.Bytes(), &item)
		switch item.Op {
		case "add":
			InternalAdd(s, item.Path)
		case "delete":
			InternalDelete(s, item.Path)
		default:
			continue
		}
	}
}

func (s *Store) AddItem(path string) {
	if s.Registary[path] {
		return
	}
	
	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		s.Journal.UpdateJournal("add", path)
		InternalAdd(s, path)
		return nil
	})

	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func (s *Store) DeleteItem(path string) {
	if !s.Registary[path] {
		return
	}
	s.Journal.UpdateJournal("delete", path)
	InternalDelete(s, path)
}

func main() {
	// path := "D:\\Learnings\\Learning GO\\SystemLevelGo\\exercises\\ch14\\index\\index.go"

	



	store := NewStore(`D:\Learnings\Learning GO\SystemLevelGo\exercises\ch14\index\store.wal`)

	store.replayJournal(`D:\Learnings\Learning GO\SystemLevelGo\exercises\ch14\index\store.wal`)

	store.AddItem("D:\\Learnings\\Learning GO\\SystemLevelGo\\exercises")
	items := store.FindItem("ch04")
	for _, i := range items {
		fmt.Println(i)
	}
	// store.DeleteItem("D:\\Learnings\\Learning GO\\SystemLevelGo\\exercises\\ch14\\index\\index.go")
	// items = store.FindItem("D:\\Learnings\\Learning GO\\SystemLevelGo\\exercises\\ch14\\index\\index.go")
	// for _, i := range items {
	// 	fmt.Println(i)
	// }
}
