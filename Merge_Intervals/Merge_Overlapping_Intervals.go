package main

func mergeOverlapInt(arr [][]int) [][]int {
	n := len(arr)
	sortInt(arr)

	ans := [][]int{}

	for i := 0; i < n; i++ {
		s := arr[i][0]
		e := arr[i][1]

		if len(ans) > 0 && e <= ans[len(ans)-1][1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if arr[j][0] <= e {
				e = max(e, arr[j][1])
			} else {
				break
			}
		}
		ans = append(ans, []int{s, e})
	}
	return ans

}
