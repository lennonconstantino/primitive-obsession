package module

import (
	"errors"
	"fmt"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	for i, tc := range []struct {
		name     string
		value    string
		expected error
	}{
		{"Email_valido_01", "lennon@email.com", nil},
		{"Email_valido_02", "<lennonconstantino@x.com>", nil},
		{"Email_invalido_01", "lennonemail.com", errors.New("mail: missing '@' or angle-addr")},
		{"Email_invalido_02", "lennon@email.", errors.New("mail: missing '@' or angle-addr")},
		{"Email_invalido_03", "lennon constantino", errors.New("mail: no angle-addr")},
		{"Email_invalido_04", "<lennonconstantino@x.com", errors.New("mail: unclosed angle-addr")},
		{"Email_invalido_05", "lennonconstantino@x.com>", fmt.Errorf("mail: expected single address, got %q", ">")},
		{"Email_invalido_06", "", errors.New("mail: no address")},
	} {
		t.Run(testName(i, tc.name), func(t *testing.T) {
			var email Email
			email.Email = tc.value
			assertEqual(t, tc.expected, email.Validation())
		})
	}
}
