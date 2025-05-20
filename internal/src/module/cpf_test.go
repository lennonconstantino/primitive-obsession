package module

import (
	"fmt"
	"testing"
)

func TestValidateCpf(t *testing.T) {
	for i, tc := range []struct {
		name     string
		value    string
		expected error
	}{
		{"CPF valido", "123.456.789-09", nil},
		{"CPF invalido", "123.456.789-00", fmt.Errorf("%s: %w", "document cpf is not valid!", nil)},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			var cpf Cpf
			cpf.Cpf = tc.value
			assertEqual(t, tc.expected, cpf.Validation())
		})
	}
}
