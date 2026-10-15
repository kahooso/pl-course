package main

func solveV2(s string) string {
	a := parse(s)
	n := len(a.letters)
	weight, lead := prepare(a)

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

	digit := make([]int, n)
	used := make([]bool, 10)

	if !assignV2(0, 0, weight, lead, minRest, maxRest, digit, used) {
		return ""
	}

	m := make(map[rune]int)
	for i, c := range a.letters {
		m[c] = digit[i]
	}
	return answer(a, m)
}

func prepare(a alphametic) ([]int, []bool) {
	n := len(a.letters)

	index := make(map[rune]int)
	for i, c := range a.letters {
		index[c] = i
	}

	weight := make([]int, n)
	for _, w := range a.addends {
		addWeights(w, +1, index, weight)
	}
	addWeights(a.result, -1, index, weight)

	lead := make([]bool, n)
	for c := range a.lead {
		lead[index[c]] = true
	}
	return weight, lead
}

func assignV2(i, sum int, weight []int, lead []bool, minRest, maxRest []int, digit []int, used []bool) bool {
	if sum+minRest[i] > 0 || sum+maxRest[i] < 0 {
		return false
	}
	if i == len(weight) {
		return sum == 0
	}

	for d := 0; d <= 9; d++ {
		if used[d] || (d == 0 && lead[i]) {
			continue
		}
		digit[i] = d
		used[d] = true
		if assignV2(i+1, sum+weight[i]*d, weight, lead, minRest, maxRest, digit, used) {
			return true
		}
		used[d] = false
	}
	return false
}
