package main

func findDuplicate(nums []int) int {
	s, f := nums[0], nums[0]

	for {
		s = nums[s]
		f = nums[nums[f]]

		if s == f {
			break
		}
	}

	s = nums[0]
	for s != f {
		s = nums[s]
		f = nums[f]
	}
	return s
}
