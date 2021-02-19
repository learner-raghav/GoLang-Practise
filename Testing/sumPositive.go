package Testing

func sumPositive(input []int) int {
	ans := 0
	for _,j := range input {
		if j > 0 {
			ans += j
		}
	}
	return ans
}