package valid

import (
	"net/mail"

	"github.com/FourSigma/validate/str"
)

func email(s string) error {
	_, err := mail.ParseAddress(s)
	return err
}

var (
	EMail = []str.Handler{email}
)
