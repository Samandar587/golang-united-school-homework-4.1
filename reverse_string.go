package reverse_string

func ReverseString(input string) (output string) {
	rev := make([]rune, 0)
	runes := []rune(input)
	k := len(runes) - 1
	for i := 0; i < len(runes); i++ {
		rev = append(rev, runes[k])
		k--
	}
	output = string(rev)
	return output
}
