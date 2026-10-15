package main

func solveV1(s string) string {
	a := parse(s)
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

	digit := make([]int, n)
	used := make([]bool, 10)

	if !assignV1(0, weight, lead, digit, used) {
		return ""
	}

	m := make(map[rune]int)
	for i, c := range a.letters {
		m[c] = digit[i]
	}
	return answer(a, m)
}

func addWeights(w string, sign int, index map[rune]int, weight []int) {
	p := 1
	for k := len(w) - 1; k >= 0; k-- {
		weight[index[rune(w[k])]] += sign * p
		p *= 10
	}
}

func assignV1(i int, weight []int, lead []bool, digit []int, used []bool) bool {
	if i == len(weight) {
		return checkV1(weight, lead, digit)
	}

	for d := 0; d <= 9; d++ {
		if used[d] {
			continue
		}
		digit[i] = d
		used[d] = true
		if assignV1(i+1, weight, lead, digit, used) {
			return true
		}
		used[d] = false
	}
	return false
}

func checkV1(weight []int, lead []bool, digit []int) bool {
	for i := range lead {
		if lead[i] && digit[i] == 0 {
			return false
		}
	}

	sum := 0
	for i, w := range weight {
		sum += w * digit[i]
	}
	return sum == 0
}
