package main

import (
	"strconv"
	"strings"
	"testing"
)

var solvers = []struct {
	name  string
	solve func(string) string
}{
	{"V0", solveV0},
	{"V1", solveV1},
	{"V2", solveV2},
	{"V4", solveV4},
}

var testPuzzles = []struct {
	name   string
	puzzle string
}{
	{"BE", "BE + BE = MOO"},
	{"SEND", "SEND + MORE = MONEY"},
	{"ELEVEN", "ELEVEN + NINE + FIVE + FIVE = THIRTY"},
}

func valid(puzzle, res string) bool {
	if !strings.Contains(res, "=") {
		return false
	}
	p := parse(puzzle)
	r := parse(res)

	if len(p.addends) != len(r.addends) {
		return false
	}
	pw := append(append([]string{}, p.addends...), p.result)
	rw := append(append([]string{}, r.addends...), r.result)

	letterToDigit := make(map[rune]byte)
	digitToLetter := make(map[byte]rune)
	for i := range pw {
		if len(pw[i]) != len(rw[i]) {
			return false
		}
		if len(rw[i]) > 1 && rw[i][0] == '0' {
			return false
		}
		for j, c := range pw[i] {
			d := rw[i][j]
			if x, ok := letterToDigit[c]; ok && x != d {
				return false
			}
			if x, ok := digitToLetter[d]; ok && x != c {
				return false
			}
			letterToDigit[c] = d
			digitToLetter[d] = c
		}
	}

	sum := 0
	for _, w := range r.addends {
		n, _ := strconv.Atoi(w)
		sum += n
	}
	n, _ := strconv.Atoi(r.result)
	return sum == n
}

func TestValid(t *testing.T) {
	cases := []struct {
		res  string
		want bool
	}{
		{"9567 + 1085 = 10652", true},
		{"9567 + 1085 = 10653", false},
		{"9567 + 1095 = 10662", false},
		{"", false},
	}
	for _, c := range cases {
		if got := valid("SEND + MORE = MONEY", c.res); got != c.want {
			t.Errorf("valid(%q) = %v, ждали %v", c.res, got, c.want)
		}
	}
}

func TestSolvers(t *testing.T) {
	for _, s := range solvers {
		for _, tp := range testPuzzles {
			res := s.solve(tp.puzzle)
			if !valid(tp.puzzle, res) {
				t.Errorf("%s, %s: получили %q — неверно", s.name, tp.puzzle, res)
			}
		}
	}
}

func BenchmarkSolvers(b *testing.B) {
	for _, s := range solvers {
		for _, tp := range testPuzzles {
			b.Run(s.name+"/"+tp.name, func(b *testing.B) {
				for b.Loop() {
					s.solve(tp.puzzle)
				}
			})
		}
	}
}
