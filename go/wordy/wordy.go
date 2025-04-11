package wordy

import "regexp"
import "strconv"

var numPattern = regexp.MustCompile(`-?\d+`)
var opPattern = regexp.MustCompile(`(divided|minus|multiplied|plus)`)
var validPattern = regexp.MustCompile(`What is -?\d+(?: (?:plus|minus|divided by|multiplied by) -?\d+)*\?`)

func Answer(question string) (int, bool) {
	if !validPattern.MatchString(question) {
		return 0, false
	}

	nums := numPattern.FindAllString(question, -1)
	ops := opPattern.FindAllString(question, -1)

	if len(ops) != len(nums)-1 {
		return 0, false
	}

	sum, _ := strconv.Atoi(nums[0])

	for i, op := range ops {
		n, _ := strconv.Atoi(nums[i+1])

		switch op {
		case "divided":
			sum /= n
		case "minus":
			sum -= n
		case "multiplied":
			sum *= n
		default:
			sum += n
		}
	}

	return sum, true
}
