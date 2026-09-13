func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	countZero := countZeros(nums)

	currProd := 1

	if countZero >= 2 {
		return res
	} else if countZero == 1 {
		for i := range nums {
			if nums[i] != 0 {
				continue
			}

			for j, num := range nums {
				if j == i {
					continue
				}

				currProd *= num
			}

			res[i] = currProd
			return res
		}
	}

	for i, num := range nums {
		if i == 0 {
			continue
		}

		currProd *= num
	}

	res[0] = currProd

	for j := range nums {
		if j == 0 {
			continue
		}

		currProd /= nums[j]
		currProd *= nums[j - 1]

		res[j] = currProd
	}

	return res
}

func countZeros(nums []int) int {
	cnt := 0
	
	for _, num := range nums {
		if num == 0 {
			cnt++
		}
	}

	return cnt
}
