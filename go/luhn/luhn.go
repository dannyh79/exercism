package luhn

func Valid(id string) bool {
	var n, d, m int

	for i := len(id) - 1; i >= 0; i-- {
		c := id[i]
		switch {
		case c == ' ':
			continue
		case c >= '0' && c <= '9':
			m = int(c - '0')
			if d%2 == 1 {
				m *= 2
			}
			if m > 9 {
				m -= 9
			}
			n += m
			d++
		default:
			return false
		}
	}
	return d > 1 && n%10 == 0
}
