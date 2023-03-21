package dna

import (
	"errors"
)

type Histogram map[Nucleotide]int

type DNA []Nucleotide

type Nucleotide rune

var nucleotides = map[Nucleotide]Nucleotide{
	'A': 'A',
	'C': 'C',
	'G': 'G',
	'T': 'T',
}

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{}

	for nucleotide := range nucleotides {
		h[nucleotide] = 0
	}

	for _, nucleotide := range d {
		if _, ok := nucleotides[nucleotide]; !ok {
			return nil, errors.New("invalid nucleotide")
		}

		h[nucleotide]++
	}

	return h, nil
}
