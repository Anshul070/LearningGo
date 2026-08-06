package main

import (
	"strings"
	"fmt"
)

type Node struct {
	prefix   string
	children []*Node
	isWord   bool
}

type RedixTree struct {
	Root *Node
}

func NewRedixTree () *RedixTree {
	return new(RedixTree{
		Root: new(Node{}),
	})
}


func compareStrings(a, b string) int {
	i := 0
	for i < len(a) && i < len(b) && a[i] == b[i] {
		i++
	}
	return i
}

func (r *RedixTree) Insert(word string) {
	current := r.Root
	for {
		var child *Node
		var i int
		for _, c := range current.children {
			i = compareStrings(word, c.prefix)
			if i > 0 {
				child = c
				break
			}
		}

		if child == nil {
			newNode := new(Node{
				prefix: word,
				isWord: true,
			})

			current.children = append(current.children, newNode);
			return
		}

		if i < len(child.prefix) {
			splitNode := new(Node{
				prefix: child.prefix[i:],
				children: child.children,
				isWord : child.isWord,
			});

			child.prefix = child.prefix[:i]
			child.children = []*Node{splitNode}
			child.isWord = false
			return
		}


		word = word[i:]
		if word == "" {
			child.isWord = true
			return
		}
		current = child
	}
}

func (r *RedixTree) Search(word string) bool {
	current := r.Root
	for {
		var child *Node
		for _, c := range current.children {
			if strings.HasPrefix(word, c.prefix) {
				child = c
				break
			}
		}

		if child == nil {
			return false
		}

		if len(word) == len(child.prefix) {
			return child.isWord
		}

		word = word[len(child.prefix):]
		current = child
	}
}

func (r *RedixTree) StartsWith(word string) bool {
	current := r.Root;
	for {
		var child *Node
		for _,c := range current.children{
			if strings.HasPrefix(c.prefix, word) {
				return true
			}
			if strings.HasPrefix(word, c.prefix){
				child = c
				break
			}
		}
		if child == nil{
			return false
		}

		if len(word) <= len(child.prefix){
			return true
		}

		word = word[len(child.prefix):]
		current = child
	}
}

func main() {
	tree := NewRedixTree()

	words := []string{"go", "gone", "good", "gopher", "goblin", "zebra"}
	for _, word := range words {
		tree.Insert(word)
	}

	tests := []string{"go", "gone", "god", "gob", "gopher", "zeb", "zebra", ""}
	for _, t := range tests {
		fmt.Printf("Search(%q): %v\n", t, tree.Search(t))
	}

	prefixes := []string{"go", "gope", "zeb", "x"}
	for _, p := range prefixes {
		fmt.Printf("StartsWith(%q): %v\n", p, tree.StartsWith(p))
	}
}