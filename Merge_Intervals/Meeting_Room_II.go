package main

import "sort"

func minRoomsNeeded(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	//Sort the intervals
	sortInt(intervals)

	//Slice as min heap for storing end timings
	end := []int{}

	end = append(end, intervals[0][1]) //Add end of 1st meeting

	for i := 1; i < len(intervals); i++ {
		s := intervals[i][0]
		e := intervals[i][1]

		if s >= end[0] {
			end[0] = e
		} else {
			end = append(end, e)
		}
		sort.Ints(end)
	}
	return len(end)
}
