package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	m := map[rune]string {
		'\u2757':     "recommendation",
		'\U0001f50d': "search",
		'\u2600':     "weather",
	}
	for _, c := range log {
		if s, ok := m[c]; ok {
			return s
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	r := []rune(log)
	for i, c := range r {
		if c == oldRune {
			r[i] = newRune
		}
	}
	return string(r)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	r := []rune(log)
	return len(r) <= limit
}
