func twoSum(nums []int, target int) []int {   
   m := make(map[int]int)

   for i, num := range nums {
	if val, ok := m[num]; ok {
		return []int{val, i}
	}

	m[target - num] = i
   }

   return []int{} 
}
