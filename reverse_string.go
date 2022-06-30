package reverse_string

func ReverseString(input string) (output string) {
	for i := len(input) - 1; i > -1; i-- {
		output += string(input[i])
	}
	return
}
