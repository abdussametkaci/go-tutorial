package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
)

// sorts letters in a word (i.e. "morning" -> "gimnnor")
func sortRunes(str string) string {
	runes := bytes.Runes([]byte(str))
	var temp rune
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[j] < runes[i] {
				temp = runes[i]
				runes[i], runes[j] = runes[j], temp
			}

		}
	}
	return string(runes)
}

// load loads content of file fname into memory as []string
func load(fname string) ([]string, error) {
	if fname == "" {
		return nil, errors.New(
			"Dictionary file name cannot be empty.")
	}

	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	/*
		file, err := os.Open(fname)
		if err != nil {
				return nil, fmt.Errorf(
					"Unable to open file %s: %s", fname, err)
		}
	*/

	var lines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func write(fname string, anagrams map[string][]string) {
	_, err := os.OpenFile(
		fname,
		os.O_WRONLY+os.O_CREATE+os.O_EXCL,
		0644,
	)
	if err != nil {
		msg := fmt.Sprintf(
			"Unable to create output file: %v", err,
		)
		panic(msg)
	}

}

func mapWords(words []string) map[string][]string {
	anagrams := make(map[string][]string)
	for _, word := range words {
		wordSig := sortRunes(word)
		anagrams[wordSig] = append(anagrams[wordSig], word)
	}

	return anagrams
}

func main() {
	words, err := load("dict.txt")
	if err != nil {
		fmt.Println("Unable to load file:", err)
		os.Exit(1)
	}
	anagrams := mapWords(words)
	write("out.txt", anagrams)
}
