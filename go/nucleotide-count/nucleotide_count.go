package dna

import "fmt"



type Histogram map[rune]int

type DNA string

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, n := range d {
		_, ok := h[n]
		if !ok {
			return nil, fmt.Errorf("Invalid nucleotide: %v", n)
		}
		h[n]++
	}
	return h, nil
}
