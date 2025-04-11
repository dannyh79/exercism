package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T any](in []T, fn func(T) bool) []T {
	var out []T
	for _, i := range in {
		if fn(i) {
			out = append(out, i)
		}
	}
	return out
}

func Discard[T any](in []T, fn func(T) bool) []T {
	var out []T
	for _, i := range in {
		if !fn(i) {
			out = append(out, i)
		}
	}
	return out
}
