package bktree-test

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
        words := []string{"dog", "cat"}
        for _, w := range words {
                tree.Add(word(w))
        }
        return tree
}
