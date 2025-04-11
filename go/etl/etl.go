package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	out := map[string]int{}
	for i, strs := range in {
		for _, s := range strs {
			out[strings.ToLower(s)] = i
		}
	}

	return out
}
