package reverse

func Reverse(input string) string {
	r := ""
	for _, s := range input {
		r = string(s) + r
	}
	return r
}
