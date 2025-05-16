package module

import (
	"fmt"
	"regexp"
	"strconv"
)

var (
	CPFRegexp = regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`)
)

type Cpf struct {
	Cpf string
}

func (cpf *Cpf) ValidationCpf() error {
	if !cpf.isCPF(cpf.Cpf) {
		return fmt.Errorf("%s: %w", "document cpf is not valid!", nil)
	}

	return nil
}

// isCPF verifies if the given string is a valid CPF document.
func (cpf *Cpf) isCPF(doc string) bool {
	const (
		size = 9
		pos  = 10
	)

	return cpf.isDocument(doc, CPFRegexp, size, pos)
}

// isDocument generates the digits for a given CPF compares it with the original digits.
func (cpf *Cpf) isDocument(doc string, pattern *regexp.Regexp, size int, position int) bool {
	if !pattern.MatchString(doc) {
		return false
	}

	cleanNonDigits(&doc)

	// Invalidates documents with all digits equal.
	if cpf.allEq(doc) {
		return false
	}

	d := doc[:size]
	digit := cpf.calculateDigit(d, position)

	d = d + digit
	digit = cpf.calculateDigit(d, position+1)

	return doc == d+digit
}

// allEq checks if every rune in a given string is equal.
func (cpf *Cpf) allEq(doc string) bool {
	base := doc[0]
	for i := 1; i < len(doc); i++ {
		if base != doc[i] {
			return false
		}
	}

	return true
}

// calculateDigit calculates the next digit for the given document.
func (cpf *Cpf) calculateDigit(doc string, position int) string {
	var sum int
	for _, r := range doc {

		sum += toInt(r) * position
		position--

		if position < 2 {
			position = 9
		}
	}

	sum %= 11
	if sum < 2 {
		return "0"
	}

	return strconv.Itoa(11 - sum)
}
