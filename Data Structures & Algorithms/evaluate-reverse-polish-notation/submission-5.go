func evalRPN(tokens []string) int {
	nums := make([]int, 0)

	for i := range tokens {
		if val, err := strconv.ParseInt(tokens[i], 10, 64); err == nil {
			nums = append(nums, int(val))
			continue
		}

		num1 := nums[len(nums) - 1]
		num2 := nums[len(nums) - 2]

		switch tokens[i] {
			case "+":
				nums = append(nums[:len(nums) - 2], num1 + num2)

			case "-":
				nums = append(nums[:len(nums) - 2], num2 - num1)

			case "*":
				nums = append(nums[:len(nums) - 2], num1 * num2)

			case "/":
				nums = append(nums[:len(nums) - 2], num2 / num1)
		}
	}

	return nums[0]
}
