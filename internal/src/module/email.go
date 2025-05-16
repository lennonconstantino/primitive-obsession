package module

import (
	"net/mail"
)

type Email struct {
	Email string
}

func (email *Email) Validation() error {

	if _, erro := mail.ParseAddress(email.Email); erro != nil {
		return erro
	}

	return nil
}
