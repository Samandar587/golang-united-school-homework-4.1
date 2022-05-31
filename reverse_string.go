package reverse_string

func ReverseString(input string) (output string) {
	rev := make([]rune, 0)
	k := len(input) - 1
	for _, val := range input {
		rev[k] = val
		k--
	}
	output = string(rev)
	return output
}
