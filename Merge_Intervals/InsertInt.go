package main

func insertInt(arr [][]int, new_arr []int) [][]int {
	var res [][]int
	s := len(arr)
	i := 0
	//Add intervals before new one
	for i < s && arr[i][1] < new_arr[0] {
		res = append(res, arr[i])
		i++
	}

	//Merge overlaps
	for i < s && arr[i][0] <= new_arr[1] {
		new_arr[0] = min(new_arr[0], arr[i][0])
		new_arr[1] = max(new_arr[1], arr[i][1])
		i++
	}
	res = append(res, new_arr)

	//Add rest
	for i < s {
		res = append(res, arr[i])
		i++
	}

	return res
}
