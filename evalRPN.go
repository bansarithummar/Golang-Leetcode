150. Evaluate Reverse Polish Notation

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, c := range tokens {
		switch c {
		case "+":
			b, a := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a+b)
		case "-":
			b, a := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, b-a)
		case "*":
			b, a := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a*b)
		case "/":
			b, a := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, b/a)
		default:
			num, _ := strconv.Atoi(c)
			stack = append(stack, num)
		}
	}
	return stack[0]
}

