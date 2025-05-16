package module

import (
	"fmt"
	"testing"
)

func assertEqual(t *testing.T, expected, actual interface{}) {
	switch v := actual.(type) {
	default:
		if expected != actual {
			decisionAssert(t, expected, actual)
			fmt.Println(v)
		}
	case error:
		if expected.(error).Error() != actual.(error).Error() {
			decisionAssert(t, expected, actual)
		}
		return
	}
}

func decisionAssert(t *testing.T, expected, actual interface{}) {
	t.Fatalf(`    
Not equal!
expected: %v
actual:   %v
`, expected, actual)
}

func testName(i int, name string) string {
	return fmt.Sprintf("%d_%s", i, name)
}
