package beesolver

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
)

type Dictionary struct {
	Words map[string][]string // must be a sorted set; A: [Aardvark, ...], B: [Babble, ...], C: ...
}

// NewDictionary reads the file in wordListPath one word per line
func NewDictionary(wordListPath string) (*Dictionary, error) {
	dictFile, err := os.Open(wordListPath)
	if err != nil {
		return nil, err
	}
	defer dictFile.Close()

	words := make(map[string][]string)
	wordScanner := bufio.NewScanner(dictFile)
	for wordScanner.Scan() {
		word := wordScanner.Text()
		firstLetter := strings.ToLower(string(word[0]))
		words[firstLetter] = append(words[firstLetter], strings.ToLower(word))
	}

	if err := wordScanner.Err(); err != nil {
		return nil, err
	}

	return &Dictionary{Words: words}, nil
}

func BeeRegex(bee *BeeSolver) (*regexp.Regexp, error) {
	letters := append(bee.AllowedLetters, bee.RequiredLetter)
	lettersSpaceSeparated := fmt.Sprintf("%s", letters) // [a b c d]

	removedWhitespace := strings.ReplaceAll(lettersSpaceSeparated, " ", "") // [abcd]

	return regexp.Compile(fmt.Sprintf("^%s{%d,%d}$", removedWhitespace, bee.MinimumLength, bee.MaximumLength)) // ^[abcd]{2,6}$
}

type BeeSolver struct {
	RequiredLetter string
	MinimumLength  uint
	MaximumLength  uint
	Dictionary     *Dictionary
	AllowedLetters []string
	ValidWordRegex *regexp.Regexp
}

func NewBeeSolver(allowedLetters []string, wordListPath, requiredLetter string, min, max uint) (*BeeSolver, error) {

	solver := &BeeSolver{
		RequiredLetter: requiredLetter,
		MinimumLength:  min,
		MaximumLength:  max,
		AllowedLetters: allowedLetters,
	}

	regex, err := BeeRegex(solver)
	if err != nil {
		log.Fatalf("failed creating regular expression: %s\n%v\n", err.Error(), solver)
		return nil, err
	}

	solver.ValidWordRegex = regex
	log.Println("starting solver with regex: " + regex.String())

	dictionary, err := NewDictionary(wordListPath)
	if err != nil {
		log.Fatalf("failed creating dictionary from path: '%s'\n%s\n", wordListPath, err.Error())
		return nil, err
	}

	solver.Dictionary = dictionary

	return solver, nil
}

func (b *BeeSolver) Solve() ([]string, error) {
	wg := sync.WaitGroup{}
	solutions := make(chan string)
	for _, letter := range append(b.AllowedLetters, b.RequiredLetter) {
		wg.Add(1)
		go func(l string) {
			defer wg.Done()
			for _, word := range b.Dictionary.Words[l] {
				if b.IsValidWord(word) {
					solutions <- word
				}
			}
		}(letter)
	}
	// start a routine to wait for the routines to finish and close the channel
	go func() {
		wg.Wait()
		close(solutions)
	}()
	result := make([]string, len(solutions))
	for s := range solutions {
		result = append(result, s)
	}

	return result, nil
}

func (b *BeeSolver) IsValidWord(word string) bool {
	return strings.Contains(word, b.RequiredLetter) && b.ValidWordRegex.FindString(word) == word
}
