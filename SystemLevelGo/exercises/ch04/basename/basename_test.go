package main

import (
	"path/filepath"
	"testing"
)

var tests = []struct {
	name     string
	path     string
	expected string
}{
	// --- Basic Filenames ---
	{"Basic filename", "file.txt", "file.txt"},

	// --- Paths with Directories ---
	{"Absolute path standard", "/usr/local/bin/go", "go"},
	{"Absolute path with version", "/usr/local/bin/go1.20", "go1.20"},

	// --- Trailing Slashes and Windows ---
	{"Windows mixed slashes", "C:/Users/John/Documents/file.txt", "file.txt"},
	{"Trailing slash single", "/usr/local/bin/", "bin"},

	// --- Hidden Files ---
	{"Hidden file standard", ".gitignore", ".gitignore"},

	// --- Mixed Cases ---
	{"Double extension", "myfile.go.txt", "myfile.go.txt"},

	// --- Edge Cases: Empty and Space Inputs ---
	{"Single space filename", " ", " "},
	{"Spaces surrounding slash", " / ", " "},
	{"Path ending in space element", "a/b/c/ ", " "},

	// --- Extreme Slashing ---
	{"Single letter with slashes", "/a/", "a"},
	{"Single letter with trailing slashes", "/a//", "a"},

	// --- Relative Paths & Current/Parent Dirs ---
	{"Current directory dot", ".", "."},
	{"Parent directory dot dot", "..", ".."},
	{"Current directory dot with slash", "./", "."},
	{"Parent directory dot dot with slash", "../", ".."},
	{"Parent directory target literal", "a/b/..", ".."},
	{"Current directory target literal", "a/b/.", "."},
	{"Relative path child", "./file.txt", "file.txt"},
	{"Relative parent path child", "../file.txt", "file.txt"},

	// --- Cross-Platform Safe Native Paths ---
	{"Absolute path execution", "/Windows/System32/cmd.exe", "cmd.exe"},
	{"Relative path child element", "dir/file", "file"},
	{"Relative trailing slash", "dir/", "dir"},

	// --- Multiple and Complex Extensions ---
	{"Tarball double extension", "archive.tar.gz", "archive.tar.gz"},
	{"Multi dot web extension", "file.spec.test.js", "file.spec.test.js"},
	{"Consecutive dots extension", "file..txt", "file..txt"},
	{"Trailing dot extension", "file.", "file."},

	// --- Special Characters & Unicode ---
	{"Emoji in filename", "💡lightbulb.png", "💡lightbulb.png"},
	{"Japanese characters path", "/usr/bin/こんにちは", "こんにちは"},
	{"Spaces in path and name", "/path/with spaces/file name.txt", "file name.txt"},
	{"URL encoded spaces filename", "foo%20bar.txt", "foo%20bar.txt"},
	{"Hyphens and underscores", "test-file_v1.0.0", "test-file_v1.0.0"},

	// --- Non-Standard/Tricky Filenames ---
	{"Triple dot filename", "...", "..."},
	{"Single hyphen filename", "-", "-"},
	{"Leading underscore filename", "_file", "_file"},
	{"Hidden file with extension", ".file.txt", ".file.txt"},
}


func TestBasename(t *testing.T) {
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resb := basename(tc.path)
			if resb != tc.expected {
				t.Errorf("Error: Basename: %v Expected: %v Output: %v", tc.name ,tc.expected, resb)
			}
		})
		t.Run(tc.name, func(t *testing.T) {
			resp := filepath.Base(tc.path)
			if resp != tc.expected {
				t.Errorf("Error: Filepath: %v Expected: %v Output: %v", tc.name ,tc.expected, resp)
			}
		})
	}
}
