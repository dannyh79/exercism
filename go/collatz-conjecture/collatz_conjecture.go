package collatzconjecture

import "errors"

func divideBy2(n int) int {
	return n / 2
}

func multiplyBy3AndAdd1(n int) int {
	return n*3 + 1
}

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return -1, errors.New("")
	}
	var i = 0
	for n != 1 {
		if n%2 == 0 {
			n = divideBy2(n)
		} else {
			n = multiplyBy3AndAdd1(n)
		}
		i++
	}
	return i, nil
}
