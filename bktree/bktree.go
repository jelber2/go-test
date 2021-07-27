package bktree

import (
        "fmt"

	"github.com/agatan/bktree"
        levenshtein "github.com/creasty/go-levenshtein"
)

type word string

// Distance calculates hamming distance.
func (x word) Distance(e bktree.Entry) int {
        a := string(x)
        b := string(e.(word))

        return levenshtein.Distance(a, b)
}

func Add() {
	var tree bktree.BKTree
	// add words
	words := []string{"AAAA", "AAAG", "AAAT", "AAAC", "AACC", "AACT", "AACG", "AAGC"}
	for _, w := range words {
		tree.Add(word(w))
	}

	// spell check
        inputword := string("AAAA")
	results := tree.Search(word(inputword), 2)
	fmt.Printf("Input is %s. \nDid you mean:\n", inputword)
	for _, result := range results {
		fmt.Printf("\t%s (distance: %d)\n", result.Entry.(word), result.Distance)
	}
}
