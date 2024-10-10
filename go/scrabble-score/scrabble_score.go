package scrabble

import "strings"

const alphabetsCount int = 26

var score_alphabets = map[int]string{
	1:  "aeioulnrst",
	2:  "dg",
	3:  "bcmp",
	4:  "fhvwy",
	5:  "k",
	8:  "jx",
	10: "qz",
}

func toScoreMap(score_alphabets_map map[int]string) map[string]int {
	m := make(map[string]int, alphabetsCount)
	for s, as := range score_alphabets_map {
		for _, a := range as {
			m[string(a)] = s
		}
	}
	return m
}

func Score(word string) int {
	s := 0
	m := toScoreMap(score_alphabets)
	for _, c := range word {
		s += m[strings.ToLower(string(c))]
	}
	return s
}
