package day1

import (
	"strconv"
	"strings"
)

type Solution struct {
}

func (s Solution) Solve(input string) string {
	numbers := strings.Split(input, "\n")
	var frequncy = 0
	for _, number := range numbers {
		i, _ := strconv.Atoi(number)
		frequncy += i
	}
	return strconv.Itoa(frequncy)
}
