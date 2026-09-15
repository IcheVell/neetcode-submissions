func threeSum(nums []int) [][]int {
	res := make([][]int, 0)

	m := make(map[[3]int]struct{})

	sort.Ints(nums)

	//fmt.Println(nums)

	for left, middle, right := 0, 1, len(nums) - 1; left < right - 1; {
		sum := nums[left] + nums[middle] + nums[right]

		for ; middle < right; {
			sum = nums[left] + nums[middle] + nums[right]

			//fmt.Println(sum, nums[left], nums[middle], nums[right])

			if sum == 0 {
				m[[3]int {nums[left], nums[middle], nums[right]}] = struct{}{}
				middle++
				right--
			} else if sum < 0 {
				middle++
			} else {
				right--
			}
		}

		left++
		middle = left + 1
		right = len(nums) - 1
	}

	for key, _ := range m {
		triplet := make([]int, 0, 3)

		triplet = append(triplet, key[0])
		triplet = append(triplet, key[1])
		triplet = append(triplet, key[2])

		res = append(res, triplet)
	}

	return res
}