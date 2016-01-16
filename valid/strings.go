package valid

import (
	"net/mail"
	"net/url"

	"github.com/FourSigma/validate/str"
)

func email(s string) error {
	_, err := mail.ParseAddress(s)
	return err
}

func urlParse(s string) error {
	_, err := url.Parse(s)

	return err
}

var (
	EMail = []str.Handler{email}
	URL   = []str.Handler{urlParse}
)
