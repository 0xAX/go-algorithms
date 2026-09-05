package stack

func isExpressionBalanced(text string) bool {
	stack := make([]string, 0, len(text))

	for i, char := range text {
		if char == '(' {
			stack = append(stack, string(char))
		}
		if char == ')' {
			stack = remove(stack, i-1)
		}
	}

	return len(stack) == 0
}

func remove(s []string, index int) []string {
	if index >= len(s) {
		return nil
	}
	return append(s[:index], s[index+1:]...)
}
