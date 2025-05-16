// age_test.go
package module

import (
	"errors"
	"testing"
)

func TestValidate(t *testing.T) {
	for i, tc := range []struct {
		name     string
		value    int
		expected error
	}{
		{"Idade_Maior_de_18", 36, nil},
		{"Idade_negativa", -14, errors.New("the age is not valid")},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			var age Age
			age.Age = tc.value
			assertEqual(t, tc.expected, age.Validate())
		})
	}
}
