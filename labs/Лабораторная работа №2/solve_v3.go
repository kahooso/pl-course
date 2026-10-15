package main

import "sort"

func solveV3(s string) string {
	a := parse(s)
	n := len(a.letters)
	weight, lead := prepare(a)

	order := make([]int, n)
	for i := range order {
		order[i] = i
	}
	sort.Slice(order, func(x, y int) bool {
		return abs(weight[order[x]]) > abs(weight[order[y]])
	})

	w := make([]int, n)
	l := make([]bool, n)
	for k, i := range order {
		w[k] = weight[i]
		l[k] = lead[i]
	}

	minRest, maxRest := bounds(w)
	digit := make([]int, n)
	used := make([]bool, 10)

	if !assignV2(0, 0, w, l, minRest, maxRest, digit, used) {
		return ""
	}

	m := make(map[rune]int)
	for k, i := range order {
		m[a.letters[i]] = digit[k]
	}
	return answer(a, m)
}

func bounds(weight []int) ([]int, []int) {
	n := len(weight)
	minRest := make([]int, n+1)
	maxRest := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		minRest[i] = minRest[i+1]
		maxRest[i] = maxRest[i+1]
		if weight[i] < 0 {
			minRest[i] += 9 * weight[i]
		} else {
			maxRest[i] += 9 * weight[i]
		}
	}
	return minRest, maxRest
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
