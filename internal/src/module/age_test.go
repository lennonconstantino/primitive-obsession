// age_test.go
package module

import (
	"errors"
	"io"
	"os"
	"testing"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		name     string
		value    int
		expected error
	}{
		{"positive numbers", 36, nil},
		{"negative numbers", -14, errors.New("the age is not valid")},
	}

	i := 0
	for _, tc := range testCases {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			var age Age
			age.Age = tc.value

			r, w, _ := os.Pipe()
			os.Stdout = w

			age.Validate()

			capturedOutput, _ := io.ReadAll(r)

			expectedOutput := "\n"
			if string(capturedOutput) != expectedOutput {
				t.Errorf("Expected output '%s', but got '%s'", expectedOutput, string(capturedOutput))
			}
		})
		i++
	}
}
