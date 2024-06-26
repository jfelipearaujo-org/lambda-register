package cpf

import (
	"strconv"
	"strings"
)

func sumDigit(s string, table []int) int {
	if len(s) != len(table) {
		return 0
	}

	sum := 0

	for i, v := range table {
		c := string(s[i])
		d, err := strconv.Atoi(c)
		if err == nil {
			sum += v * d
		}
	}

	return sum
}

func Clean(cpf string) string {
	charsToRemove := []string{"/", ".", "-", " "}
	for _, char := range charsToRemove {
		cpf = strings.ReplaceAll(cpf, char, "")
	}
	return cpf
}
