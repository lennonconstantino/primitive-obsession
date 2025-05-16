package module

import (
	"net/mail"
)

type Email struct {
	Email string
}

// Validation performs validations regarding the email
func (email *Email) Validation() error {

	if _, erro := mail.ParseAddress(email.Email); erro != nil {
		return erro
	}

	return nil
}
