package reverse_string

func ReverseString(input string) (output string) {
	rev := make([]uint8, 0)
	bytes := []byte(input)
	k := len(input) - 1
	for i := 0; i < len(input); i++ {
		rev = append(rev, bytes[k])
		k--
	}
	output = string(rev)
	return output
}
